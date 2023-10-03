import { useI18n } from "@/providers/I18nProvider";
import { Dialog as HeadlessDialog } from "@headlessui/react";
import clsx from "clsx";
import React, { useCallback } from "react";

export type DialogProps = {
  children: React.ReactNode;

  title: string;
  submitText: string;
  open: boolean;
  disabled: boolean;

  onSubmit: () => void;
  onClose: () => void;
};

export default function Dialog({
  children,

  title,
  submitText,
  open,
  disabled,

  onSubmit,
  onClose,
}: DialogProps) {
  const { t } = useI18n();

  const handleClose = useCallback(() => {
    if (disabled) return;
    onClose();
  }, [disabled, onClose]);

  const handleSubmit = useCallback(() => {
    if (disabled) return;
    onSubmit();
  }, [disabled, onSubmit]);

  return (
    <HeadlessDialog open={open} onClose={handleClose}>
      <div className="fixed inset-0 bg-black/30" />

      <div className="fixed inset-0 flex items-center justify-center">
        <HeadlessDialog.Panel className="flex flex-col items-center gap-4 rounded bg-white p-4">
          <HeadlessDialog.Title>{title}</HeadlessDialog.Title>

          {children}

          <div className="flex gap-2">
            <button
              className={clsx(
                "button-secondary",
                "w-64 flex-grow rounded py-2 shadow-md",
              )}
              onClick={handleClose}
              disabled={disabled}
            >
              {t.cancel}
            </button>
            <button
              className={clsx(
                "button-primary",
                "w-64 flex-grow rounded py-2 shadow-md transition",
              )}
              onClick={handleSubmit}
              disabled={disabled}
            >
              {submitText}
            </button>
          </div>
        </HeadlessDialog.Panel>
      </div>
    </HeadlessDialog>
  );
}
