import Dialog from "@/components/util/Dialog";
import { ModelsReportType } from "@/lib/generated/api";
import { lgtmUrl } from "@/lib/image";
import { useSendReport } from "@/lib/models/report/reportHooks";
import { useI18n } from "@/providers/I18nProvider";
import { RadioGroup } from "@headlessui/react";
import { CheckCircleIcon as CheckCircleIconOutline } from "@heroicons/react/24/outline";
import { CheckCircleIcon } from "@heroicons/react/24/solid";
import clsx from "clsx";
import React, { useCallback, useMemo, useState } from "react";

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
    return true;
  }, [type]);

  const handleClose = useCallback(() => {
    onClose();
    setType(null);
    setText("");
  }, [onClose]);

  const handleClickSend = useCallback(async () => {
    if (lgtmId == null) return;
    if (type == null) return;

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
      const text = e.target.value;
      if (text.length > 1000) return;
      setText(text);
    },
    [],
  );

  return (
    <Dialog
      submitText={t.send}
      open={Boolean(lgtmId)}
      loading={sending}
      disabled={!isValid}
      onSubmit={handleClickSend}
      onClose={handleClose}
    >
      {lgtmId && (
        <img
          className="max-h-60 max-w-full border sm:max-h-72"
          src={lgtmUrl(lgtmId)}
          alt=""
        />
      )}

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
                <span>
                  {checked ? (
                    <CheckCircleIcon className="h-6 w-6 text-primary-main" />
                  ) : (
                    <CheckCircleIconOutline className="h-6 w-6 text-gray-400" />
                  )}
                </span>
                <span className="text-sm sm:text-base">{t[type]}</span>
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
    </Dialog>
  );
}
