import React from "react";
import LgtmCardList from "./LgtmCardList";
import { useI18n } from "@/providers/I18nProvider";

export type FavoritePanelProps = {
  favorites: string[];
  onChangeFavorites: (favorites: string[]) => void;
};

export default function FavoritePanel({
  favorites,
  onChangeFavorites,
}: FavoritePanelProps) {
  const { t } = useI18n();

  if (favorites.length === 0) {
    return <p className="text-center text-gray-500">{t.noFavorites}</p>;
  }

  return (
    <LgtmCardList
      lgtmIds={favorites}
      favorites={favorites}
      onChangeFavorites={onChangeFavorites}
    />
  );
}
