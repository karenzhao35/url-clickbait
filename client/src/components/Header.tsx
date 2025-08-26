import { Sparkles, Zap, TrendingUp } from 'lucide-react';

const Header = () => {
  const techBuzzwords = [
    "REVOLUTIONARY",
    "AI-POWERED", 
    "DISRUPTIVE",
    "BLOCKCHAIN-NATIVE",
    "GAME-CHANGING",
    "NEXT-GEN"
  ];

  return (
    <div className="text-center mb-8 lg:mb-12 relative">
      <div className="absolute top-0 left-0 right-0 flex flex-wrap justify-center items-center gap-2 lg:gap-4 animate-bounce px-4">
        {techBuzzwords.map((word, i) => (
          <span key={i} className="text-xs lg:text-sm bg-gradient-to-r from-pink-500 to-violet-500 px-2 py-1 rounded-full font-bold opacity-80">
            {word}
          </span>
        ))}
      </div>
      
      <h1 className="text-4xl md:text-5xl lg:text-6xl font-black bg-gradient-to-r from-cyan-400 via-purple-400 to-pink-400 bg-clip-text text-transparent mb-4 mt-8 lg:pt-10 m:pt-18 pt-25">
        ClickBait.<span className="text-yellow-400">ai</span>
      </h1>
      
      <div className="flex justify-center space-x-2 mb-6">
        <Sparkles className="text-yellow-400 animate-spin w-5 h-5 lg:w-6 lg:h-6" />
        <Zap className="text-cyan-400 animate-pulse w-5 h-5 lg:w-6 lg:h-6" />
        <TrendingUp className="text-green-400 animate-bounce w-5 h-5 lg:w-6 lg:h-6" />
      </div>
      
      <p className="text-lg md:text-xl font-bold mb-2">
        GET TECH BROS TO CLICK YOUR BORING URL!
      </p>
      <p className="text-base md:text-lg opacity-90 max-w-2xl mx-auto mb-4 px-4">
        Transform your mundane links into <span className="text-cyan-400 font-bold">EXPONENTIALLY DISRUPTIVE</span> clickbait that leverages cutting-edge AI to drive engagement through the roof! 
      </p>
    </div>
  );
};

export default Header;
