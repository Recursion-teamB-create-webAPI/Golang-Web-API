import axios, { HttpStatusCode } from "axios";
import { jwtDecode } from "jwt-decode";
import { ChangeEvent, FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";

type Props = {
  username: string;
  token: string;
  status: number;
  error: string;
};

const SignIn = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  sessionStorage.removeItem("authUsername");

  const handleSignIn = async (e: FormEvent<HTMLButtonElement>) => {
    e.preventDefault();
    if (username === "" || password === "") {
      alert("ユーザー名とパスワードを入力してください");
      return;
    }
    const resp = await axios.post("/api/signin", {
      username: username,
      password: password,
    });

    const data: Props = resp.data;
    if (data.status === HttpStatusCode.Ok) {
      const jwtToken: JwtPayload = jwtDecode(data.token);
      sessionStorage.setItem("authUsername", jwtToken.username);
      navigate(`/${jwtToken.username}`);
    } else if (data.error.toLowerCase().includes("password")) {
      alert("パスワードが違います。");
      return;
    } else {
      alert("そのようなユーザは存在しません。");
      return;
    }
  };

  const handleGoSignUpPage = () => {
    navigate("/signup");
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
        <p className="text-2xl text-center my-5 text-blue-400">Sign In Form</p>
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
        <div className="flex justify-center items-center mt-4">
          <button
            className="bg-blue-400 hover:bg-blue-600 text-white rounded-md p-3 text-lg"
            onClick={handleSignIn}
          >
            Sign In
          </button>
          <button
            className="text-blue-500 rounded-xl p-1 text-sm"
            onClick={handleGoSignUpPage}
          >
            アカウントを作る
          </button>
        </div>
      </form>
    </div>
  );
};

export default SignIn;
