import { Api } from "@/lib/generated/api";

export const api = new Api({
  baseUrl: "http://localhost:8080", // TODO: from env
  baseApiParams: {
    cache: "no-cache",
    format: "json",
  },
});
