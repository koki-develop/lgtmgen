import React, { useMemo } from "react";
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
  const buttons = useMemo(
    () => [
      {
        icon: <DocumentDuplicateIcon className="h-6 w-6" />,
        additionalClass: clsx(
          "button-primary",
          "border-t-primary-main hover:border-t-primary-dark",
        ),
      },
      {
        icon: <HeartIcon className="h-6 w-6" />,
        additionalClass: clsx(
          "bg-white hover:bg-favorite-light",
          "hover:border-t-favorite-light",
          "text-favorite-dark",
        ),
      },
      {
        icon: <FlagIcon className="h-6 w-6" />,
        additionalClass: clsx(
          "border-t-report-main hover:border-t-report-dark",
          "bg-report-main hover:bg-report-dark",
        ),
      },
    ],
    [],
  );

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
        {buttons.map(({ icon, additionalClass }, index) => (
          <button
            key={index}
            className={clsx(
              additionalClass,
              "flex flex-grow justify-center",
              "border-t py-2 transition",
            )}
          >
            {icon}
          </button>
        ))}
      </div>
    </div>
  );
}
