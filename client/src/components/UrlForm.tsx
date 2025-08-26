import { Globe, Clock, Rocket, Brain } from 'lucide-react';

interface UrlFormProps {
  url: string;
  setUrl: (url: string) => void;
  expiry: number;
  setExpiry: (expiry: number) => void;
  onGenerate: () => void;
  isGenerating: boolean;
}

const UrlForm = ({ url, setUrl, expiry, setExpiry, onGenerate, isGenerating }: UrlFormProps) => {
  return (
    <div className="w-full max-w-2xl mx-auto bg-black/30 backdrop-blur-lg rounded-3xl p-6 lg:p-8 border border-purple-500/30 shadow-2xl">
      <div className="space-y-6">
        {/* URL Input */}
        <div>
          <label className="block text-base lg:text-lg font-bold mb-3 flex items-center space-x-2">
            <Globe className="text-cyan-400 w-5 h-5 lg:w-6 lg:h-6" />
            <span>YOUR BASIC, NON-DISRUPTIVE URL:</span>
          </label>
          <input
            type="url"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            placeholder="https://boring-corporate-website.com"
            className="w-full px-4 py-3 lg:py-4 bg-gray-900/50 border-2 border-purple-500/50 rounded-xl text-white placeholder-gray-400 focus:border-cyan-400 focus:outline-none transition-all duration-300 text-sm lg:text-base"
          />
        </div>

        {/* Expiry Input */}
        <div>
          <label className="block text-base lg:text-lg font-bold mb-3 flex items-center space-x-2">
            <Clock className="text-yellow-400 w-5 h-5 lg:w-6 lg:h-6" />
            <span>EXPIRATION TIME (DAYS):</span>
          </label>
          <input
            type="number"
            value={expiry}
            onChange={(e) => setExpiry(Number(e.target.value))}
            min="1"
            max="365"
            className="w-full px-4 py-3 lg:py-4 bg-gray-900/50 border-2 border-purple-500/50 rounded-xl text-white focus:border-cyan-400 focus:outline-none transition-all duration-300 text-sm lg:text-base"
          />
          <p className="text-xs lg:text-sm text-gray-400 mt-2">
            Pro tip: Shorter expiry = more FOMO = higher conversion rates!
          </p>
        </div>

        {/* Generate Button */}
        <button
          onClick={onGenerate}
          disabled={!url || isGenerating}
          className={`w-full py-3 lg:py-4 px-8 rounded-xl font-black text-lg lg:text-xl transition-all duration-300 transform ${
            !url || isGenerating 
              ? 'bg-gray-600 cursor-not-allowed' 
              : 'bg-gradient-to-r from-pink-500 via-purple-500 to-cyan-500 hover:scale-105 hover:shadow-2xl active:scale-95'
          }`}
        >
          {isGenerating ? (
            <div className="flex items-center justify-center space-x-3">
              <Brain className="animate-spin w-5 h-5 lg:w-6 lg:h-6" />
              <span className="text-sm lg:text-xl">AI IS OPTIMIZING FOR MAXIMUM ENGAGEMENT...</span>
            </div>
          ) : (
            <div className="flex items-center justify-center space-x-3">
              <Rocket className="w-5 h-5 lg:w-6 lg:h-6" />
              <span className="text-sm lg:text-xl">GENERATE DISRUPTIVE CLICKBAIT</span>
            </div>
          )}
        </button>
      </div>
    </div>
  );
};

export default UrlForm;
