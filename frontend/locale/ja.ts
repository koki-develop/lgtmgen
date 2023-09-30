import { I18n } from "@/lib/i18n";

const ja: I18n = {
  app: "LGTM Generator",
  lgtm: "LGTM",
  searchImage: "画像検索",
  favorite: "お気に入り",

  loadMore: "もっと見る",
  upload: "アップロード",
  confirmGeneration: "この画像で LGTM を生成しますか？",
  generate: "生成",
  cancel: "キャンセル",

  successToGenerate: "LGTM を生成しました",
  failedToGenerate: "LGTM の生成に失敗しました",
  unsupportedImageFormat: "サポートされていない画像形式です",
} as const;

export default ja;
