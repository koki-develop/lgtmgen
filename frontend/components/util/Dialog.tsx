import { useI18n } from "@/providers/I18nProvider";
import { Dialog as HeadlessDialog } from "@headlessui/react";
import clsx from "clsx";
import React, { useCallback } from "react";

export type DialogProps = {
  children: React.ReactNode;

  title?: string;
  submitText: string;
  open: boolean;
  loading: boolean;
  disabled: boolean;

  onSubmit: () => void;
  onClose: () => void;
};

export default function Dialog({
  children,

  title,
  submitText,
  open,
  loading,
  disabled,

  onSubmit,
  onClose,
}: DialogProps) {
  const { t } = useI18n();

  const handleClose = useCallback(() => {
    if (loading) return;
    onClose();
  }, [loading, onClose]);

  const handleSubmit = useCallback(() => {
    if (loading) return;
    onSubmit();
  }, [loading, onSubmit]);

  return (
    <HeadlessDialog open={open} onClose={handleClose}>
      <div className="fixed inset-0 bg-black/30" />

      <div className="fixed inset-0 flex items-center justify-center">
        <HeadlessDialog.Panel className="flex flex-col items-center gap-4 rounded bg-white p-4">
          {title && <HeadlessDialog.Title>{title}</HeadlessDialog.Title>}

          {children}

          <div className="flex gap-2">
            <button
              className={clsx(
                "button-secondary",
                "w-64 flex-grow rounded py-2 shadow-md",
              )}
              onClick={handleClose}
              disabled={loading}
            >
              {t.cancel}
            </button>
            <button
              className={clsx(
                "button-primary",
                "w-64 flex-grow rounded py-2 shadow-md transition",
              )}
              onClick={handleSubmit}
              disabled={loading || disabled}
            >
              {submitText}
            </button>
          </div>
        </HeadlessDialog.Panel>
      </div>
    </HeadlessDialog>
  );
}
