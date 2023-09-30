import { Dialog } from "@headlessui/react";
import { useI18n } from "@/providers/I18nProvider";
import React from "react";

export type LgtmPreviewProps = {
  dataUrl: string | null;

  onGenerate: () => void;
  onCancel: () => void;
};

export default function LgtmPreview({
  dataUrl,
  onCancel,
  onGenerate,
}: LgtmPreviewProps) {
  const { t } = useI18n();

  if (dataUrl == null) {
    return null;
  }

  return (
    <Dialog open onClose={onCancel}>
      <div className="fixed left-0 top-0 flex h-full w-full items-center justify-center bg-black bg-opacity-25">
        <Dialog.Panel className="flex flex-col items-center gap-4 rounded bg-white px-8 py-4">
          <Dialog.Description>{t.confirmGeneration}</Dialog.Description>

          <img className="max-h-72 max-w-full border" src={dataUrl} alt="" />

          <div className="flex w-full gap-2">
            <button
              className="w-64 flex-grow rounded bg-gray-200 py-2 shadow-md transition hover:bg-gray-400"
              onClick={onCancel}
            >
              {t.cancel}
            </button>
            <button
              className="w-64 flex-grow rounded bg-primary-main py-2 text-white shadow-md transition hover:bg-primary-dark"
              onClick={onGenerate}
            >
              {t.generate}
            </button>
          </div>
        </Dialog.Panel>
      </div>
    </Dialog>
  );
}
