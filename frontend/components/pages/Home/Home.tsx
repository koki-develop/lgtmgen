import React from "react";
import { api } from "@/lib/api";
import Main from "./Main";

const perPage = 2; // TODO: -> 40

export type HomeProps = {
  params: {
    locale: string;
  };
};

export default async function Home({ params: { locale } }: HomeProps) {
  const listing = api.v1.lgtmsList({ limit: perPage });
  const randomListing = api.v1.lgtmsList({ limit: perPage, random: true });

  const [listingResponse, randomListingResponse] = await Promise.all([
    listing,
    randomListing,
  ]);
  if (!listingResponse.ok) throw listingResponse.error;
  if (!randomListingResponse.ok) throw randomListingResponse.error;

  return (
    <Main
      locale={locale}
      initialData={listingResponse.data}
      initialRandomData={randomListingResponse.data}
      perPage={perPage}
    />
  );
}
