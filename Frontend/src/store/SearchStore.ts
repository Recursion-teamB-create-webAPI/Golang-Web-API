import { create } from "zustand";

interface SearchState {
  searchString: string;
  searchReady: boolean;
  setSearchString: (searchString: string) => void;
  getSearchString: () => string;
  setSearchReady: (ready: boolean) => void;
}
export const useSearchState = create<SearchState>((set, get) => ({
  searchString: "",
  searchReady: false,
  setSearchString: (searchString: string) => {
    set({ searchString });
  },
  getSearchString: () => {
    return get().searchString;
  },
  setSearchReady: (searchReady: boolean) => {
    set({ searchReady });
  },
}));
