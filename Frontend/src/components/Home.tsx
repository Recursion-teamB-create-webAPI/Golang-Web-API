import { useSearchResultState } from '../store/SearchResultStore'
import SearchResultCard from './SearchResultCard'

const Home = () => {

  const searchResults = useSearchResultState((state) => state.searchResults)

  return (
    <>
      {/* This should be user name. After implementeing of auth, change this. */}
      <h1 className='mt-3'>
        Hello TlexCypher
      </h1>
      {searchResults ? (
        <p className='mt-3 mx-auto font-bold text-xl'>検索結果</p>
      ) : (
        <p className='mt-3 mx-auto font-bold text-xl'>検索結果なし</p>
      )}
      <div className='mx-auto'>
        {searchResults && searchResults.map((imageURL) =>
          <SearchResultCard
            imageURL={imageURL}
          />
        )}
        {/* This should be vanished. */}
        {searchResults && console.log("search results>>", searchResults)}
      </div>
    </>
  )
}

export default Home
