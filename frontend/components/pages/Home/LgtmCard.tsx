import React, { useCallback, useEffect, useState } from "react";
import ImageCard from "./ImageCard";
import ImageCardButtons, { ImageCardButtonsProps } from "./ImageCardButtons";
import { useI18n } from "@/providers/I18nProvider";
import { useToast } from "@/lib/toast";
import copy from "copy-to-clipboard";
import { lgtmUrl } from "@/lib/image";
import { CheckIcon, DocumentDuplicateIcon } from "@heroicons/react/24/outline";

export type LgtmCardProps = {
  className?: string;

  lgtmId: string;
  favorites: string[];
} & Omit<ImageCardButtonsProps, "lgtmId" | "favorited">;

export default function LgtmCard({
  className,
  lgtmId,
  favorites,
  ...buttonProps
}: LgtmCardProps) {
  const { t } = useI18n();
  const { enqueueToast } = useToast();

  const [copied, setCopied] = useState<boolean>(false);

  const handleClickLgtm = useCallback(() => {
    copy(`![LGTM](${lgtmUrl(lgtmId)})`);
    setCopied(true);
    enqueueToast(t.copiedToClipboard);
  }, [enqueueToast, lgtmId, t]);

  useEffect(() => {
    if (copied) {
      const timer = setTimeout(() => {
        setCopied(false);
      }, 2000);

      return () => {
        clearTimeout(timer);
      };
    }
  }, [copied]);

  return (
    <ImageCard
      className={className}
      src={lgtmUrl(lgtmId)}
      alt="LGTM"
      icon={
        copied ? (
          <CheckIcon className="h-16 w-16 text-green-500" />
        ) : (
          <DocumentDuplicateIcon className="h-16 w-16 text-white" />
        )
      }
      onClick={handleClickLgtm}
    >
      <ImageCardButtons
        lgtmId={lgtmId}
        favorites={favorites}
        {...buttonProps}
      />
    </ImageCard>
  );
}
