import { useCallback } from "react";

const keys = {
  favoriteIds: "FAVORITE_IDS",
  randomly: "RANDOMLY",
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

  const loadRandomly = useCallback((): boolean => {
    if (typeof window === "undefined") return false;

    const value = window.localStorage.getItem(keys.randomly);
    return value === "true";
  }, []);

  const saveRandomly = useCallback((randomly: boolean) => {
    if (typeof window === "undefined") return;

    window.localStorage.setItem(keys.randomly, randomly ? "true" : "false");
  }, []);

  return {
    loadFavorites,
    saveFavorites,
    loadRandomly,
    saveRandomly,
  };
};
