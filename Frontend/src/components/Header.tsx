import Signout from "./Signout";
import Avatar from "react-avatar";

const Header = () => {
  return (
    <>
      <header className="mt-3">
        <div className="flex justify-between items-center px-3">
          <h1 className="text-4xl text-[#0055d1]">Pix Finder</h1>
          <div className="flex items-center">
            <Avatar name="TlexCypher Swapman" round size="50" color="#0055D1" />
            <Signout />
          </div>
        </div>
      </header>
    </>
  );
};

export default Header;
