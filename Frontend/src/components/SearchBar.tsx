import { MagnifyingGlassIcon } from "@heroicons/react/20/solid";
import { useSearchResultState } from "../store/SearchResultStore";
import { useSearchState } from "../store/SearchStore";

const SearchBar = () => {
  const getSearchResultState = useSearchResultState(
    (state) => state.getSearchResultState
  );
  const [searchString, setSearchString] = useSearchState((state) => [
    state.searchString,
    state.setSearchString,
  ]);

  const handleSearch = async () => {
    if (searchString === "") return;
    getSearchResultState({ keyword: searchString });
  };

  const handleSearchString = (e) => {
    setSearchString(e.target.value);
  };

  return (
    <>
      <div className="flex items-center justify-between mt-10 px-4">
        <div className="flex items-center bg-white w-full h-full p-2">
          <MagnifyingGlassIcon className="h-10 w-10 text-gray-400" />
          <input
            className="h-10 w-full ml-2 outline-none rounded-md"
            placeholder="Search"
            onChange={handleSearchString}
          />
        </div>
        <button
          className="bg-[#0055D1] px-3 py-3 ml-2 
          text-white rounded-xl hover:bg-blue-400"
          onClick={handleSearch}
        >
          Search
        </button>
      </div>
    </>
  );
};

export default SearchBar;
