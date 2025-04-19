import React, { useState } from 'react';

interface HeroAppProps {
  // Add any props your component needs here
}

const HeroApp: React.FC<HeroAppProps> = () => {
  const [sentence, setSentence] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  const fetchHeroSentence = async (): Promise<void> => {
    setLoading(true);
    setError('');
    
    try {
      const response = await fetch('/api/heros');
      const data = await response.json();
      setSentence(data.sentence);
    } catch (err) {
      setError('Failed to fetch data. Make sure your API is running.');
      console.error('Error fetching data:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="app-container">
      <h1 className="app-title">My F*cking Nimble 5e Hero</h1>
      <button
        onClick={fetchHeroSentence}
        disabled={loading}
        className="btn"
      >
        {loading ? 'Loading...' : 'Let\'s See!'}
      </button>
      
      {sentence && (
        <div className="sentence-container">
          <p className="sentence-text">{sentence}</p>
        </div>
      )}
      
      {error && (
        <div className="error-container">
          <p>{error}</p>
        </div>
      )}
    </div>
  );
};

export default HeroApp;
