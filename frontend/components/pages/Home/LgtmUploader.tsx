import { api } from "@/lib/api";
import { ModelsLGTM, ServiceErrCode } from "@/lib/generated/api";
import { dataUrlToBase64, fileToDataUrl } from "@/lib/image";
import React, { useCallback } from "react";

export type LgtmUploaderProps = {
  onUploaded: (lgtm: ModelsLGTM) => void;
};

export const LgtmUploader = ({ onUploaded }: LgtmUploaderProps) => {
  const [file, setFile] = React.useState<File | null>(null);
  const [imageDataUrl, setImageDataUrl] = React.useState<string | null>(null);

  const handleChangeFile = useCallback(
    async (e: React.ChangeEvent<HTMLInputElement>) => {
      if (!e.target.files) return;

      const file = e.target.files[0];
      setFile(file);
      const dataUrl = await fileToDataUrl(file);
      setImageDataUrl(dataUrl);
    },
    [],
  );

  const handleClickUpload = useCallback(async () => {
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
    <div>
      <div>
        <button onClick={handleClickUpload}>upload</button>
      </div>
      <div>
        <input type="file" accept="image/*" onChange={handleChangeFile} />
        {file && imageDataUrl && (
          <img
            src={imageDataUrl}
            alt={file.name}
            style={{ width: 400, maxHeight: 400, objectFit: "contain" }}
          />
        )}
      </div>
    </div>
  );
};
