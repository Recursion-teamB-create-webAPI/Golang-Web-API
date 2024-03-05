const SignUp = () => {
  return (
    <div className="flex justify-center mt-20 drop-shadow-xl">
      <form className="bg-white px-10 pb-10 rounded-xl w-80 h-80">
        <p className="text-2xl text-center my-5 text-blue-400">Sign Up Form</p>
        <div>
          <p className="text-lg"> Username </p>
          <input
            type="text"
            placeholder="sample"
            className="border rounded-md w-full outline-none px-3 h-10"
          />
        </div>
        <div className="mt-6">
          <p className="text-lg"> Password </p>
          <input
            type="password"
            placeholder="password"
            className="border rounded-md w-full outline-none px-3 h-10"
          />
        </div>
        <div className="flex justify-center items-center mt-4">
          <button className="bg-blue-400 hover:bg-blue-600 text-white rounded-md p-3 text-lg">
            Sign Up
          </button>
        </div>
      </form>
    </div>
  );
};

export default SignUp;
