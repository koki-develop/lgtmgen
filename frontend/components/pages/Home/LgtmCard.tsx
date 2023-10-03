import { ModelsLGTM } from "@/lib/generated/api";
import React, { useCallback, useEffect, useMemo, useState } from "react";
import ImageCard from "./ImageCard";
import ImageCardButtons, { ImageCardButtonsProps } from "./ImageCardButtons";
import { useI18n } from "@/providers/I18nProvider";
import { useToast } from "@/lib/toast";
import copy from "copy-to-clipboard";
import { lgtmUrl } from "@/lib/image";
import { CheckIcon, DocumentDuplicateIcon } from "@heroicons/react/24/outline";

export type LgtmCardProps = {
  className?: string;

  lgtm: ModelsLGTM;
  favorites: string[];
} & Omit<ImageCardButtonsProps, "lgtmId" | "favorited">;

export default function LgtmCard({
  className,
  lgtm,
  favorites,
  ...buttonProps
}: LgtmCardProps) {
  const { t } = useI18n();
  const { enqueueToast } = useToast();

  const [copied, setCopied] = useState<boolean>(false);

  const favorited = useMemo(
    () => favorites.includes(lgtm.id),
    [favorites, lgtm.id],
  );

  const handleClickLgtm = useCallback(() => {
    copy(`![LGTM](${lgtmUrl(lgtm.id)})`);
    setCopied(true);
    enqueueToast(t.copiedToClipboard);
  }, [enqueueToast, lgtm.id, t]);

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
      src={lgtmUrl(lgtm.id)}
      alt="LGTM"
      icon={
        copied ? (
          <CheckIcon className="text-green-500" />
        ) : (
          <DocumentDuplicateIcon className="text-white" />
        )
      }
      onClick={handleClickLgtm}
    >
      <ImageCardButtons
        lgtmId={lgtm.id}
        favorited={favorited}
        {...buttonProps}
      />
    </ImageCard>
  );
}
