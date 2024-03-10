import React from "react";
import { Navigate, useLocation, useParams } from "react-router-dom";

type Props = {
  component: React.ReactNode;
};

const RouteAuthGuard = ({ component }: Props) => {
  const location = useLocation();
  const authUsername = sessionStorage.getItem("authUsername");
  if (location.state?.from === "/signup") {
    return <>{component}</>;
  }
  if (!authUsername) {
    alert("コンテンツの閲覧にはログインが必要です");
    return <Navigate to="/signin" />;
  }
  const { username } = useParams();
  if (authUsername && authUsername === username) {
    return <>{component}</>;
  } else {
    alert("コンテンツの閲覧にはログインが必要です");
    return <Navigate to="/signin" />;
  }
};

export default RouteAuthGuard;
