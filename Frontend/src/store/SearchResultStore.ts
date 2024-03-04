import axios from "axios";
import { create } from "zustand"

interface SearchResultState {
  searchResults: SearchResult[]
  setSearchResultState: (results: SearchResult[]) => void
  getSearchResultState: (params: SearchParams) => void;
}
export const useSearchResultState = create<SearchResultState>((set) => ({
  searchResults: [],
  setSearchResultState: (searchResults: SearchResult[]) => set({ searchResults }),
  getSearchResultState: async (params: SearchParams) => {
    const { data } = await axios.get("/api/search", { params });
    console.log("data>>", data);
    set({ searchResults: data.images })
  }
}))
