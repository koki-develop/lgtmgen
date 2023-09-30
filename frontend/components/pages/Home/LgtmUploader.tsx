import { api } from "@/lib/api";
import { ModelsLGTM, ServiceErrCode } from "@/lib/generated/api";
import { dataUrlToBase64, fileToDataUrl } from "@/lib/image";
import { useI18n } from "@/providers/I18nProvider";
import { PlusCircleIcon } from "@heroicons/react/24/solid";
import React, { useCallback, useRef } from "react";
import LgtmPreview from "./LgtmPreview";

export type LgtmUploaderProps = {
  onUploaded: (lgtm: ModelsLGTM) => void;
};

export const LgtmUploader = ({ onUploaded }: LgtmUploaderProps) => {
  const { t } = useI18n();

  const [file, setFile] = React.useState<File | null>(null);
  const [imageDataUrl, setImageDataUrl] = React.useState<string | null>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  const [inputKey, setInputKey] = React.useState(0);

  const handleClickUpload = useCallback(() => {
    inputRef.current?.click();
  }, []);

  const handleChangeFile = useCallback(
    async (e: React.ChangeEvent<HTMLInputElement>) => {
      setInputKey((key) => key + 1); // reset input
      if (!e.target.files) return;

      const file = e.target.files[0];
      setFile(file);
      const dataUrl = await fileToDataUrl(file);
      setImageDataUrl(dataUrl);
    },
    [],
  );

  const handleClosePreview = useCallback(() => {
    setFile(null);
    setImageDataUrl(null);
  }, []);

  const handleGenerate = useCallback(async () => {
    if (!file || !imageDataUrl) return;

    const base64 = dataUrlToBase64(imageDataUrl);
    const response = await api.v1.lgtmsCreate({ base64 });
    if (response.ok) {
      onUploaded(response.data);
      setFile(null);
      setImageDataUrl(null);
      return;
    }

    switch (response.error.code) {
      case ServiceErrCode.ErrCodeUnsupportedImageFormat:
        alert("Unsupported image format");
        break;
      case ServiceErrCode.ErrCodeInternalServerError:
        alert("Internal server error");
        break;
      default:
        throw response.error;
    }
  }, [file, imageDataUrl, onUploaded]);

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
        className="flex gap-2 fixed bottom-4 right-4 bg-primary-main hover:bg-primary-dark transition text-white py-4 px-4 rounded-full shadow-md"
        onClick={handleClickUpload}
      >
        <PlusCircleIcon className="w-6 h-6" />
        {t.upload}
      </button>

      <LgtmPreview
        dataUrl={imageDataUrl}
        onCancel={handleClosePreview}
        onGenerate={handleGenerate}
      />
    </>
  );
};
