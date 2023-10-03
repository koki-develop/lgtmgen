"use client";

import React, { useCallback, useEffect, useState } from "react";
import { api } from "@/lib/api";
import { ModelsImage, ModelsLGTM } from "@/lib/generated/api";
import LgtmPanel from "./LgtmPanel";
import SearchImagePanel from "./SearchImagePanel";
import { Tab } from "@headlessui/react";
import { i18n } from "@/lib/i18n";
import clsx from "clsx";
import { LgtmUploader } from "./LgtmUploader";
import { useStorage } from "@/lib/storage";

export type MainProps = {
  locale: string;
  initialData: ModelsLGTM[];
  perPage: number;
};

export default function Main({ locale, initialData, perPage }: MainProps) {
  const { loadFavorites, saveFavorites } = useStorage();
  const t = i18n(locale);

  /*
   * LGTM
   */

  const [lgtms, setLgtms] = useState<ModelsLGTM[]>(initialData);

  const handleLoaded = useCallback((loadedLgtms: ModelsLGTM[]) => {
    setLgtms((prev) => [...prev, ...loadedLgtms]);
  }, []);

  const handleGenerated = useCallback((lgtm: ModelsLGTM) => {
    setLgtms((prev) => [lgtm, ...prev]);
  }, []);

  /*
   * SearchImage
   */

  const [images, setImages] = useState<ModelsImage[]>([]);
  const [query, setQuery] = useState<string>("");

  const handleChangeQuery = useCallback((query: string) => {
    setQuery(query);
  }, []);

  const handleSearched = useCallback((images: ModelsImage[]) => {
    setImages(images);
  }, []);

  /*
   * Favorite
   */

  const [favorites, setFavorites] = useState<string[]>([]);

  useEffect(() => {
    const favorites = loadFavorites();
    setFavorites(favorites);
  }, [loadFavorites]);

  const handleFavorite = useCallback(
    (id: string) => {
      setFavorites((prev) => {
        const next = [...prev, id];
        saveFavorites(next);
        return next;
      });
    },
    [saveFavorites],
  );

  const handleUnfavorite = useCallback(
    (id: string) => {
      setFavorites((prev) => {
        const next = prev.filter((v) => v !== id);
        saveFavorites(next);
        return next;
      });
    },
    [saveFavorites],
  );

  // Render
  return (
    <div>
      <Tab.Group>
        <Tab.List className="mb-4 flex rounded-t bg-white shadow-md">
          {[t.lgtm, t.searchImage, t.favorite].map((label) => (
            <Tab
              key={label}
              className={clsx(
                "flex-grow py-4 outline-none transition",
                "border-b-2 border-b-white",
                "hover:border-b-gray-100 hover:bg-gray-100",
                "ui-selected:border-b-primary-main ui-selected:font-semibold ui-selected:text-primary-main ui-not-selected:text-gray-400",
              )}
            >
              {label}
            </Tab>
          ))}
        </Tab.List>

        <Tab.Panels>
          <Tab.Panel>
            <LgtmUploader onUploaded={handleGenerated} />
            <LgtmPanel
              lgtms={lgtms}
              favorites={favorites}
              perPage={perPage}
              onLoaded={handleLoaded}
              onFavorite={handleFavorite}
              onUnfavorite={handleUnfavorite}
            />
          </Tab.Panel>

          <Tab.Panel>
            <SearchImagePanel
              images={images}
              query={query}
              onChangeQuery={handleChangeQuery}
              onSearched={handleSearched}
              onGenerated={handleGenerated}
            />
          </Tab.Panel>
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
}
