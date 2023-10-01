import { api } from "@/lib/api";
import { ServiceCreateReportInput, ServiceErrCode } from "@/lib/generated/api";
import { useToast } from "@/lib/toast";
import { useI18n } from "@/providers/I18nProvider";
import { useCallback, useState } from "react";

export const useSendReport = () => {
  const { t } = useI18n();
  const { enqueueToast } = useToast();

  const [sending, setSending] = useState<boolean>(false);

  const sendReport = useCallback(
    async (input: ServiceCreateReportInput) => {
      setSending(true);
      return await api.v1
        .reportsCreate(input)
        .then((response) => {
          if (response.ok) {
            enqueueToast(t.successToSend);
            return response.data;
          }

          switch (response.error.code) {
            case ServiceErrCode.ErrCodeInternalServerError:
              enqueueToast(t.failedToSend);
              break;
          }

          return null;
        })
        .finally(() => {
          setSending(false);
        });
    },
    [enqueueToast, t],
  );

  return { sendReport, sending };
};
