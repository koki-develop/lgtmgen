import { ServiceCreateLGTMInput, ServiceErrCode } from "@/lib/generated/api";
import { useCallback, useState } from "react";
import { api } from "@/lib/api";
import { useI18n } from "@/providers/I18nProvider";

export const useGenerateLgtm = () => {
  const { t } = useI18n();

  const [generating, setGenerating] = useState<boolean>(false);

  const generateLgtm = useCallback(async (input: ServiceCreateLGTMInput) => {
    setGenerating(true);
    return await api.v1
      .lgtmsCreate(input)
      .then((response) => {
        if (response.ok) {
          return response.data;
        }

        switch (response.error.code) {
          case ServiceErrCode.ErrCodeUnsupportedImageFormat:
            alert(t.unsupportedImageFormat);
            break;
          case ServiceErrCode.ErrCodeInternalServerError:
            alert(t.failedToGenerate);
            break;
          default:
            throw response.error;
        }

        return null;
      })
      .finally(() => {
        setGenerating(false);
      });
  }, []);

  return { generateLgtm, generating };
};
