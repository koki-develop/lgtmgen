import Dialog from "@/components/util/Dialog";
import { useI18n } from "@/providers/I18nProvider";
import React from "react";

export type LgtmPreviewProps = {
  src: string | null;
  alt: string | null;
  generating: boolean;
  open: boolean;

  onGenerate: () => void;
  onCancel: () => void;
};

export default function LgtmPreview({
  src,
  alt,
  generating,
  open,
  onCancel,
  onGenerate,
}: LgtmPreviewProps) {
  const { t, locale } = useI18n();

  return (
    <Dialog
      title={t.confirmGeneration}
      submitText={t.generate}
      open={open}
      loading={generating}
      disabled={!Boolean(src) || generating}
      onSubmit={onGenerate}
      onClose={onCancel}
    >
      <div>
        {src ? (
          <img
            className="max-h-60 max-w-full border sm:max-h-72"
            src={src}
            alt={alt ?? ""}
          />
        ) : (
          <div className="flex h-full items-center">
            <div className="loader" />
          </div>
        )}
      </div>
      <div className="text-sm">{t.pleaseReadUsagePrecautions(locale)}</div>
    </Dialog>
  );
}
