import { useI18n } from "@/providers/I18nProvider";
import { useState } from "react";
import LgtmCardList from "./LgtmCardList";

export type FavoritePanelProps = {
  favorites: string[];
  onChangeFavorites: (favorites: string[]) => void;
};

export default function FavoritePanel({
  favorites,
  onChangeFavorites,
}: FavoritePanelProps) {
  const { t } = useI18n();

  const [currentFavorites, _] = useState<string[]>(favorites); // clear favorites when unmount

  if (currentFavorites.length === 0) {
    return <p className="text-center text-gray-500">{t.noFavorites}</p>;
  }

  return (
    <LgtmCardList
      lgtmIds={currentFavorites}
      favorites={favorites}
      onChangeFavorites={onChangeFavorites}
    />
  );
}
