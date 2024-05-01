import { api } from "@/lib/api";
import Main from "./Main";

const perPage = 20;

export type HomeProps = {
  params: {
    locale: string;
  };
};

export default async function Home({ params: { locale } }: HomeProps) {
  // const lang = ["ja", "en"].includes(locale) ? locale : "ja";
  const newsListing = api.v1.newsList({ locale });
  // const lgtmsListing = api.v1.lgtmsList({ limit: perPage });
  // const lgtmsRandomListing = api.v1.lgtmsList({ limit: perPage, random: true });
  // const categoriesListing = api.v1.categoriesList({ lang });

  const [
    newsListResponse,
    // lgtmsListResponse,
    // lgtmListRandomResponse,
    // categoriesListResponse,
  ] = await Promise.all([
    newsListing,
    // lgtmsListing,
    // lgtmsRandomListing,
    // categoriesListing,
  ]);
  if (!newsListResponse.ok) throw newsListResponse.error;
  // if (!lgtmsListResponse.ok) throw lgtmsListResponse.error;
  // if (!lgtmListRandomResponse.ok) throw lgtmListRandomResponse.error;
  // if (!categoriesListResponse.ok) throw categoriesListResponse.error;

  return (
    <Main
      locale={locale}
      newsList={newsListResponse.data}
      initialData={[]}
      initialRandomData={[]}
      categories={[]}
      perPage={perPage}
    />
  );
}
