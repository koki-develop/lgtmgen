import { i18n } from "@/lib/i18n";
import "./global.css";
import "@fontsource/archivo-black";
import type { Metadata } from "next";
import Link from "next/link";
import Providers from "@/providers/Providers";

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

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
      <body className="bg-primary-light">
        <header className="shadow bg-primary-main py-2 px-4">
          <h1
            className="text-3xl text-white"
            style={{ fontFamily: "Archivo Black" }}
          >
            <Link href={`/${locale}`}>{t.app}</Link>
          </h1>
        </header>

        <Providers locale={locale}>
          <main className="container mx-auto p-4 py-8">{children}</main>
        </Providers>

        <footer className="flex flex-col items-center">
          <small>&copy; 2023 Koki Sato</small>
        </footer>
      </body>
    </html>
  );
}
