import { useSearchResultState } from "../store/SearchResultStore";
import SearchBar from "./SearchBar";
import SearchResultCard from "./SearchResultCard";
import { useSearchState } from "../store/SearchStore";
import { useEffect } from "react";

const Home = () => {
  const [searchResults, getSearchResultState] = useSearchResultState(
    (state) => [state.searchResults, state.getSearchResultState]
  );

  const [searchString] = useSearchState((state) => [state.searchString]);

  useEffect(() => {
    getSearchResultState({ keyword: searchString });
  }, []);

  return (
    <>
      <SearchBar />
      {searchResults.length > 0 ? (
        <p className="mt-3 mx-auto font-bold text-xl text-[#0055D1] px-4">
          検索結果一覧
        </p>
      ) : (
        <p className="mt-3 mx-auto font-bold text-xl text-[#0055D1] px-4">
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
