"use client";

import { usePathname, useSearchParams } from "next/navigation";
import Script from "next/script";
import React, { useEffect, useState } from "react";

export default function GoogleAnalytics() {
  const [rendered, setRendered] = useState<boolean>(false);
  const pathname = usePathname();
  const searchParams = useSearchParams();

  useEffect(() => {
    if (!rendered) {
      setRendered(true);
      return;
    }
    const pagePath = pathname + "?" + searchParams.toString();
    if (process.env.NEXT_PUBLIC_STAGE !== "prd") {
      console.log("GoogleAnalytics:", pagePath);
      return;
    }

    window.gtag("config", process.env.NEXT_PUBLIC_GA_MEASUREMENT_ID, {
      page_path: pagePath,
    });
  }, [pathname, rendered, searchParams]);

  if (process.env.NEXT_PUBLIC_STAGE !== "prd") {
    return null;
  }

  return (
    <>
      <Script
        src={`https://www.googletagmanager.com/gtag/js?id=${process.env.NEXT_PUBLIC_GA_MEASUREMENT_ID}`}
      />
      <Script id="google-analytics">
        {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', '${process.env.NEXT_PUBLIC_GA_MEASUREMENT_ID}');
        `}
      </Script>
    </>
  );
}
