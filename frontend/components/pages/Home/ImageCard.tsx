import React from "react";
import clsx from "clsx";

export type ImageCardProps = {
  className?: string;
  children?: React.ReactNode;

  src: string;
  alt: string;
};

export default function ImageCard({
  className,
  children,
  src,
  alt,
}: ImageCardProps) {
  return (
    <div
      className={clsx(
        className,
        "flex flex-col gap-2 rounded bg-white shadow-md",
      )}
    >
      <div className="flex h-40 flex-grow items-center justify-center p-2">
        <img className="max-w-100 max-h-36 border" src={src} alt={alt} />
      </div>

      {children}
    </div>
  );
}
