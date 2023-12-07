"use client";

import React from "react";
import { Toaster } from "react-hot-toast";
import I18nProvider from "./I18nProvider";

export type ProvidersProps = {
  locale: string;
  children: React.ReactNode;
};

export default function Providers({ children, locale }: ProvidersProps) {
  return (
    <>
      <Toaster />
      <I18nProvider locale={locale}>{children}</I18nProvider>
    </>
  );
}
