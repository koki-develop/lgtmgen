import React from "react";
import LgtmCardList from "./LgtmCardList";

export type FavoritePanelProps = {
  favorites: string[];
  onChangeFavorites: (favorites: string[]) => void;
};

export default function FavoritePanel({
  favorites,
  onChangeFavorites,
}: FavoritePanelProps) {
  return (
    <LgtmCardList
      lgtmIds={favorites}
      favorites={favorites}
      onChangeFavorites={onChangeFavorites}
    />
  );
}
