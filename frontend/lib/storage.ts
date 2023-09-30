import { useCallback } from "react";

const keys = {
  favoriteIds: "FAVORITE_IDS",
} as const;

export const useStorage = () => {
  const loadFavorites = useCallback((): string[] => {
    if (typeof window === "undefined") return [];

    const value = window.localStorage.getItem(keys.favoriteIds);
    if (value) {
      return JSON.parse(value);
    }
    return [];
  }, []);

  const saveFavorites = useCallback((ids: string[]) => {
    if (typeof window === "undefined") return;

    window.localStorage.setItem(keys.favoriteIds, JSON.stringify(ids));
  }, []);

  return {
    loadFavorites,
    saveFavorites,
  };
};
