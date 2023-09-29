import React from "react";
import { ModelsLGTM } from "@/lib/generated/api";
import { lgtmUrl } from "@/lib/image";
import { LgtmUploader } from "./LgtmUploader";

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
      <div>
        <LgtmUploader onUploaded={onUploaded} />
      </div>
      <div>
        <ul>
          {lgtms.map((lgtm) => (
            <li key={lgtm.id}>
              <img
                src={lgtmUrl(lgtm.id)}
                style={{
                  width: 200,
                  maxHeight: 200,
                  objectFit: "contain",
                }}
                alt="LGTM"
              />
            </li>
          ))}
        </ul>
        {hasNextPage && <button onClick={onLoadMore}>Load more</button>}
      </div>
    </div>
  );
}
