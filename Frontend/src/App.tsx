import { BrowserRouter, Route, Routes } from "react-router-dom"
import Home from "./components/Home"
import Layout from "./components/Layout"

function App() {

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
