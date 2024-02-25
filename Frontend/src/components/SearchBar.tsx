import { MagnifyingGlassIcon } from '@heroicons/react/20/solid'
import React from 'react'

const SearchBar = () => {
  const handleSearch = () => {

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
