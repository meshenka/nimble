function HeroApp() {
  const [sentence, setSentence] = React.useState('');
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState('');

  const fetchHeroSentence = async () => {
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
        <h1 className="app-title">Who the fuck is my Nimble Character?</h1>
        <button 
          onClick={fetchHeroSentence}
          disabled={loading}
          className="btn"
        >
          {loading ? 'Loading...' : 'Yea'}
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
}
