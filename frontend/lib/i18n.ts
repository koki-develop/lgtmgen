import ja from "@/locale/ja";
import en from "@/locale/en";

export const i18n = (locale: string) => {
  const t = { ja, en }[locale];
  if (!t) throw new Error(`Unsupported locale: ${locale}`);

  return t;
};
