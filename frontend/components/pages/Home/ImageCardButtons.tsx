import React, { useCallback, useEffect, useState } from "react";
import clsx from "clsx";
import {
  CheckIcon,
  DocumentDuplicateIcon,
  HeartIcon as HeartIconOutline,
  FlagIcon,
} from "@heroicons/react/24/outline";
import { HeartIcon as HeartIconSolid } from "@heroicons/react/24/solid";
import { Menu } from "@headlessui/react";
import copy from "copy-to-clipboard";
import { useToast } from "@/lib/toast";
import { useI18n } from "@/providers/I18nProvider";
import { lgtmUrl } from "@/lib/image";

export type ImageCardButtonsProps = {
  lgtmId: string;
  favorited: boolean;

  onFavorite: (id: string) => void;
  onUnfavorite: (id: string) => void;
  onStartReport: (id: string) => void;
};

// TODO: Refactor
export default function ImageCardButtons({
  lgtmId,
  favorited,
  onFavorite,
  onUnfavorite,
  onStartReport,
}: ImageCardButtonsProps) {
  const { t } = useI18n();
  const { enqueueToast } = useToast();

  const [copied, setCopied] = useState<boolean>(false);

  const handleClickMarkdown = useCallback(() => {
    copy(`![LGTM](${lgtmUrl(lgtmId)})`);
    setCopied(true);
    enqueueToast(t.copiedToClipboard);
  }, [enqueueToast, lgtmId, t]);

  const handleClickHTML = useCallback(() => {
    copy(`<img src="${lgtmUrl(lgtmId)}" alt="LGTM" />`);
    setCopied(true);
    enqueueToast(t.copiedToClipboard);
  }, [enqueueToast, lgtmId, t]);

  const handleClickFavorite = useCallback(() => {
    if (favorited) {
      onUnfavorite(lgtmId);
    } else {
      onFavorite(lgtmId);
    }
  }, [lgtmId, favorited, onFavorite, onUnfavorite]);

  const handleClickReport = useCallback(() => {
    onStartReport(lgtmId);
  }, [lgtmId, onStartReport]);

  const baseClass = clsx(
    "flex flex-grow justify-center",
    "border-t py-2 transition",
  );

  useEffect(() => {
    if (copied) {
      const timer = setTimeout(() => {
        setCopied(false);
      }, 2000);
      return () => clearTimeout(timer);
    }
  }, [copied]);

  return (
    <div className="relative flex rounded-b text-white">
      {/* Copy */}
      <Menu>
        <Menu.Button
          className={clsx(
            baseClass,
            "rounded-bl",
            "button-primary",
            "border-t-primary-main hover:border-t-primary-dark",
          )}
        >
          {copied ? (
            <CheckIcon className="h-6 w-6 text-green-500" />
          ) : (
            <DocumentDuplicateIcon className="h-6 w-6" />
          )}
        </Menu.Button>
        <Menu.Items className="absolute -top-16 left-6 flex flex-col divide-y rounded bg-white text-gray-600 shadow-md">
          <Menu.Item>
            {({ active }) => (
              <button
                className={clsx("px-4 py-2 transition", {
                  "bg-gray-200": active,
                })}
                onClick={handleClickMarkdown}
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
                onClick={handleClickHTML}
              >
                HTML
              </button>
            )}
          </Menu.Item>
        </Menu.Items>
      </Menu>

      {/* Favorite */}
      <button
        className={clsx(baseClass, "text-favorite-dark", {
          "bg-white hover:border-t-favorite-light hover:bg-favorite-light":
            !favorited,
          "bg-favorite-light hover:border-t-favorite-main hover:bg-favorite-main":
            favorited,
        })}
        onClick={handleClickFavorite}
      >
        {favorited ? (
          <HeartIconSolid className="h-6 w-6" />
        ) : (
          <HeartIconOutline className="h-6 w-6" />
        )}
      </button>

      {/* Report */}
      <button
        className={clsx(
          baseClass,
          "rounded-br",
          "border-t-report-main hover:border-t-report-dark",
          "bg-report-main hover:bg-report-dark",
        )}
        onClick={handleClickReport}
      >
        <FlagIcon className="h-6 w-6" />
      </button>
    </div>
  );
}
