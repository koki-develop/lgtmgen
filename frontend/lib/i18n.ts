import ja from "@/locale/ja";
import en from "@/locale/en";

export type I18n = {
  app: string;
  lgtm: string;
  searchImage: string;
  favorite: string;

  loadMore: string;
  upload: string;
  confirmGeneration: string;
  generate: string;
  cancel: string;

  successToGenerate: string;
  failedToGenerate: string;
  unsupportedImageFormat: string;
  copiedToClipboard: string;
};

export const i18n = (locale: string): I18n => {
  const t = { ja, en }[locale];
  if (!t) throw new Error(`Unsupported locale: ${locale}`);

  return t;
};
