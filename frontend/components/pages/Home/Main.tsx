"use client";

import { ModelsCategory, ModelsLGTM, ModelsNews } from "@/lib/generated/api";
// import { i18n } from "@/lib/i18n";
// import { useStorage } from "@/lib/storage";
// import { Tab } from "@headlessui/react";
// import clsx from "clsx";
// import { useCallback, useEffect, useState } from "react";
// import FavoritePanel from "./FavoritePanel";
// import LgtmPanel from "./LgtmPanel";
// import SearchImagePanel from "./SearchImagePanel";

export type MainProps = {
  locale: string;
  newsList: ModelsNews[];
  initialData: ModelsLGTM[];
  initialRandomData: ModelsLGTM[];
  categories: ModelsCategory[];
  perPage: number;
};

export default function Main({
  // locale,
  newsList,
  // initialData,
  // initialRandomData,
  // categories,
  // perPage,
}: MainProps) {
  // const { loadFavorites, loadRandomly } = useStorage();
  // const t = i18n(locale);
  // const [rendered, setRendered] = useState<boolean>(false);

  /*
   * LGTM
   */

  // const [lgtms, setLgtms] = useState<ModelsLGTM[]>([]);
  // const [hasNextPage, setHasNextPage] = useState<boolean>(false);
  // const [randomly, setRandomly] = useState<boolean>(false);
  // const [selectedCategoryName, setSelectedCategoryName] = useState<
  //   string | null
  // >(null);

  // const handleLoaded = useCallback(
  //   (loadedLgtms: ModelsLGTM[]) => {
  //     setLgtms((prev) => [...prev, ...loadedLgtms]);
  //     setHasNextPage(loadedLgtms.length === perPage);
  //   },
  //   [perPage],
  // );

  // const handleClear = useCallback(() => {
  //   setLgtms([]);
  // }, []);

  // const handleChangeRandomly = useCallback((randomly: boolean) => {
  //   setRandomly(randomly);
  // }, []);

  // const handleChangeCategory = useCallback(
  //   (category: ModelsCategory | null) => {
  //     setSelectedCategoryName(category?.name ?? null);
  //   },
  //   [],
  // );

  /*
   * SearchImage
   */

  /*
   * Favorite
   */

  // const [favorites, setFavorites] = useState<string[]>([]);

  // useEffect(() => {
  //   if (rendered) return;

  //   const favorites = loadFavorites();
  //   setFavorites(favorites);

  //   const randomly = loadRandomly();
  //   setRandomly(randomly);

  //   if (randomly) {
  //     setLgtms(initialRandomData);
  //     setHasNextPage(initialRandomData.length === perPage);
  //   } else {
  //     setLgtms(initialData);
  //     setHasNextPage(initialData.length === perPage);
  //   }

  //   setRendered(true);
  // }, [
  //   rendered,
  //   loadFavorites,
  //   loadRandomly,
  //   initialData,
  //   initialRandomData,
  //   perPage,
  // ]);

  // const handleChangeFavorites = useCallback((favorites: string[]) => {
  //   setFavorites(favorites);
  // }, []);

  // Render
  return (
    <div className="flex flex-col">
      {newsList.length > 0 && (
        <div className="mb-4 flex flex-col gap-2">
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
      )}

      {/* <Tab.Group> */}
      {/*   <Tab.List className="mb-4 flex rounded-t bg-white shadow-md"> */}
      {/*     {[t.lgtm, t.searchImage, t.favorite].map((label) => ( */}
      {/*       <Tab */}
      {/*         key={label} */}
      {/*         className={clsx( */}
      {/*           "text-sm sm:text-base", */}
      {/*           "flex-grow py-3 outline-none transition sm:py-4", */}
      {/*           "border-b-2 border-b-white", */}
      {/*           "hover:border-b-gray-100 hover:bg-gray-100", */}
      {/*           "ui-selected:border-b-primary-main ui-selected:font-semibold ui-selected:text-primary-main ui-not-selected:text-gray-400", */}
      {/*         )} */}
      {/*       > */}
      {/*         {label} */}
      {/*       </Tab> */}
      {/*     ))} */}
      {/*   </Tab.List> */}

      {/*   <Tab.Panels> */}
      {/*     <Tab.Panel> */}
      {/*       {rendered ? ( */}
      {/*         <LgtmPanel */}
      {/*           lgtms={lgtms} */}
      {/*           categories={categories} */}
      {/*           selectedCategoryName={selectedCategoryName} */}
      {/*           randomly={randomly} */}
      {/*           onChangeRandomly={handleChangeRandomly} */}
      {/*           favorites={favorites} */}
      {/*           perPage={perPage} */}
      {/*           hasNextPage={hasNextPage} */}
      {/*           onLoaded={handleLoaded} */}
      {/*           onClear={handleClear} */}
      {/*           onChangeFavorites={handleChangeFavorites} */}
      {/*           onChangeCategory={handleChangeCategory} */}
      {/*         /> */}
      {/*       ) : ( */}
      {/*         <div className="flex justify-center"> */}
      {/*           <div className="loader" /> */}
      {/*         </div> */}
      {/*       )} */}
      {/*     </Tab.Panel> */}

      {/*     <Tab.Panel> */}
      {/*       <SearchImagePanel /> */}
      {/*     </Tab.Panel> */}

      {/*     <Tab.Panel> */}
      {/*       {rendered ? ( */}
      {/*         <FavoritePanel */}
      {/*           favorites={favorites} */}
      {/*           onChangeFavorites={handleChangeFavorites} */}
      {/*         /> */}
      {/*       ) : ( */}
      {/*         <div className="flex justify-center"> */}
      {/*           <div className="loader" /> */}
      {/*         </div> */}
      {/*       )} */}
      {/*     </Tab.Panel> */}
      {/*   </Tab.Panels> */}
      {/* </Tab.Group> */}
    </div>
  );
}
