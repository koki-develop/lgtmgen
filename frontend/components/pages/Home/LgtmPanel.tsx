import React from "react";
import { ModelsLGTM } from "@/lib/generated/api";
import { lgtmUrl } from "@/lib/image";
import { LgtmUploader } from "./LgtmUploader";
import ImageCard from "./ImageCard";

export type LgtmPanelProps = {
  lgtms: ModelsLGTM[];
  hasNextPage: boolean;

  onLoadMore: () => void;
  onUploaded: (lgtm: ModelsLGTM) => void;
};

export default function LgtmPanel({
  lgtms,
  hasNextPage,
  onLoadMore,
  onUploaded,
}: LgtmPanelProps) {
  return (
    <div>
      <LgtmUploader onUploaded={onUploaded} />
      <ul className="grid grid-cols-4 gap-4">
        {lgtms.map((lgtm) => (
          <li key={lgtm.id}>
            <ImageCard className="h-full" src={lgtmUrl(lgtm.id)} alt="LGTM" />
          </li>
        ))}
      </ul>
      {hasNextPage && <button onClick={onLoadMore}>Load more</button>}
    </div>
  );
}
