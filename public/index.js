function HeroApp() {
  const [sentence, setSentence] = React.useState('');
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState('');

  const fetchHeroSentence = async () => {
    setLoading(true);
    setError('');
    
    try {
      const response = await fetch('https://nimble-holy-meadow-800.fly.dev/api/heros');
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
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
      <div className="bg-white p-6 rounded-lg shadow-md w-full max-w-md">
        <h1 className="text-2xl font-bold text-center mb-6">Who the fuck is my Nimble Character?</h1>
        <button 
          onClick={fetchHeroSentence}
          disabled={loading}
          className="w-full bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded transition-colors duration-300 disabled:bg-blue-300 yea-button"
        >
          {loading ? 'Loading...' : 'Yea'}
        </button>
        
        {sentence && (
          <div className="mt-6 p-4 bg-gray-50 rounded-md sentence-container">
            <p className="text-lg text-center sentence-text">{sentence}</p>
          </div>
        )}
        
        {error && (
          <div className="mt-4 p-3 bg-red-100 text-red-700 rounded-md error-container">
            <p>{error}</p>
          </div>
        )}
      </div>
    </div>
  );
}
