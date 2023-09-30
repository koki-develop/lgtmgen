import React from "react";
import { ModelsLGTM } from "@/lib/generated/api";
import { lgtmUrl } from "@/lib/image";
import ImageCard from "./ImageCard";
import clsx from "clsx";
import { useI18n } from "@/providers/I18nProvider";
import { useFetchLgtms } from "@/lib/models/lgtm/lgtmHooks";
import ImageCardButtons from "./ImageCardButtons";

export type LgtmPanelProps = {
  perPage: number;
  lgtms: ModelsLGTM[];
  favorites: string[];

  onLoaded: (lgtms: ModelsLGTM[]) => void;
  onFavorite: (id: string) => void;
  onUnfavorite: (id: string) => void;
};

export default function LgtmPanel({
  perPage,
  lgtms,
  favorites,

  onLoaded,
  onFavorite,
  onUnfavorite,
}: LgtmPanelProps) {
  const { t } = useI18n();
  const { fetchLgtms, fetching } = useFetchLgtms(perPage);

  const [hasNextPage, setHasNextPage] = React.useState<boolean>(
    lgtms.length === perPage,
  );

  const onLoadMore = React.useCallback(async () => {
    const after = lgtms.slice(-1)[0]?.id;
    const loadedLgtms = await fetchLgtms(after);
    onLoaded(loadedLgtms);
    setHasNextPage(loadedLgtms.length === perPage);
  }, [fetchLgtms, lgtms, onLoaded, perPage]);

  return (
    <div className="flex flex-col gap-4">
      <ul className="grid grid-cols-4 gap-4">
        {lgtms.map((lgtm) => (
          <li key={lgtm.id}>
            <ImageCard className="h-full" src={lgtmUrl(lgtm.id)} alt="LGTM">
              <ImageCardButtons
                lgtmId={lgtm.id}
                favorited={favorites.includes(lgtm.id)}
                onFavorite={onFavorite}
                onUnfavorite={onUnfavorite}
              />
            </ImageCard>
          </li>
        ))}
      </ul>

      <div className="flex justify-center">
        <button
          className={clsx(
            { hidden: !hasNextPage || fetching },
            "button-primary rounded px-4 py-2 shadow-md",
          )}
          onClick={onLoadMore}
        >
          {t.loadMore}
        </button>

        <div className={clsx("loader", { hidden: !fetching })} />
      </div>
    </div>
  );
}
