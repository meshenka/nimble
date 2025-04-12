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

  // Create button element
  const button = React.createElement(
    'button',
    { 
      onClick: fetchHeroSentence,
      disabled: loading,
      className: 'btn'
    },
    loading ? 'Loading...' : 'Yea'
  );

  // Create sentence display if sentence exists
  let sentenceElement = null;
  if (sentence) {
    sentenceElement = React.createElement(
      'div',
      { className: 'sentence-container' },
      React.createElement(
        'p',
        { className: 'sentence-text' },
        sentence
      )
    );
  }

  // Create error display if error exists
  let errorElement = null;
  if (error) {
    errorElement = React.createElement(
      'div',
      { className: 'error-container' },
      React.createElement('p', null, error)
    );
  }

  // Create app title
  const title = React.createElement(
    'h1',
    { className: 'app-title' },
    'Hero Sentence App'
  );

  // Create app container
  const appContainer = React.createElement(
    'div',
    { className: 'app-container' },
    [title, button, sentenceElement, errorElement].filter(Boolean)
  );

  return appContainer;
}
