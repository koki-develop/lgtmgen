import "./global.css";
import "@fontsource/archivo-black";
import type { Metadata } from "next";

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
  return (
    <html lang={locale}>
      <body>{children}</body>
    </html>
  );
}
