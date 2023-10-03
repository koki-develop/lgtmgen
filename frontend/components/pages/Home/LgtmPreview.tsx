import Dialog from "@/components/util/Dialog";
import { useI18n } from "@/providers/I18nProvider";
import React from "react";

export type LgtmPreviewProps = {
  dataUrl: string | null;
  generating: boolean;

  onGenerate: () => void;
  onCancel: () => void;
};

export default function LgtmPreview({
  dataUrl,
  generating,
  onCancel,
  onGenerate,
}: LgtmPreviewProps) {
  const { t } = useI18n();

  return (
    <Dialog
      title={t.confirmGeneration}
      submitText={t.generate}
      open={Boolean(dataUrl)}
      disabled={generating}
      onSubmit={onGenerate}
      onClose={onCancel}
    >
      {dataUrl && (
        <img
          className="max-h-72
          max-w-full border"
          src={dataUrl}
          alt=""
        />
      )}
    </Dialog>
  );
}
