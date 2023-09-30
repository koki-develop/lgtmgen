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
      <div className="fixed h-full w-full top-0 left-0 flex justify-center items-center bg-black bg-opacity-25">
        <Dialog.Panel className="bg-white rounded flex flex-col gap-4 items-center px-8 py-4">
          <Dialog.Description>{t.confirmGeneration}</Dialog.Description>

          <img className="max-w-full max-h-72 border" src={dataUrl} alt="" />

          <div className="flex gap-2 w-full">
            <button
              className="flex-grow shadow-md w-64 rounded bg-gray-200 hover:bg-gray-400 transition py-2"
              onClick={onCancel}
            >
              {t.cancel}
            </button>
            <button
              className="flex-grow shadow-md w-64 rounded text-white bg-primary-main hover:bg-primary-dark transition py-2"
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
