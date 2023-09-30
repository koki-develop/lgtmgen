import React, { useMemo } from "react";
import clsx from "clsx";
import {
  DocumentDuplicateIcon,
  HeartIcon,
  FlagIcon,
} from "@heroicons/react/24/outline";
import { Menu } from "@headlessui/react";

export type ImageCardProps = {
  className?: string;

  src: string;
  alt: string;
};

export default function ImageCard({ className, src, alt }: ImageCardProps) {
  const baseClass = clsx(
    "flex flex-grow justify-center",
    "border-t py-2 transition",
  );

  const buttons = useMemo(
    () => [
      {
        button: (
          <Menu>
            <Menu.Button
              className={clsx(
                baseClass,
                "button-primary",
                "border-t-primary-main hover:border-t-primary-dark",
              )}
            >
              <DocumentDuplicateIcon className="h-6 w-6" />
            </Menu.Button>
            <Menu.Items className="absolute -top-16 left-6 flex flex-col divide-y rounded bg-white text-gray-500 shadow-md">
              <Menu.Item>
                {({ active }) => (
                  <button
                    className={clsx("px-4 py-2 transition", {
                      "bg-gray-200": active,
                    })}
                  >
                    Markdown
                  </button>
                )}
              </Menu.Item>
              <Menu.Item>
                {({ active }) => (
                  <button
                    className={clsx("px-4 py-2 transition", {
                      "bg-gray-200": active,
                    })}
                  >
                    HTML
                  </button>
                )}
              </Menu.Item>
            </Menu.Items>
          </Menu>
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
      <div className="relative flex rounded-b text-white">
        {buttons.map(
          ({ button, icon, additionalClass }, index) =>
            button || (
              <button key={index} className={clsx(baseClass, additionalClass)}>
                {icon}
              </button>
            ),
        )}
      </div>
    </div>
  );
}
