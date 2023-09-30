import { Dialog } from "@headlessui/react";
import { useI18n } from "@/providers/I18nProvider";
import React, { useCallback } from "react";
import clsx from "clsx";

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

  if (dataUrl == null) {
    return null;
  }

  const handleClose = useCallback(() => {
    if (generating) return;
    onCancel();
  }, [generating, onCancel]);

  return (
    <Dialog open onClose={handleClose}>
      <div className="fixed left-0 top-0 flex h-full w-full items-center justify-center bg-black bg-opacity-25">
        <Dialog.Panel className="flex flex-col items-center gap-4 rounded bg-white px-8 py-4">
          <Dialog.Description>{t.confirmGeneration}</Dialog.Description>

          <img className="max-h-72 max-w-full border" src={dataUrl} alt="" />

          <div className="flex w-full gap-2">
            <button
              className={clsx(
                "button-secondary",
                "w-64 flex-grow rounded py-2 shadow-md",
              )}
              onClick={onCancel}
              disabled={generating}
            >
              {t.cancel}
            </button>
            <button
              className={clsx(
                "button-primary",
                "w-64 flex-grow rounded py-2 shadow-md transition",
              )}
              onClick={onGenerate}
              disabled={generating}
            >
              {t.generate}
            </button>
          </div>
        </Dialog.Panel>
      </div>
    </Dialog>
  );
}
