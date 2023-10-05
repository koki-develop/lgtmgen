import React, { useCallback, useState } from "react";
import { ModelsLGTM } from "@/lib/generated/api";
import clsx from "clsx";
import { useI18n } from "@/providers/I18nProvider";
import { useFetchLgtms } from "@/lib/models/lgtm/lgtmHooks";
import LgtmCardList from "./LgtmCardList";
import { useStorage } from "@/lib/storage";
import { Switch } from "@headlessui/react";

export type LgtmPanelProps = {
  perPage: number;
  hasNextPage: boolean;
  randomly: boolean;
  lgtms: ModelsLGTM[];
  favorites: string[];

  onLoaded: (lgtms: ModelsLGTM[]) => void;
  onClear: () => void;
  onChangeRandomly: (randomly: boolean) => void;
  onChangeFavorites: (favorites: string[]) => void;
};

export default function LgtmPanel({
  perPage,
  hasNextPage,
  randomly,
  lgtms,
  favorites,

  onLoaded,
  onClear,
  onChangeRandomly,
  onChangeFavorites,
}: LgtmPanelProps) {
  const { t } = useI18n();
  const { fetchLgtms, fetching } = useFetchLgtms(perPage);
  const { saveRandomly } = useStorage();

  const handleClickLoadMore = useCallback(async () => {
    const after = lgtms.slice(-1)[0]?.id;
    const loadedLgtms = await fetchLgtms({ after });
    onLoaded(loadedLgtms);
  }, [fetchLgtms, lgtms, onLoaded]);

  const handleChangeRandomly = useCallback(
    async (randomly: boolean) => {
      onChangeRandomly(randomly);
      saveRandomly(randomly);
      onClear();

      if (randomly) {
        const loadedLgtms = await fetchLgtms({ random: true });
        onLoaded(loadedLgtms);
      } else {
        const loadedLgtms = await fetchLgtms({});
        onLoaded(loadedLgtms);
      }
    },
    [onChangeRandomly, saveRandomly, fetchLgtms, onLoaded, onClear],
  );

  return (
    <>
      <div className="flex flex-col gap-4">
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
          {!randomly && (
            <button
              className={clsx(
                { hidden: !hasNextPage || fetching },
                "button-primary rounded px-4 py-2 shadow-md",
              )}
              onClick={handleClickLoadMore}
            >
              {t.loadMore}
            </button>
          )}

          <div className={clsx("loader", { hidden: !fetching })} />
        </div>
      </div>
    </>
  );
}
