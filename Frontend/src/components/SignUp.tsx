import axios, { HttpStatusCode } from "axios";
import { ChangeEvent, FormEvent, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";

const SignUp = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const from = useLocation().pathname;
  const naviate = useNavigate();

  const handleSignUp = async (e: FormEvent<HTMLButtonElement>) => {
    e.preventDefault();
    const { data } = await axios.post("/api/signup", {
      username: username,
      password: password,
    });
    console.log("SignUp Response>>", data);
    if (data.status === HttpStatusCode.Ok) {
      naviate(`/${username}`, {
        state: { username: username, from: from },
      });
    } else {
      alert("ユーザ作成に失敗しました。すでにアカウントを有していませんか？");
    }
  };

  const handleGoSignInPage = () => {
    naviate("/signin");
  };

  const handleUsername = (e: ChangeEvent<HTMLInputElement>) => {
    setUsername(e.target.value);
  };

  const handlePassword = (e: ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  return (
    <div className="flex justify-center mt-20 drop-shadow-xl">
      <form className="bg-white px-10 pb-10 rounded-xl w-84 h-80">
        <p className="text-2xl text-center my-5 text-blue-400">Sign Up Form</p>
        <div>
          <p className="text-lg"> Username </p>
          <input
            type="text"
            placeholder="sample"
            className="border rounded-md w-full outline-none px-3 h-10"
            onChange={handleUsername}
          />
        </div>
        <div className="mt-6">
          <p className="text-lg"> Password </p>
          <input
            type="password"
            placeholder="password"
            className="border rounded-md w-full outline-none px-3 h-10"
            onChange={handlePassword}
          />
        </div>
        <div className="flex justify-center items-center mt-4 space-x-2">
          <button
            className="bg-blue-400 hover:bg-blue-600 text-white rounded-md p-3 text-lg"
            onClick={handleSignUp}
          >
            Sign Up
          </button>
          <button
            className="text-blue-500 rounded-xl p-1 text-sm"
            onClick={handleGoSignInPage}
          >
            すでにアカウントをお持ちですか？
          </button>
        </div>
      </form>
    </div>
  );
};

export default SignUp;
