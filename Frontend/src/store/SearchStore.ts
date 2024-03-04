import axios from "axios";
import { create } from "zustand"

interface SearchState {
  searchString: string;
  setSearchString: (searchString: string) => void;
}
export const useSearchState = create<SearchState>((set) => ({
  searchString: "",
  setSearchString: (searchString: string) => {
    set({ searchString })
  },
}))
