import { create } from "zustand"

interface SearchResultState {
  searchResults: SearchResult[]
  setSearchResultState: (results: SearchResult[]) => void
}

export const useSearchResultState = create<SearchResultState>((set) => ({
  searchResults: [],
  setSearchResultState: (searchResults: SearchResult[]) => set({ searchResults }),
}))
