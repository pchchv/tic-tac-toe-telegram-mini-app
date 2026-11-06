import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { init, miniApp } from '@telegram-apps/sdk-react'
import App from './App.tsx'
import './index.css'

try {
  init();

  if (miniApp.ready.isAvailable()) {
    miniApp.ready();
  }
} catch (error) {
  console.warn('Telegram SDK is not initialized (standard browser):', error);
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
