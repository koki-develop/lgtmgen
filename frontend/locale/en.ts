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
  send: "Send",
  supplement: "( Optional ) Supplement",
  illegal:
    "Illegal ( Copyright infringement, invasion of privacy, defamation, etc. )",
  inappropriate: "Inappropriate content",
  other: "Other",
  keyword: "Keyword",
  cancel: "Cancel",

  successToGenerate: "Successfully generated LGTM.",
  failedToGenerate: "Failed to generate LGTM.",
  successToSend: "Sent.",
  failedToSend: "Failed to send.",
  unsupportedImageFormat: "Unsupported image format.",
  copiedToClipboard: "Copied to clipboard.",
  rateLimitReached:
    "The request limit has been reached.\nPlease try again later.",
} as const;

export default en;
