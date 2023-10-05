import { I18n, i18n } from "@/lib/i18n";
import ja from "@/locale/ja";
import React, { createContext, useContext } from "react";

export type I18nContextType = {
  t: I18n;
  locale: string;
};

export const I18nContext = createContext<I18nContextType>({
  t: ja,
  locale: "",
});

export type I18nProviderProps = {
  locale: string;
  children: React.ReactNode;
};

export default function I18nProvider({ children, locale }: I18nProviderProps) {
  return (
    <I18nContext.Provider value={{ t: i18n(locale), locale }}>
      {children}
    </I18nContext.Provider>
  );
}

export const useI18n = () => useContext(I18nContext);
