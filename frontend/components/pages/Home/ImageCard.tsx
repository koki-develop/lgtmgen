import React from "react";
import clsx from "clsx";
import {
  DocumentDuplicateIcon,
  HeartIcon,
  FlagIcon,
} from "@heroicons/react/24/outline";

export type ImageCardProps = {
  className?: string;

  src: string;
  alt: string;
};

export default function ImageCard({ className, src, alt }: ImageCardProps) {
  return (
    <div
      className={clsx(
        className,
        "bg-white shadow-md rounded flex flex-col gap-2",
      )}
    >
      <div className="flex items-center justify-center flex-grow p-2">
        <img className="max-w-100 max-h-36 border" src={src} alt={alt} />
      </div>
      <div className="flex text-white rounded-b overflow-hidden">
        <button
          className={clsx(
            "button-primary",
            "border-t border-t-primary-main hover:border-t-primary-dark",
            "flex-grow flex justify-center py-2",
          )}
        >
          <DocumentDuplicateIcon className="w-6 h-6" />
        </button>
        <button
          className={clsx(
            "flex-grow transition bg-white text-favorite-dark hover:bg-favorite-light border-t hover:border-t-favorite-light flex justify-center py-2",
          )}
        >
          <HeartIcon className="w-6 h-6" />
        </button>
        <button className="flex-grow transition bg-report-main hover:bg-report-dark flex justify-center border-t border-t-report-main hover:border-t-report-dark py-2">
          <FlagIcon className="w-6 h-6" />
        </button>
      </div>
    </div>
  );
}
