import React from "react";
import { api } from "@/lib/api";
import Main from "./Main";

const perPage = 20; // TODO: -> 40

export type HomeProps = {
  params: {
    locale: string;
  };
};

export default async function Home({ params: { locale } }: HomeProps) {
  const resp = await api.v1.lgtmsList({ limit: perPage });
  if (!resp.ok) throw resp.error;
  const lgtms = resp.data;

  return <Main locale={locale} initialData={lgtms} perPage={perPage} />;
}
