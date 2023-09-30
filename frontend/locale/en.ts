import { I18n } from "@/lib/i18n";

const en: I18n = {
  app: "LGTM Generator",

  lgtm: "LGTM",
  searchImage: "Search Image",
  favorite: "Favorite",

  upload: "Upload",
  confirmGeneration: "Would you like to generate an LGTM with this image?",
  generate: "Generate",
  cancel: "Cancel",

  failedToGenerate: "Failed to generate LGTM",
  unsupportedImageFormat: "Unsupported image format",
} as const;

export default en;
