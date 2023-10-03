import Dialog from "@/components/util/Dialog";
import { useI18n } from "@/providers/I18nProvider";
import React from "react";

export type LgtmPreviewProps = {
  src: string | null;
  generating: boolean;

  onGenerate: () => void;
  onCancel: () => void;
};

export default function LgtmPreview({
  src,
  generating,
  onCancel,
  onGenerate,
}: LgtmPreviewProps) {
  const { t } = useI18n();

  return (
    <Dialog
      title={t.confirmGeneration}
      submitText={t.generate}
      open={Boolean(src)}
      loading={generating}
      disabled={generating}
      onSubmit={onGenerate}
      onClose={onCancel}
    >
      {src && (
        <img
          className="max-h-72
          max-w-full border"
          src={src}
          alt=""
        />
      )}
    </Dialog>
  );
}
