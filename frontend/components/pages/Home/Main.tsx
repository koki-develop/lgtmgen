"use client";

import React, { useCallback, useEffect, useState } from "react";
import { ModelsImage, ModelsLGTM, ModelsNews } from "@/lib/generated/api";
import LgtmPanel from "./LgtmPanel";
import SearchImagePanel from "./SearchImagePanel";
import { Tab } from "@headlessui/react";
import { i18n } from "@/lib/i18n";
import clsx from "clsx";
import { LgtmUploader } from "./LgtmUploader";
import { useStorage } from "@/lib/storage";
import FavoritePanel from "./FavoritePanel";

export type MainProps = {
  locale: string;
  newsList: ModelsNews[];
  initialData: ModelsLGTM[];
  initialRandomData: ModelsLGTM[];
  perPage: number;
};

export default function Main({
  locale,
  newsList,
  initialData,
  initialRandomData,
  perPage,
}: MainProps) {
  const { loadFavorites, loadRandomly } = useStorage();
  const t = i18n(locale);
  const [rendered, setRendered] = useState<boolean>(false);

  /*
   * LGTM
   */

  const [lgtms, setLgtms] = useState<ModelsLGTM[]>([]);
  const [hasNextPage, setHasNextPage] = useState<boolean>(false);
  const [randomly, setRandomly] = useState<boolean>(false);

  const handleLoaded = useCallback(
    (loadedLgtms: ModelsLGTM[]) => {
      setLgtms((prev) => [...prev, ...loadedLgtms]);
      setHasNextPage(loadedLgtms.length === perPage);
    },
    [perPage],
  );

  const handleClear = useCallback(() => {
    setLgtms([]);
  }, []);

  const handleGenerated = useCallback((lgtm: ModelsLGTM) => {
    setLgtms((prev) => [lgtm, ...prev]);
  }, []);

  const handleChangeRandomly = useCallback((randomly: boolean) => {
    setRandomly(randomly);
  }, []);

  /*
   * SearchImage
   */

  const [images, setImages] = useState<ModelsImage[]>([]);
  const [query, setQuery] = useState<string>("");

  const handleChangeQuery = useCallback((query: string) => {
    if (query.length > 255) return;
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
    if (rendered) return;

    const favorites = loadFavorites();
    setFavorites(favorites);

    const randomly = loadRandomly();
    setRandomly(randomly);

    if (randomly) {
      setLgtms(initialRandomData);
      setHasNextPage(initialRandomData.length === perPage);
    } else {
      setLgtms(initialData);
      setHasNextPage(initialData.length === perPage);
    }

    setRendered(true);
  }, [
    rendered,
    loadFavorites,
    loadRandomly,
    initialData,
    initialRandomData,
    perPage,
  ]);

  const handleChangeFavorites = useCallback((favorites: string[]) => {
    setFavorites(favorites);
  }, []);

  // Render
  return (
    <div className="flex flex-col">
      <div className="flex flex-col gap-2">
        {newsList.map((news, i) => (
          <div
            key={i}
            className="flex flex-col gap-1 rounded border border-primary-main bg-blue-100 p-2 text-primary-dark shadow-md"
          >
            <div>
              {news.title && (
                <div className="text-lg font-bold">{news.title}</div>
              )}

              {news.date && <div className="mb-1 text-xs">{news.date}</div>}
            </div>

            {news.content && (
              <div
                className="whitespace-pre-wrap text-sm"
                dangerouslySetInnerHTML={{
                  __html: news.content,
                }}
              />
            )}
          </div>
        ))}
      </div>

      <Tab.Group>
        <Tab.List className="mb-4 flex rounded-t bg-white shadow-md">
          {[t.lgtm, t.searchImage, t.favorite].map((label) => (
            <Tab
              key={label}
              className={clsx(
                "text-sm sm:text-base",
                "flex-grow py-3 outline-none transition sm:py-4",
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
          <LgtmUploader onUploaded={handleGenerated} />
          <Tab.Panel>
            {rendered ? (
              <LgtmPanel
                lgtms={lgtms}
                randomly={randomly}
                onChangeRandomly={handleChangeRandomly}
                favorites={favorites}
                perPage={perPage}
                hasNextPage={hasNextPage}
                onLoaded={handleLoaded}
                onClear={handleClear}
                onChangeFavorites={handleChangeFavorites}
              />
            ) : (
              <div className="flex justify-center">
                <div className="loader" />
              </div>
            )}
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

          <Tab.Panel>
            {rendered ? (
              <FavoritePanel
                favorites={favorites}
                onChangeFavorites={handleChangeFavorites}
              />
            ) : (
              <div className="flex justify-center">
                <div className="loader" />
              </div>
            )}
          </Tab.Panel>
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
}
