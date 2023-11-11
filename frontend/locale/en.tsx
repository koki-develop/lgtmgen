import { I18n } from "@/lib/i18n";

const en: I18n = {
  app: "LGTM Generator",
  description:
    "This is a simple LGTM image generation service. LGTM stands for 'Looks Good To Me', a type of internet slang used during code reviews, among other situations.",

  feedback: "Feedback",
  feedbackLink: "https://forms.gle/4xcwMaZveyXK1jvu6",

  lgtm: "LGTM",
  searchImage: "Search Image",
  favorite: "Favorite",

  loadMore: "More",
  upload: "Upload Image",
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
  random: "Random",
  reload: "Reload",

  privacyPolicy: "Privacy Policy",
  usagePrecautions: "Usage Precautions",
  pleaseReadUsagePrecautions: (locale: string) => (
    <>
      Please read the{" "}
      <a
        className="text-primary-main underline"
        href={`/${locale}/usage-precautions`}
        target="_blank"
        rel="noopener noreferrer"
      >
        usage precautions
      </a>{" "}
      before generating LGTM.
    </>
  ),

  noFavorites: "No favorite LGTM.",
  successToGenerate: "Successfully generated LGTM.",
  failedToGenerate: "Failed to generate LGTM.",
  successToSend: "Sent.",
  failedToSend: "Failed to send.",
  fileTooLarge: "File size is too large.",
  unsupportedImageFormat: "Unsupported image format.",
  copiedToClipboard: "Copied to clipboard.",
  rateLimitReached:
    "The request limit has been reached.\nPlease try again later.",

  aboutTrafficAnalysisTool: "About Traffic Analysis Tool",
  aboutTrafficAnalysisToolContent: (
    <>
      On this site, we utilize the access analysis tool &quot;Google
      Analytics&quot; provided by Google. Google Analytics uses cookies for the
      purpose of collecting traffic data. This traffic data is collected
      anonymously and does not identify individuals. You can opt out of data
      collection by disabling cookies in your browser settings. Please refer to
      the{" "}
      <a
        className="text-primary-main underline"
        href="https://marketingplatform.google.com/about/analytics/terms/us/"
        target="_blank"
        rel="noopener noreferrer"
      >
        Terms of Service of Google Analytics
      </a>{" "}
      for more details regarding this policy.
    </>
  ),

  aboutPrivacyPolicyChange: "About Privacy Policy Change",
  aboutPrivacyPolicyChangeContent:
    "This site complies with Japanese laws and regulations applicable to personal information, and strives to review and improve the contents of this policy as appropriate. The updated privacy policy, when amended, will always be disclosed on this page.",

  usagePrecautionsItems: [
    "All responsibilities regarding the images generated through the use of this service shall be borne by the user. In the event that a third party suffers damage in relation to images created by the user, the operators of this service shall not bear any responsibility on behalf of the user.",
    "Images generated through the use of this service will be published on the internet. Please be mindful of the copyrights of the original images and related laws. Do not create images that are against public order and morals or are illegal. Such images, along with others deemed inappropriate by the operators, may be deleted without notice.",
    "Please refrain from sending an excessive number of requests and causing strain on the service.",
    "Furthermore, in the event that malicious use is identified, specific users may be prohibited from accessing the service without notice.",
  ],
} as const;

export default en;
