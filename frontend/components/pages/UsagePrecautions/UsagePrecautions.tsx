"use client";

import { useI18n } from "@/providers/I18nProvider";
import React from "react";

export default function UsagePrecautions() {
  const { t } = useI18n();

  return (
    <div className="flex flex-col gap-4">
      <h2 className="text-3xl">{t.usagePrecautions}</h2>

      <ul className="list-disc pl-4">
        {t.usagePrecautionsItems.map((item, i) => (
          <li key={i} className="">
            {item}
          </li>
        ))}
      </ul>
    </div>
  );
}
