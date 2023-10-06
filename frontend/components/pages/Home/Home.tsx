import React from "react";
import { api } from "@/lib/api";
import Main from "./Main";

const perPage = 20;

export type HomeProps = {
  params: {
    locale: string;
  };
};

export default async function Home({ params: { locale } }: HomeProps) {
  const newsListing = api.v1.newsList({ locale });
  const lgtmsListing = api.v1.lgtmsList({ limit: perPage });
  const lgtmsRandomListing = api.v1.lgtmsList({ limit: perPage, random: true });

  const [newsListResponse, lgtmsListResponse, lgtmListRandomResponse] =
    await Promise.all([newsListing, lgtmsListing, lgtmsRandomListing]);
  if (!newsListResponse.ok) throw newsListResponse.error;
  if (!lgtmsListResponse.ok) throw lgtmsListResponse.error;
  if (!lgtmListRandomResponse.ok) throw lgtmListRandomResponse.error;

  return (
    <Main
      locale={locale}
      newsList={newsListResponse.data}
      initialData={lgtmsListResponse.data}
      initialRandomData={lgtmListRandomResponse.data}
      perPage={perPage}
    />
  );
}
