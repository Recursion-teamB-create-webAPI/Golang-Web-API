import { ReactNode } from "react";
import Header from "./Header";
import { useLocation } from "react-router-dom";

type Props = {
  children: ReactNode;
};

const Layout = ({ children }: Props) => {
  const location = useLocation();
  return (
    <>
      <div
        className="
          absolute
          top-0
          left-0
          w-full
          h-full
          bg-gradient-to-br
          from-pink-400
          to-[#0055D1]
          rounded-md
          filter
          blur-3xl
          opacity-50
          -z-50
          "
      />
      {location &&
        location.pathname !== "/signin" &&
        location.pathname != "/signup" && <Header />}
      {children}
    </>
  );
};

export default Layout;
