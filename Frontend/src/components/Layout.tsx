import { ReactNode } from 'react'
import Header from './Header'
import Footer from './Footer'

type Props = {
  children: ReactNode
}

const Layout = ({ children }: Props) => {
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
      <Header />
      {children}
      {/* <Footer /> */}
    </>
  )
}

export default Layout
