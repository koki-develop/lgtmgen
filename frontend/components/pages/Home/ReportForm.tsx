import { lgtmUrl } from "@/lib/image";
import { Dialog, RadioGroup } from "@headlessui/react";
import React, { useCallback, useMemo, useState } from "react";
import clsx from "clsx";
import { useI18n } from "@/providers/I18nProvider";
import { useSendReport } from "@/lib/models/report/reportHooks";
import { ModelsReportType } from "@/lib/generated/api";
import { CheckCircleIcon as CheckCircleIconOutline } from "@heroicons/react/24/outline";
import { CheckCircleIcon } from "@heroicons/react/24/solid";

export type ReportFormProps = {
  lgtmId: string | null;

  onClose: () => void;
};

export default function ReportForm({ lgtmId, onClose }: ReportFormProps) {
  const { t } = useI18n();
  const { sendReport, sending } = useSendReport();

  const [type, setType] = useState<ModelsReportType | null>(null);
  const [text, setText] = useState<string>("");

  const isValid = useMemo(() => {
    if (type == null) return false;
    if (text.trim() === "") return false;
    return true;
  }, [type, text]);

  const handleClose = useCallback(() => {
    if (sending) return;
    onClose();
    setType(null);
    setText("");
  }, [onClose, sending]);

  const handleClickSend = useCallback(async () => {
    if (lgtmId == null) return;
    if (type == null) return;
    if (text.trim() === "") return;

    const report = await sendReport({
      lgtm_id: lgtmId,
      type: type,
      text: text,
    });
    if (report) {
      handleClose();
    }
  }, [sendReport, handleClose, lgtmId, text, type]);

  const handleChangeType = useCallback((reportType: ModelsReportType) => {
    setType(reportType);
  }, []);

  const handleChangeText = useCallback(
    (e: React.ChangeEvent<HTMLTextAreaElement>) => {
      setText(e.target.value);
    },
    [],
  );

  if (lgtmId === null) {
    return null;
  }

  return (
    <Dialog open onClose={handleClose}>
      <div className="fixed left-0 top-0 flex h-full w-full items-center justify-center bg-black bg-opacity-25">
        <Dialog.Panel className="flex flex-col items-center gap-4 rounded bg-white px-8 py-4">
          <img
            className="max-h-72 max-w-full border"
            src={lgtmUrl(lgtmId)}
            alt=""
          />

          <RadioGroup
            className="flex w-full flex-col gap-2"
            value={type}
            onChange={handleChangeType}
          >
            {Object.values(ModelsReportType).map((type) => (
              <RadioGroup.Option key={type} value={type}>
                {({ checked }) => (
                  <div
                    className={clsx(
                      "flex cursor-pointer items-center gap-2 rounded border p-2 transition",
                      {
                        "hover:bg-gray-100": !checked,
                        "bg-primary-light": checked,
                      },
                    )}
                  >
                    {checked ? (
                      <CheckCircleIcon className="h-6 w-6 text-primary-main" />
                    ) : (
                      <CheckCircleIconOutline className="h-6 w-6 text-gray-400" />
                    )}
                    {t[type]}
                  </div>
                )}
              </RadioGroup.Option>
            ))}
          </RadioGroup>

          <textarea
            className="w-full rounded border p-2 outline-none"
            rows={4}
            placeholder={t.supplement}
            value={text}
            onChange={handleChangeText}
          />

          <div className="flex w-full gap-2">
            <button
              className={clsx(
                "button-secondary",
                "w-64 flex-grow rounded py-2 shadow-md",
              )}
              onClick={handleClose}
              disabled={sending}
            >
              {t.cancel}
            </button>
            <button
              className={clsx(
                "button-primary",
                "w-64 flex-grow rounded py-2 shadow-md transition",
              )}
              onClick={handleClickSend}
              disabled={!isValid || sending}
            >
              {t.send}
            </button>
          </div>
        </Dialog.Panel>
      </div>
    </Dialog>
  );
}
