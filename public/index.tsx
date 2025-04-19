import React from 'react';
import ReactDOM from 'react-dom/client';
import HeroApp from './HeroApp'; // Adjust the path as needed

// Render the app
const container = document.getElementById('root');
if (container) {
  const root = ReactDOM.createRoot(container);
  root.render(<HeroApp />);
} else {
  console.error('Root element not found');
}
