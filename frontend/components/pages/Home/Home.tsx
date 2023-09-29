import React from "react";
import { api } from "@/lib/api";
import Main from "./Main";

const lgtmsLimit = 2; // TODO: -> 40

export default async function Home() {
  const resp = await api.v1.lgtmsList({ limit: lgtmsLimit });
  if (!resp.ok) throw resp.error;
  const lgtms = resp.data;

  return (
    <div>
      <Main initialData={lgtms} />
    </div>
  );
}
