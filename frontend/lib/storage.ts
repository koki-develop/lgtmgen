const keys = {
  favoriteIds: "FAVORITE_IDS",
} as const;

export const useStorage = () => {
  const loadFavorites = (): string[] => {
    const value = localStorage.getItem(keys.favoriteIds);
    if (value) {
      return JSON.parse(value);
    }
    return [];
  };

  const saveFavorites = (ids: string[]) => {
    localStorage.setItem(keys.favoriteIds, JSON.stringify(ids));
  };

  return {
    loadFavorites,
    saveFavorites,
  };
};
