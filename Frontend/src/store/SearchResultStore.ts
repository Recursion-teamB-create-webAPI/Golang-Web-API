import axios from "axios";
import { create } from "zustand";

interface SearchResultState {
  searchResults: string[];
  setSearchResultState: (results: string[]) => void;
  getSearchResultState: (params: SearchParams) => void;
}
export const useSearchResultState = create<SearchResultState>((set) => ({
  searchResults: [],
  setSearchResultState: (searchResults: string[]) => set({ searchResults }),
  getSearchResultState: async (params: SearchParams) => {
    if (params.keyword === "") {
      set({ searchResults: [] });
    }
    const { data } = await axios.get("/api/search", { params });
    set({ searchResults: data.imageData.images });
  },
}));
