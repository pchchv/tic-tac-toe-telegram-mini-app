import { useEffect, useState } from 'react';
import { themeParams, hapticFeedback } from '@telegram-apps/sdk-react';
import { calculateWinner } from './utils';
import type { CellValue } from './types';
import './App.css';

function App() {
  // Game state
  const [board, setBoard] = useState<CellValue[]>(Array(9).fill(null));
  const [isXNext, setIsXNext] = useState<boolean>(true);

  // Determine game status
  const winner = calculateWinner(board);
  const isDraw = !winner && board.every((cell) => cell !== null);

  useEffect(() => {
    // Check if theme parameters component is available
    if (themeParams) {
      const root = document.documentElement;
      const state = themeParams.state();
      
      // Map current Telegram theme variables to CSS custom properties with safe fallbacks
      root.style.setProperty('--tg-theme-bg-color', state.bgColor || '#17212b');
      root.style.setProperty('--tg-theme-text-color', state.textColor || '#f5f5f5');
      root.style.setProperty('--tg-theme-button-color', state.buttonColor || '#2481cc');
      root.style.setProperty('--tg-theme-button-text-color', state.buttonTextColor || '#ffffff');
      root.style.setProperty('--tg-theme-hint-color', state.hintColor || '#708499');
      root.style.setProperty('--tg-theme-secondary-bg-color', state.secondaryBgColor || '#101921');
    }
  }, []);

  const handleCellClick = (index: number) => {
    // Return early if the cell is already taken or game is finished
    if (board[index] || winner || isDraw) {
      return;
    }

    const newBoard = [...board];
    const currentTurn = isXNext ? 'X' : 'O';
    newBoard[index] = currentTurn;
    
    // Calculate intermediate results to trigger haptics immediately
    const nextWinner = calculateWinner(newBoard);
    const nextDraw = !nextWinner && newBoard.every((cell) => cell !== null);

    // Trigger Telegram Haptic Feedback safely
    if (hapticFeedback) {
      if (nextWinner) {
        // Success notification for a win
        if (hapticFeedback.notificationOccurred.isAvailable()) {
          hapticFeedback.notificationOccurred('success');
        }
      } else if (nextDraw) {
        // Warning notification for a draw
        if (hapticFeedback.notificationOccurred.isAvailable()) {
          hapticFeedback.notificationOccurred('warning');
        }
      } else {
        // Soft impact tap for a regular turn
        if (hapticFeedback.impactOccurred.isAvailable()) {
          hapticFeedback.impactOccurred('light');
        }
      }
    }

    setBoard(newBoard);
    setIsXNext(!isXNext);
  };

  const handleReset = () => {
    // Light vibration when resetting the field
    if (hapticFeedback && hapticFeedback.impactOccurred.isAvailable()) {
      hapticFeedback.impactOccurred('medium');
    }
    setBoard(Array(9).fill(null));
    setIsXNext(true);
  };

  // Generate status text message
  let statusMessage = `Next turn: ${isXNext ? 'X' : 'O'}`;
  if (winner) {
    statusMessage = `Winner: ${winner} 🎉`;
  } else if (isDraw) {
    statusMessage = `It's a draw! 🤝`;
  }

  return (
    <div className="app-container">
      <header className="app-header">
        <h1>Tic-Tac-Toe ❌⭕</h1>
        <div className={`status-turn ${winner ? 'status-winner' : ''}`}>
          {statusMessage}
        </div>
      </header>

      <main className="game-zone">
        <div className="board">
          {board.map((value, index) => (
            <button 
              key={index} 
              className={`cell ${value ? `cell-${value.toLowerCase()}` : ''}`}
              onClick={() => handleCellClick(index)}
              disabled={!!value || !!winner || isDraw}
            >
              {value}
            </button>
          ))}
        </div>
      </main>

      <footer className="app-footer">
        <button className="reset-btn" onClick={handleReset}>
          Restart Game
        </button>
      </footer>
    </div>
  );
}

export default App;
