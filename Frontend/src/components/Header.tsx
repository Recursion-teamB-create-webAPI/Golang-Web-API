import { Avatar } from '@mui/material'
import React from 'react'
import Logout from './Logout'
import { pink } from '@mui/material/colors'
import SearchBar from './SearchBar'

const Header = () => {
  return (
    <>
      <header className='mt-3'>
        <div
          className=
          'flex justify-between items-center'
        >
          <h1
            className='text-4xl text-blue-500'
          >
            Golang Web API
          </h1>
          <div
            className='flex items-center'
          >
            <Avatar
              sx={{ bgcolor: pink[300] }}
            />
            <Logout />
          </div>
        </div>
        <SearchBar />
      </header>
    </>
  )
}

export default Header
