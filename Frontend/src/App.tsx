import { BrowserRouter, Route, Routes } from "react-router-dom"
import Home from "./components/Home"
import Layout from "./components/Layout"
import axios from "axios"
import SignIn from "./components/SignIn"
import SignUp from "./components/SignUp"

function App() {

  axios.defaults.baseURL = "http://localhost:8000"

  return (
    <>
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/signin" element={<SignIn />} />
            <Route path="/signup" element={<SignUp />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </>
  )
}

export default App
