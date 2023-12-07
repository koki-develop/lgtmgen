import clsx from "clsx";
import React from "react";

export type ImageCardProps = {
  className?: string;
  children?: React.ReactNode;

  src: string;
  alt: string;
  icon?: React.ReactNode;

  onClick: () => void;
};

export default function ImageCard({
  className,
  children,

  src,
  alt,
  icon,

  onClick,
}: ImageCardProps) {
  return (
    <div
      className={clsx(className, "flex flex-col rounded bg-white shadow-md")}
    >
      <button
        className="group relative flex h-40 flex-grow items-center justify-center p-2 transition hover:bg-gray-100"
        onClick={onClick}
      >
        {icon && (
          <span className="absolute opacity-0 transition group-hover:opacity-80">
            {icon}
          </span>
        )}
        <img className="max-w-100 max-h-36 border" src={src} alt={alt} />
      </button>

      {children}
    </div>
  );
}
