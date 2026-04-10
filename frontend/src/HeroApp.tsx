import React, { useState, useEffect } from 'react';
import { HandlerHeroResponse } from './types';
import Link from './Link';

interface HeroAppProps {
  // Add any props your component needs here
}

const HeroApp: React.FC<HeroAppProps> = () => {
  const [response, setResponse] = useState<HandlerHeroResponse | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  const fetchHeroSentence = async (seed?: string): Promise<void> => {
    setLoading(true);
    setError('');

    try {
      const url = seed ? `/api/heros/${seed}` : '/api/heros';
      const response = await fetch(url);
      const data = await response.json();
      setResponse(data)
    } catch (err) {
      setError('Failed to fetch data. Make sure your API is running.');
      console.error('Error fetching data:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const seed = params.get('seed');
    if (seed) {
      fetchHeroSentence(seed);
    }
  }, []);

  return (
    <div className="app-container">
      <h1 className="app-title">My F*cking Nimble 5e Hero</h1>
      <button
        onClick={() => fetchHeroSentence()}
        disabled={loading}
        className="btn"
      >
        {loading ? 'Loading...' : 'Let\'s See!'}
      </button>

      {response && (
        <div className="sentence-container">
          <p className="sentence-text">{response.sentence}</p>
          <Link href={`/?seed=${response.id}`} text='Bookmark' />
        </div>
      )}

      {error && (
        <div className="error-container">
          <p>{error}</p>
        </div>
      )}
    </div >
  )
};

export default HeroApp;
