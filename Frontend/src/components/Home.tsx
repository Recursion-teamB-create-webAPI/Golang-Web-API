import { useSearchResultState } from "../store/SearchResultStore";
import SearchBar from "./SearchBar";
import SearchResultCard from "./SearchResultCard";

const Home = () => {
  const searchResults = useSearchResultState((state) => state.searchResults);

  return (
    <>
      <SearchBar />
      {searchResults.length > 0 ? (
        <p className="mt-3 mx-auto font-bold text-xl text-blue-500">
          検索結果一覧
        </p>
      ) : (
        <p className="mt-3 mx-auto font-bold text-xl text-blue-500">
          検索結果なし
        </p>
      )}
      <div className="mx-auto">
        {searchResults.length > 0 && (
          <SearchResultCard
            imageURL={searchResults[0]}
            totalResults={searchResults}
          />
        )}
      </div>
    </>
  );
};

export default Home;
