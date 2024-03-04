import { BrowserRouter, Route, Routes } from "react-router-dom"
import Home from "./components/Home"
import Layout from "./components/Layout"
import axios from "axios"

function App() {

  axios.defaults.baseURL = "http://localhost:8000"

  return (
    <>
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/" element={<Home />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </>
  )
}

export default App
