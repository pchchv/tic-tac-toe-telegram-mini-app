package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TelegramUser represents the parsed user object from initData
type TelegramUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// ValidateInitData verifies the cryptographic signature of Telegram WebApp initData
// and returns the authenticated TelegramUser if valid.
func ValidateInitData(initData, botToken string, expireAfter time.Duration) (*TelegramUser, error) {
	if initData == "" {
		return nil, errors.New("initData is empty")
	}
	if botToken == "" {
		return nil, errors.New("bot token is not configured")
	}

	// Parse query string into key-value pairs
	values, err := url.ParseQuery(initData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse initData: %w", err)
	}

	receivedHash := values.Get("hash")
	if receivedHash == "" {
		return nil, errors.New("hash is missing in initData")
	}

	// Check auth_date expiration if expireAfter is specified
	if expireAfter > 0 {
		authDateStr := values.Get("auth_date")
		authDateUnix, err := strconv.ParseInt(authDateStr, 10, 64)
		if err != nil {
			return nil, errors.New("invalid auth_date format")
		}
		if time.Since(time.Unix(authDateUnix, 0)) > expireAfter {
			return nil, errors.New("initData has expired")
		}
	}

	// Collect all keys except "hash" and sort them alphabetically
	var keys []string
	for k := range values {
		if k != "hash" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// Build data-check-string formatted as key=value joined by newline
	var dataCheckParts []string
	for _, k := range keys {
		dataCheckParts = append(dataCheckParts, fmt.Sprintf("%s=%s", k, values.Get(k)))
	}
	dataCheckString := strings.Join(dataCheckParts, "\n")

	// Step 1: Compute secret key = HMAC_SHA256("WebAppData", botToken)
	secretKeyMac := hmac.New(sha256.New, []byte("WebAppData"))
	secretKeyMac.Write([]byte(botToken))
	secretKey := secretKeyMac.Sum(nil)

	// Step 2: Compute expected hash = HMAC_SHA256(secretKey, dataCheckString)
	hashMac := hmac.New(sha256.New, secretKey)
	hashMac.Write([]byte(dataCheckString))
	calculatedHash := hex.EncodeToString(hashMac.Sum(nil))

	// Compare hashes securely using constant-time comparison
	if !hmac.Equal([]byte(calculatedHash), []byte(receivedHash)) {
		return nil, errors.New("invalid initData signature")
	}

	// Extract and unmarshal user JSON data
	userJSON := values.Get("user")
	if userJSON == "" {
		return nil, errors.New("user data is missing in initData")
	}

	var user TelegramUser
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	return &user, nil
}
