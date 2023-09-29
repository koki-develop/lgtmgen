import ja from "@/app/locale/ja";
import en from "@/app/locale/en";

export const i18n = (locale: string) => {
  const t = { ja, en }[locale];
  if (!t) throw new Error(`Unsupported locale: ${locale}`);

  return t;
};
