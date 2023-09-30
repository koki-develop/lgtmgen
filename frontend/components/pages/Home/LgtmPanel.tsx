import React from "react";
import { ModelsLGTM } from "@/lib/generated/api";
import { lgtmUrl } from "@/lib/image";
import { LgtmUploader } from "./LgtmUploader";
import {
  DocumentDuplicateIcon,
  HeartIcon,
  FlagIcon,
} from "@heroicons/react/24/outline";

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
        <ul className="grid grid-cols-4 gap-4">
          {lgtms.map((lgtm) => (
            <li
              key={lgtm.id}
              className="bg-white shadow-md rounded flex flex-col gap-2"
            >
              <div className="flex items-center justify-center flex-grow p-2">
                <img
                  className="max-w-100 max-h-36 border"
                  src={lgtmUrl(lgtm.id)}
                  alt="LGTM"
                />
              </div>
              <div className="flex text-white rounded-b overflow-hidden">
                <button className="bg-primary-main border-t border-t-primary-main hover:border-t-primary-dark transition hover:bg-primary-dark flex-grow flex justify-center py-2">
                  <DocumentDuplicateIcon className="w-6 h-6 " />
                </button>
                <button className="flex-grow transition bg-white text-favorite-dark hover:bg-favorite-light border-t hover:border-t-favorite-light flex justify-center py-2">
                  <HeartIcon className="w-6 h-6" />
                </button>
                <button className="flex-grow transition bg-report-main hover:bg-report-dark flex justify-center border-t border-t-report-main hover:border-t-report-dark py-2">
                  <FlagIcon className="w-6 h-6" />
                </button>
              </div>
            </li>
          ))}
        </ul>
        {hasNextPage && <button onClick={onLoadMore}>Load more</button>}
      </div>
    </div>
  );
}
