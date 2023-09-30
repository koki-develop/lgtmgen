const keys = {
  favoriteIds: "FAVORITE_IDS",
} as const;

export const useStorage = () => {
  const loadFavorites = (): string[] => {
    if (typeof window === "undefined") return [];

    const value = window.localStorage.getItem(keys.favoriteIds);
    if (value) {
      return JSON.parse(value);
    }
    return [];
  };

  const saveFavorites = (ids: string[]) => {
    if (typeof window === "undefined") return;

    window.localStorage.setItem(keys.favoriteIds, JSON.stringify(ids));
  };

  return {
    loadFavorites,
    saveFavorites,
  };
};
