import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./components/Home";
import Layout from "./components/Layout";
import axios from "axios";
import SignIn from "./components/SignIn";
import SignUp from "./components/SignUp";
import Description from "./components/Description";
import RouteAuthGuard from "./components/RouteAuthGuard";

function App() {
  axios.defaults.baseURL = "http://localhost:8000";

  return (
    <>
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/signin" element={<SignIn />} />
            <Route path="/signup" element={<SignUp />} />
            <Route path="/:username">
              <Route
                path=""
                element={<RouteAuthGuard component={<Home />} />}
              />
              <Route
                path="description/:imageURL"
                element={<RouteAuthGuard component={<Description />} />}
              />
            </Route>
          </Routes>
        </Layout>
      </BrowserRouter>
    </>
  );
}

export default App;
