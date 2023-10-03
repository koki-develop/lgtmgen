import React from "react";
import clsx from "clsx";

export type ImageCardProps = {
  className?: string;
  children?: React.ReactNode;

  src: string;
  alt: string;

  onClick: () => void;
};

export default function ImageCard({
  className,
  children,
  src,
  alt,
  onClick,
}: ImageCardProps) {
  return (
    <div
      className={clsx(className, "flex flex-col rounded bg-white shadow-md")}
    >
      <button
        className="flex h-40 flex-grow items-center justify-center p-2 transition hover:bg-gray-100"
        onClick={onClick}
      >
        <img className="max-w-100 max-h-36 border" src={src} alt={alt} />
      </button>

      {children}
    </div>
  );
}
