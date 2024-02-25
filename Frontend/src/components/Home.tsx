import SearchResultCard from './SearchResultCard'

const Home = () => {
  const searchResults: SearchResult[] = useSearchResultState((state) =>
    state.searchResults
  )
  return (
    <>
      {/* This should be user name. After implementeing of auth, change this. */}
      <h1>Hello User</h1>
      {searchResults ? (
        <p>検索結果はこちら!</p>
      ) : (
        <p>検索結果なし</p>
      )}
      <div className='flex items-center justify-center'>
        {searchResults && searchResults.map(({ id, searchWord, imageURL }) =>
          <SearchResultCard
            id={id}
            searchWord={searchWord}
            imageURL={imageURL}
          />
        )}
      </div>
    </>
  )
}

export default Home
