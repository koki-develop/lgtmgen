import React from "react";
import { api } from "@/lib/api";
import Main from "./Main";

const perPage = 20; // TODO: -> 40

export default async function Home() {
  const resp = await api.v1.lgtmsList({ limit: perPage });
  if (!resp.ok) throw resp.error;
  const lgtms = resp.data;

  return (
    <div>
      <Main initialData={lgtms} perPage={perPage} />
    </div>
  );
}
