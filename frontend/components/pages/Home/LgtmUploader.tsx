import { ModelsLGTM } from "@/lib/generated/api";
import { dataUrlToBase64, fileToDataUrl, resizeDataUrl } from "@/lib/image";
import { useI18n } from "@/providers/I18nProvider";
import { PlusCircleIcon } from "@heroicons/react/24/solid";
import React, { useCallback, useRef } from "react";
import LgtmPreview from "./LgtmPreview";
import clsx from "clsx";
import { useGenerateLgtm } from "@/lib/models/lgtm/lgtmHooks";
import { useToast } from "@/lib/toast";

export type LgtmUploaderProps = {
  onUploaded: (lgtm: ModelsLGTM) => void;
};

const maxSideLength = 425;

export const LgtmUploader = ({ onUploaded }: LgtmUploaderProps) => {
  const { t } = useI18n();
  const { enqueueToast } = useToast();

  const [file, setFile] = React.useState<File | null>(null);
  const [imageDataUrl, setImageDataUrl] = React.useState<string | null>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  const [inputKey, setInputKey] = React.useState(0);

  const { generateLgtm, generating } = useGenerateLgtm();

  const handleClickUpload = useCallback(() => {
    inputRef.current?.click();
  }, []);

  const handleChangeFile = useCallback(
    async (e: React.ChangeEvent<HTMLInputElement>) => {
      setInputKey((key) => key + 1); // reset input
      if (!e.target.files) return;

      const file = e.target.files[0];
      let dataUrl = await fileToDataUrl(file);

      switch (file.type) {
        case "image/jpeg":
        case "image/png":
          dataUrl = await resizeDataUrl(dataUrl, maxSideLength, file.type);
          break;
        default:
          const sizeLimit = 1024 * 1024 * 4; // 4MB
          if (file.size > sizeLimit) {
            enqueueToast(`${t.fileTooLarge}: ${file.name}`, "error");
            return;
          }
          break;
      }

      setFile(file);
      setImageDataUrl(dataUrl);
    },
    [enqueueToast, t],
  );

  const handleClosePreview = useCallback(() => {
    setFile(null);
    setImageDataUrl(null);
  }, []);

  const handleGenerate = useCallback(async () => {
    if (!file || !imageDataUrl) return;

    const base64 = dataUrlToBase64(imageDataUrl);
    const lgtm = await generateLgtm({ base64 });
    if (lgtm) {
      onUploaded(lgtm);
      setFile(null);
      setImageDataUrl(null);
    }
  }, [file, imageDataUrl, generateLgtm, onUploaded]);

  return (
    <>
      <input
        ref={inputRef}
        key={inputKey}
        className="hidden"
        type="file"
        accept="image/*"
        onChange={handleChangeFile}
      />
      <button
        className={clsx(
          "button-primary rounded-full text-white shadow-md",
          "fixed bottom-4 right-4 z-10",
          "flex gap-2 px-4 py-4",
        )}
        onClick={handleClickUpload}
      >
        <PlusCircleIcon className="h-6 w-6" />
        {t.upload}
      </button>

      <LgtmPreview
        generating={generating}
        src={imageDataUrl}
        onCancel={handleClosePreview}
        onGenerate={handleGenerate}
      />
    </>
  );
};
