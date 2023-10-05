import ja from "@/locale/ja";
import en from "@/locale/en";
import React from "react";

export type I18n = {
  app: string;
  lgtm: string;
  searchImage: string;
  favorite: string;

  loadMore: string;
  upload: string;
  confirmGeneration: string;
  generate: string;
  send: string;
  supplement: string;
  illegal: string;
  inappropriate: string;
  other: string;
  keyword: string;
  noFavorites: string;
  cancel: string;

  privacyPolicy: string;
  usagePrecautions: string;

  successToGenerate: string;
  failedToGenerate: string;
  successToSend: string;
  failedToSend: string;

  fileTooLarge: string;
  unsupportedImageFormat: string;
  copiedToClipboard: string;
  rateLimitReached: string;

  aboutTrafficAnalysisTool: string;
  aboutTrafficAnalysisToolContent: React.ReactNode;
  aboutPrivacyPolicyChange: string;
  aboutPrivacyPolicyChangeContent: string;

  usagePrecautionsItems: readonly string[];
};

export const i18n = (locale: string): I18n => {
  const t = { ja, en }[locale];
  if (!t) throw new Error(`Unsupported locale: ${locale}`);

  return t;
};
