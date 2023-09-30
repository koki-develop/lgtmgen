"use client";

import React from "react";
import I18nProvider from "./I18nProvider";

export type ProvidersProps = {
  locale: string;
  children: React.ReactNode;
};

export default function Providers({ children, locale }: ProvidersProps) {
  return <I18nProvider locale={locale}>{children}</I18nProvider>;
}
