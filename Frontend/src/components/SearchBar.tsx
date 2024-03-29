import { MagnifyingGlassIcon } from "@heroicons/react/20/solid";
import { useSearchResultState } from "../store/SearchResultStore";
import { useSearchState } from "../store/SearchStore";
import { Button, Menu, MenuButton, MenuItem, MenuList } from "@chakra-ui/react";
import { ChangeEvent, useEffect, useState } from "react";
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

  const handleSearch = () => {
    if (searchString === "") return;
    getSearchResultState({ keyword: searchString });
    handleSearchCandidate();
  };

  const handleSearchString = (e: ChangeEvent<HTMLInputElement>) => {
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
            color="black"
            fontWeight={"bold"}
          >
            検索候補の表示
          </MenuButton>
          <MenuList mt={10}>
            {searchCandidates.length > 0 &&
              searchCandidates.map((candidate, index) => (
                <MenuItem
                  key={index}
                  onClick={() => setSearchString(candidate.item)}
                >
                  <p className="bg-white font-bold min-w-96 text-lg px-4 py-2 mb-2 rounded-md flex justify-between">
                    <span className="text-[#0055d1] pr-5">{`検索用語: ${candidate.item}`}</span>
                    <span className="pr-5 text-pink-500">{`検索回数: ${candidate.search_count}`}</span>
                    <span className="text-emerald-500">{`最後に調べた日時: ${candidate.updated_at}`}</span>
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
