import axios from "axios";
import { FormEvent } from "react";
import { useNavigate } from "react-router-dom";

const Signout = () => {
  const navigate = useNavigate();
  const handleSignout = async (e: FormEvent<HTMLButtonElement>) => {
    e.preventDefault();
    await axios.post("/api/signout", {
      msg: "signout",
    });
    navigate("/signin");
  };

  return (
    <button
      className="bg-[#0055D1] rounded mx-3 px-2 py-2 shadow-xl text-white hover:bg-blue-400"
      onClick={handleSignout}
    >
      Sign Out
    </button>
  );
};

export default Signout;
