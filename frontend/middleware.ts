import Negotiator from "negotiator";
import { NextRequest, NextResponse } from "next/server";

const supportedLocales = ["ja", "en"];
const defaultLocale = "ja";

const extractLocale = (headers: Negotiator.Headers) => {
  return (
    new Negotiator({ headers }).language(supportedLocales) ?? defaultLocale
  );
};

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;

  const excluded: (string | RegExp)[] = [
    // static files
    "/favicon.ico",
    "/robots.txt",
    "/sitemap.xml",
    "/card.png",
    "/logo192.png",
    "/logo512.png",
  ];
  if (excluded.some((path) => pathname.match(path))) return;

  const pathnameHasLocale = supportedLocales.some(
    (locale) => pathname.startsWith(`/${locale}/`) || pathname === `/${locale}`,
  );
  if (pathnameHasLocale) return;

  const headers = {
    "accept-language": request.headers.get("accept-language") ?? "",
  };
  const locale = extractLocale(headers);

  request.nextUrl.pathname = `/${locale}${pathname}`;
  return NextResponse.redirect(request.nextUrl);
}

export const config = {
  matcher: ["/((?!_next).*)"],
};
