import { ServiceCreateLGTMInput, ServiceErrCode } from "@/lib/generated/api";
import { useCallback, useState } from "react";
import { api } from "@/lib/api";
import { useI18n } from "@/providers/I18nProvider";
import { useToast } from "@/lib/toast";

export const useFetchLgtms = (perPage: number) => {
  const [fetching, setFetching] = useState<boolean>(false);

  const fetchLgtms = useCallback(
    async (after: string) => {
      setFetching(true);

      return await api.v1
        .lgtmsList({ limit: perPage, after })
        .then((response) => {
          if (!response.ok) throw response.error;
          return response.data;
        })
        .finally(() => {
          setFetching(false);
        });
    },
    [perPage],
  );

  return { fetchLgtms, fetching };
};

export const useGenerateLgtm = () => {
  const { t } = useI18n();
  const { enqueueToast } = useToast();

  const [generating, setGenerating] = useState<boolean>(false);

  const generateLgtm = useCallback(
    async (input: ServiceCreateLGTMInput) => {
      setGenerating(true);
      return await api.v1
        .lgtmsCreate(input)
        .then((response) => {
          if (response.ok) {
            enqueueToast(t.successToGenerate);
            return response.data;
          }

          switch (response.error.code) {
            case ServiceErrCode.ErrCodeUnsupportedImageFormat:
              enqueueToast(t.unsupportedImageFormat, "error");
              break;
            case ServiceErrCode.ErrCodeRateLimitReached:
              enqueueToast(t.rateLimitReached, "error");
              break;
            case ServiceErrCode.ErrCodeInternalServerError:
              enqueueToast(t.failedToGenerate, "error");
              break;
            default:
              throw response.error;
          }

          return null;
        })
        .finally(() => {
          setGenerating(false);
        });
    },
    [enqueueToast, t],
  );

  return { generateLgtm, generating };
};
