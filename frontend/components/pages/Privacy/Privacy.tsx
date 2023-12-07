"use client";

import { useI18n } from "@/providers/I18nProvider";

export default function Privacy() {
  const { t } = useI18n();

  return (
    <div className="flex flex-col gap-4">
      <h2 className="text-3xl">{t.privacyPolicy}</h2>

      <div>
        <h3 className="mb-2 text-xl">{t.aboutTrafficAnalysisTool}</h3>
        <div>{t.aboutTrafficAnalysisToolContent}</div>
      </div>

      <div>
        <h3 className="mb-2 text-xl">{t.aboutPrivacyPolicyChange}</h3>
        <div>{t.aboutPrivacyPolicyChangeContent}</div>
      </div>
    </div>
  );
}
