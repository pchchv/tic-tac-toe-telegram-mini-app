import { useState } from 'react';
import type { CellValue } from './types';
import './App.css';

function App() {
  // Local state to store the status of 9 cells (for now, we’ll populate it with test data for layout purposes)
  const [board, setBoard] = useState<CellValue[]>([
    'X', null, 'O',
    null, 'X', null,
    null, null, 'O'
  ]);

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
