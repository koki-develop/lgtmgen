import { ModelsCategory, ModelsLGTM } from "@/lib/generated/api";
import { useFetchLgtms } from "@/lib/models/lgtm/lgtmHooks";
import { useStorage } from "@/lib/storage";
import { useI18n } from "@/providers/I18nProvider";
import { Switch } from "@headlessui/react";
import clsx from "clsx";
import { useCallback } from "react";
import LgtmCardList from "./LgtmCardList";

export type LgtmPanelProps = {
  perPage: number;
  hasNextPage: boolean;
  randomly: boolean;
  lgtms: ModelsLGTM[];
  categories: ModelsCategory[];
  selectedCategoryName: string | null;
  favorites: string[];

  onLoaded: (lgtms: ModelsLGTM[]) => void;
  onClear: () => void;
  onChangeRandomly: (randomly: boolean) => void;
  onChangeFavorites: (favorites: string[]) => void;
  onChangeCategory: (category: ModelsCategory | null) => void;
};

export default function LgtmPanel({
  perPage,
  hasNextPage,
  randomly,
  lgtms,
  categories,
  selectedCategoryName,
  favorites,

  onLoaded,
  onClear,
  onChangeRandomly,
  onChangeFavorites,
  onChangeCategory,
}: LgtmPanelProps) {
  const { t } = useI18n();
  const { fetchLgtms, fetching } = useFetchLgtms(perPage);
  const { saveRandomly } = useStorage();

  const handleClickLoadMore = useCallback(async () => {
    const after = lgtms.slice(-1)[0]?.id;
    const loadedLgtms = await fetchLgtms({
      after,
      category: selectedCategoryName ?? undefined,
    });
    onLoaded(loadedLgtms);
  }, [fetchLgtms, lgtms, onLoaded, selectedCategoryName]);

  const handleClickReload = useCallback(async () => {
    onClear();
    const loadedLgtms = await fetchLgtms({
      random: true,
      category: selectedCategoryName ?? undefined,
    });
    onLoaded(loadedLgtms);
  }, [selectedCategoryName, fetchLgtms, onLoaded, onClear]);

  const handleChangeRandomly = useCallback(
    async (randomly: boolean) => {
      onChangeRandomly(randomly);
      saveRandomly(randomly);
      onClear();

      const category = selectedCategoryName ?? undefined;

      const loadedLgtms = await fetchLgtms({ random: randomly, category });
      onLoaded(loadedLgtms);
    },
    [
      onChangeRandomly,
      saveRandomly,
      fetchLgtms,
      onLoaded,
      onClear,
      selectedCategoryName,
    ],
  );

  const handleClickCategory = useCallback(
    async (category: ModelsCategory) => {
      onClear();

      if (category.name === selectedCategoryName) {
        onChangeCategory(null);
        const loadedLgtms = await fetchLgtms({ random: randomly });
        onLoaded(loadedLgtms);
      } else {
        onChangeCategory(category);
        const loadedLgtms = await fetchLgtms({
          random: randomly,
          category: category.name,
        });
        onLoaded(loadedLgtms);
      }
    },
    [
      randomly,
      onClear,
      fetchLgtms,
      onLoaded,
      onChangeCategory,
      selectedCategoryName,
    ],
  );

  return (
    <>
      <div className="flex flex-col gap-4">
        <div className="flex flex-wrap gap-2">
          {categories.map((category) => (
            <button
              key={category.name}
              className={clsx(
                "rounded-full border border-primary-main px-2 py-1 text-sm",
                {
                  "bg-white text-primary-main":
                    category.name !== selectedCategoryName,
                  "bg-primary-main text-white":
                    category.name === selectedCategoryName,
                },
              )}
              onClick={() => handleClickCategory(category)}
            >
              {category.name}
            </button>
          ))}
        </div>
        <div className="flex justify-end">
          <Switch
            className="flex items-center justify-end gap-2"
            checked={randomly}
            onChange={handleChangeRandomly}
          >
            <span
              className={clsx(
                "relative flex items-center rounded-full border shadow-inner transition",
                "h-[28px] w-[56px]",
                "sm:h-[32px] sm:w-[60px]",
                {
                  "bg-primary-main": randomly,
                  "bg-gray-300": !randomly,
                },
              )}
            >
              <span
                className={clsx(
                  "absolute inline-block rounded-full bg-white transition",
                  "h-[20px] w-[20px]",
                  "sm:h-[24px] sm:w-[24px]",
                  {
                    "translate-x-[4px]": !randomly,
                    "translate-x-[30px]": randomly,
                  },
                )}
              />
            </span>
            <span className="text-sm sm:text-base">{t.random}</span>
          </Switch>
        </div>

        <LgtmCardList
          lgtmIds={lgtms.map((lgtm) => lgtm.id)}
          favorites={favorites}
          onChangeFavorites={onChangeFavorites}
        />

        <div className="flex justify-center">
          <button
            className={clsx(
              { hidden: (!hasNextPage && !randomly) || fetching },
              "button-primary rounded px-4 py-2 shadow-md",
            )}
            onClick={randomly ? handleClickReload : handleClickLoadMore}
          >
            {randomly ? t.reload : t.loadMore}
          </button>

          <div className={clsx("loader", { hidden: !fetching })} />
        </div>
      </div>
    </>
  );
}
