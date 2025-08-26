import { useState } from 'react';

import Header from './components/Header';
import UrlForm from './components/UrlForm';
import Results from './components/Results';
import Background from './components/Background';
import Footer from './components/Footer';

function App() {
  const [url, setUrl] = useState('');
  const [expiry, setExpiry] = useState(30);
  const [generatedUrl, setGeneratedUrl] = useState('');
  const [isGenerating, setIsGenerating] = useState(false);

  const handleGenerate = async () => {
    if (!url) return;

    const data = {
      old_url: url,
      expired_at: expiry.toString(),
    }
    
    setIsGenerating(true);
    try {
      const res = await fetch('http://localhost:3000/clickbait', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
      const result = await res.json();
      setGeneratedUrl(result.short_url);
      setIsGenerating(false);
    } catch (error) {
      console.error('Error sending data:', error);
      setIsGenerating(false);
    }
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(generatedUrl);
  };

  return (
    <div className="min-h-screen flex flex-col bg-gradient-to-br from-purple-900 via-blue-900 to-indigo-900 text-white relative overflow-hidden">
      <Background />

      <div className="relative z-10 flex flex-col min-h-screen">
        <div className="flex-1 flex flex-col justify-center px-4 py-6 lg:py-8">
          <div className="container mx-auto max-w-6xl">
            <Header />
            
            <div className="flex flex-col items-center space-y-6">
              <UrlForm
                url={url}
                setUrl={setUrl}
                expiry={expiry}
                setExpiry={setExpiry}
                onGenerate={handleGenerate}
                isGenerating={isGenerating}
              />
              
              <Results
                generatedUrl={generatedUrl}
                onCopy={copyToClipboard}
              />
            </div>
          </div>
        </div>
        
        <Footer />
      </div>
    </div>
  );
}

export default App
