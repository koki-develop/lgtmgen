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
      <div className="fixed inset-0 z-30 bg-black/30" />

      <div className="container fixed inset-0 z-40 mx-auto flex w-full items-center justify-center px-4 py-2">
        <HeadlessDialog.Panel className="flex max-h-full max-w-full flex-col items-center gap-4 overflow-y-auto rounded bg-white p-4">
          {title && <HeadlessDialog.Title>{title}</HeadlessDialog.Title>}

          {children}

          <div className="flex w-full gap-2">
            <button
              className={clsx(
                "button-secondary text-sm sm:text-base",
                "w-full flex-grow rounded py-2 shadow-md",
              )}
              onClick={handleClose}
              disabled={loading}
            >
              {t.cancel}
            </button>
            <button
              className={clsx(
                "button-primary text-sm sm:text-base",
                "w-full flex-grow rounded py-2 shadow-md transition",
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
