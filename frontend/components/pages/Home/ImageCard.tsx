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
        "flex flex-col gap-2 rounded bg-white shadow-md",
      )}
    >
      <div className="flex flex-grow items-center justify-center p-2">
        <img className="max-w-100 max-h-36 border" src={src} alt={alt} />
      </div>
      <div className="flex overflow-hidden rounded-b text-white">
        <button
          className={clsx(
            "button-primary",
            "border-t border-t-primary-main hover:border-t-primary-dark",
            "flex flex-grow justify-center py-2",
          )}
        >
          <DocumentDuplicateIcon className="h-6 w-6" />
        </button>
        <button
          className={clsx(
            "flex flex-grow justify-center border-t bg-white py-2 text-favorite-dark transition hover:border-t-favorite-light hover:bg-favorite-light",
          )}
        >
          <HeartIcon className="h-6 w-6" />
        </button>
        <button className="flex flex-grow justify-center border-t border-t-report-main bg-report-main py-2 transition hover:border-t-report-dark hover:bg-report-dark">
          <FlagIcon className="h-6 w-6" />
        </button>
      </div>
    </div>
  );
}
