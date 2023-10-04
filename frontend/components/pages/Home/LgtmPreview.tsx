import Dialog from "@/components/util/Dialog";
import { useI18n } from "@/providers/I18nProvider";
import React from "react";

export type LgtmPreviewProps = {
  src: string | null;
  generating: boolean;
  open: boolean;

  onGenerate: () => void;
  onCancel: () => void;
};

export default function LgtmPreview({
  src,
  generating,
  open,
  onCancel,
  onGenerate,
}: LgtmPreviewProps) {
  const { t } = useI18n();

  return (
    <Dialog
      title={t.confirmGeneration}
      submitText={t.generate}
      open={open}
      loading={generating}
      disabled={generating}
      onSubmit={onGenerate}
      onClose={onCancel}
    >
      <div className="h-72">
        {src ? (
          <img className="h-full max-w-full border" src={src} alt="" />
        ) : (
          <div className="flex h-full items-center">
            <div className="loader" />
          </div>
        )}
      </div>
    </Dialog>
  );
}
