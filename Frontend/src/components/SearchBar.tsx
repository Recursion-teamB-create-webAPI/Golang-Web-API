import { MagnifyingGlassIcon } from "@heroicons/react/20/solid";
import { useSearchResultState } from "../store/SearchResultStore";
import { useSearchState } from "../store/SearchStore";
import { Button, Menu, MenuButton, MenuItem, MenuList } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import axios from "axios";

type SearchCandidate = {
  item: string;
  search_count: number;
  updated_at: Date;
};

const SearchBar = () => {
  const getSearchResultState = useSearchResultState(
    (state) => state.getSearchResultState
  );
  const [searchString, setSearchString] = useSearchState((state) => [
    state.searchString,
    state.setSearchString,
  ]);

  const [searchCandidates, setSearchCandidates] = useState<SearchCandidate[]>(
    []
  );

  const handleSearch = async () => {
    if (searchString === "") return;
    getSearchResultState({ keyword: searchString });
  };

  const handleSearchString = (e) => {
    setSearchString(e.target.value);
  };

  const handleSearchCandidate = async () => {
    const resp = await axios.get("/api/total_result");
    const data: SearchCandidate[] = resp.data.totalResult;
    setSearchCandidates(data);
  };

  useEffect(() => {
    handleSearchCandidate();
  }, []);

  return (
    <>
      <div className="flex items-center justify-between mt-10 px-4">
        <div className="flex items-center bg-white w-full h-full p-2">
          <MagnifyingGlassIcon className="h-10 w-10 text-gray-400" />
          <input
            className="h-10 w-full ml-2 outline-none rounded-md"
            placeholder="Search"
            onChange={handleSearchString}
            value={searchString}
          />
        </div>
        <button
          className="bg-[#0055D1] px-3 py-3 ml-2 
          text-white rounded-xl hover:bg-blue-400"
          onClick={handleSearch}
        >
          Search
        </button>
        <Menu>
          <MenuButton
            as={Button}
            ml={4}
            p={4}
            color="black" // テキストの色を設定
            fontWeight={"bold"}
          >
            検索候補の表示
          </MenuButton>
          <MenuList>
            {searchCandidates.length > 0 &&
              searchCandidates.map((candidate, index) => (
                <MenuItem
                  key={index}
                  onClick={() => setSearchString(candidate.item)}
                >
                  <p className="bg-white text-blue-500 font-bold w-full text-lg px-4 py-2 mb-2 rounded-md">
                    {candidate.item}
                  </p>
                </MenuItem>
              ))}
          </MenuList>
        </Menu>
      </div>
    </>
  );
};

export default SearchBar;
