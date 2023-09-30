import React from "react";
import { ModelsLGTM } from "@/lib/generated/api";
import { lgtmUrl } from "@/lib/image";
import { LgtmUploader } from "./LgtmUploader";
import ImageCard from "./ImageCard";
import clsx from "clsx";
import { useI18n } from "@/providers/I18nProvider";

export type LgtmPanelProps = {
  lgtms: ModelsLGTM[];
  hasNextPage: boolean;
  loading: boolean;

  onLoadMore: () => void;
  onUploaded: (lgtm: ModelsLGTM) => void;
};

export default function LgtmPanel({
  lgtms,
  hasNextPage,
  loading,

  onLoadMore,
  onUploaded,
}: LgtmPanelProps) {
  const { t } = useI18n();

  return (
    <>
      <LgtmUploader onUploaded={onUploaded} />

      <div className="flex flex-col gap-4">
        <ul className="grid grid-cols-4 gap-4">
          {lgtms.map((lgtm) => (
            <li key={lgtm.id}>
              <ImageCard className="h-full" src={lgtmUrl(lgtm.id)} alt="LGTM" />
            </li>
          ))}
        </ul>
        <div className="flex justify-center">
          <button
            className={clsx(
              {
                hidden: !hasNextPage || loading,
              },
              "button-primary rounded px-4 py-2 shadow-md",
            )}
            onClick={onLoadMore}
          >
            {t.loadMore}
          </button>
          <div
            className={clsx("loader", {
              hidden: !loading,
            })}
          />
        </div>
      </div>
    </>
  );
}
