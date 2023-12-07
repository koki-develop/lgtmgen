import GoogleAnalytics from "@/components/util/GoogleAnalytics";
import { i18n } from "@/lib/i18n";
import Providers from "@/providers/Providers";
import "@fontsource/archivo-black";
import type { Metadata } from "next";
import Link from "next/link";
import { AiOutlineGithub } from "react-icons/ai";
import Header from "./Header";
import "./global.css";

export async function generateMetadata({
  params: { locale },
}: {
  params: { locale: string };
}): Promise<Metadata> {
  const t = i18n(locale);

  return {
    title: t.app,
    description: t.description,
    themeColor: "#1E90FF",

    icons: {
      icon: "/favicon.ico",
      apple: "/logo192.png",
    },

    openGraph: {
      title: t.app,
      siteName: t.app,
      description: t.description,
      locale: locale === "en" ? "en_US" : "ja_JP",
      images: {
        url: "/card.png",
        width: 600,
        height: 314,
      },
      type: "website",
    },

    twitter: {
      card: "summary_large_image",
      site: "@koki_develop",
      images: {
        url: "/card.png",
        width: 600,
        height: 314,
      },
    },
  };
}

export default function RootLayout({
  params: { locale },
  children,
}: {
  params: { locale: string };
  children: React.ReactNode;
}) {
  const t = i18n(locale);

  return (
    <html lang={locale}>
      <body className="bg-primary-light text-gray-600">
        <Header />

        <Providers locale={locale}>
          <GoogleAnalytics />
          <main className="container mx-auto p-4 py-8">{children}</main>
        </Providers>

        <footer className="flex flex-col items-center gap-4 pb-8">
          <ul className="flex flex-col items-center justify-center gap-2">
            <li>
              <a
                href="https://github.com/koki-develop/lgtmgen"
                target="_blank"
                rel="noopener noreferrer"
              >
                <AiOutlineGithub size={30} />
              </a>
            </li>
            <li>
              <Link href={`/${locale}/privacy`}>{t.privacyPolicy}</Link>
            </li>
            <li>
              <Link href={`/${locale}/usage-precautions`}>
                {t.usagePrecautions}
              </Link>
            </li>
            <li>
              <a
                href={t.feedbackLink}
                target="_blank"
                rel="noopener noreferrer"
              >
                {t.feedback}
              </a>
            </li>
          </ul>

          <small>&copy; 2023 Koki Sato</small>
        </footer>
      </body>
    </html>
  );
}
