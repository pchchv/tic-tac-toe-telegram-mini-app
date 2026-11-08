import { useEffect, useState } from 'react';
import { themeParams } from '@telegram-apps/sdk-react';
import type { CellValue } from './types';
import './App.css';

function App() {
  // Initialize empty 3x3 board matrix
  const [board, setBoard] = useState<CellValue[]>(Array(9).fill(null));

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

  // Placeholder for cell click
  const handleCellClick = (index: number) => {
    console.log(`Clicked cell index: ${index}`);
  };

  return (
    <div className="app-container">
      <header className="app-header">
        <h1>Tic-Tac-Toe ❌⭕</h1>
        <div className="status-turn">Next turn: <span>X</span></div>
      </header>

      <main className="game-zone">
        <div className="board">
          {board.map((value, index) => (
            <button 
              key={index} 
              className={`cell ${value ? `cell-${value.toLowerCase()}` : ''}`}
              onClick={() => handleCellClick(index)}
            >
              {value}
            </button>
          ))}
        </div>
      </main>

      <footer className="app-footer">
        <button className="reset-btn" onClick={() => setBoard(Array(9).fill(null))}>
          Reset Board
        </button>
      </footer>
    </div>
  );
}

export default App;
