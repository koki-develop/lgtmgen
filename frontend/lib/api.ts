import { Api } from "@/lib/generated/api";

export const api = new Api({
  baseUrl: process.env.NEXT_PUBLIC_API_BASE_URL,
  baseApiParams: {
    cache: "no-cache",
    format: "json",
  },
});
