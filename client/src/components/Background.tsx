const Background = () => {
  return (
    <div className="absolute inset-0 overflow-hidden">
      <div className="absolute -top-20 lg:-top-40 -right-20 lg:-right-40 w-40 h-40 lg:w-80 lg:h-80 bg-purple-500 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-pulse"></div>
      <div className="absolute -bottom-20 lg:-bottom-40 -left-20 lg:-left-40 w-40 h-40 lg:w-80 lg:h-80 bg-blue-500 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-pulse animation-delay-2000"></div>
      <div className="absolute top-20 lg:top-40 left-1/2 w-40 h-40 lg:w-80 lg:h-80 bg-indigo-500 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-pulse animation-delay-4000"></div>
    </div>
  );
};

export default Background;
