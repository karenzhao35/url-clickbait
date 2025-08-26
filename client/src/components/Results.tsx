import { Sparkles, Copy } from 'lucide-react';

interface ResultsProps {
  generatedUrl: string;
  onCopy: () => void;
}

const Results = ({ generatedUrl, onCopy }: ResultsProps) => {
  if (!generatedUrl) return null;

  return (
    <div className="w-full max-w-2xl mx-auto mt-6 lg:mt-8 bg-green-900/30 backdrop-blur-lg rounded-3xl p-6 lg:p-8 border border-green-500/30 shadow-2xl animate-fade-in">
      <h3 className="text-xl lg:text-2xl font-bold text-green-400 mb-4 flex items-center space-x-2">
        <Sparkles className="animate-spin w-5 h-5 lg:w-6 lg:h-6" />
        <span>YOUR EXPONENTIALLY OPTIMIZED URL</span>
      </h3>
      
      <div className="bg-gray-900/50 rounded-xl p-4 mb-4">
        <div className="flex flex-col sm:flex-row items-start sm:items-center gap-3 sm:gap-0 sm:justify-between">
          <code className="text-cyan-400 font-mono text-xs lg:text-sm break-all flex-1 pr-0 sm:pr-4">
            {generatedUrl}
          </code>
          <button
            onClick={onCopy}
            className="bg-purple-600 hover:bg-purple-500 px-3 py-2 rounded-lg transition-colors flex items-center space-x-1 shrink-0 w-full sm:w-auto justify-center sm:justify-start"
          >
            <Copy size={16} />
            <span className="text-sm font-bold">COPY</span>
          </button>
        </div>
      </div>
      
      <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-center">
        <div className="bg-purple-900/50 rounded-lg p-4">
          <div className="text-xl lg:text-2xl font-bold text-yellow-400">+{Math.floor(Math.random() * 5000)}%</div>
          <div className="text-xs lg:text-sm">Click-through Rate</div>
        </div>
        <div className="bg-blue-900/50 rounded-lg p-4">
          <div className="text-xl lg:text-2xl font-bold text-green-400">+{Math.floor(Math.random() * 5000)}%</div>
          <div className="text-xs lg:text-sm">Tech Bro Engagement</div>
        </div>
      </div>
      
      <p className="text-center mt-4 text-gray-300 text-xs lg:text-sm px-4">
        Congratulations! Your URL is now optimized for maximum disruption in the B2B space!
      </p>
    </div>
  );
};

export default Results;
