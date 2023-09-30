import { I18n } from "@/lib/i18n";

const en: I18n = {
  app: "LGTM Generator",

  lgtm: "LGTM",
  searchImage: "Search Image",
  favorite: "Favorite",

  loadMore: "More",
  upload: "Upload",
  confirmGeneration: "Would you like to generate an LGTM with this image?",
  generate: "Generate",
  cancel: "Cancel",

  successToGenerate: "Successfully generated LGTM",
  failedToGenerate: "Failed to generate LGTM",
  unsupportedImageFormat: "Unsupported image format",
  copiedToClipboard: "Copied to clipboard",
} as const;

export default en;
