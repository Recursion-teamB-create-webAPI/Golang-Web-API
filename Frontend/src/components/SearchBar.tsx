import { MagnifyingGlassIcon } from '@heroicons/react/20/solid'
import axios from 'axios'
import React from 'react'

const SearchBar = () => {
  const handleSearch = async () => {
    const { data } = await axios.get("/api/search")
    /* add logging.*/
    console.log(data)
  }
  return (
    <>
      <div className='flex items-center justify-between bg-white shadow-md rounded-md p-2 m-5'>
        <MagnifyingGlassIcon
          className='h-10 w-10 text-gray-400' />
        <input
          className='h-10 w-full ml-2 outline-none'
          placeholder='Search'
          onChange={handleSearch}
        />
      </div>
    </>
  )
}

export default SearchBar
