import React, { useCallback, useState } from "react";
import { ModelsLGTM } from "@/lib/generated/api";
import clsx from "clsx";
import { useI18n } from "@/providers/I18nProvider";
import { useFetchLgtms } from "@/lib/models/lgtm/lgtmHooks";
import LgtmCardList from "./LgtmCardList";

export type LgtmPanelProps = {
  perPage: number;
  lgtms: ModelsLGTM[];
  favorites: string[];

  onLoaded: (lgtms: ModelsLGTM[]) => void;
  onChangeFavorites: (favorites: string[]) => void;
};

export default function LgtmPanel({
  perPage,
  lgtms,
  favorites,

  onLoaded,
  onChangeFavorites,
}: LgtmPanelProps) {
  const { t } = useI18n();
  const { fetchLgtms, fetching } = useFetchLgtms(perPage);

  const [hasNextPage, setHasNextPage] = useState<boolean>(
    lgtms.length === perPage,
  );

  const handleClickLoadMore = useCallback(async () => {
    const after = lgtms.slice(-1)[0]?.id;
    const loadedLgtms = await fetchLgtms(after);
    onLoaded(loadedLgtms);
    setHasNextPage(loadedLgtms.length === perPage);
  }, [fetchLgtms, lgtms, onLoaded, perPage]);

  return (
    <>
      <div className="flex flex-col gap-4">
        <LgtmCardList
          lgtmIds={lgtms.map((lgtm) => lgtm.id)}
          favorites={favorites}
          onChangeFavorites={onChangeFavorites}
        />

        <div className="flex justify-center">
          <button
            className={clsx(
              { hidden: !hasNextPage || fetching },
              "button-primary rounded px-4 py-2 shadow-md",
            )}
            onClick={handleClickLoadMore}
          >
            {t.loadMore}
          </button>

          <div className={clsx("loader", { hidden: !fetching })} />
        </div>
      </div>
    </>
  );
}
