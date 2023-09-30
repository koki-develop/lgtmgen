const keys = {
  favoriteIds: "FAVORITE_IDS",
} as const;

export const useStorage = () => {
  const getFavoriteIds = (): string[] => {
    const value = localStorage.getItem(keys.favoriteIds);
    if (value) {
      return JSON.parse(value);
    }
    return [];
  };

  const setFavoriteIds = (ids: string[]) => {
    localStorage.setItem(keys.favoriteIds, JSON.stringify(ids));
  };

  return {
    getFavoriteIds,
    setFavoriteIds,
  };
};
