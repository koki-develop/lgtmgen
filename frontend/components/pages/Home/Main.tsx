"use client";

import React, { useCallback, useState } from "react";
import { api } from "@/lib/api";
import { ModelsLGTM } from "@/lib/generated/api";
import LgtmPanel from "./LgtmPanel";

export type MainProps = {
  initialData: ModelsLGTM[];
};

const perPage = 2; // TODO: -> 40

export default function Main({ initialData }: MainProps) {
  const [lgtms, setLgtms] = useState<ModelsLGTM[]>(initialData);
  const [hasNextPage, setHasNextPage] = useState<boolean>(
    initialData.length === perPage,
  );

  const handleLoadMore = useCallback(async () => {
    const after = lgtms.slice(-1)[0]?.id;
    const resp = await api.v1.lgtmsList({ after, limit: perPage });
    if (!resp.ok) throw resp.error;
    setLgtms((prev) => [...prev, ...resp.data]);
    setHasNextPage(resp.data.length === perPage);
  }, [lgtms]);

  const handleUploaded = useCallback((lgtm: ModelsLGTM) => {
    setLgtms((prev) => [lgtm, ...prev]);
  }, []);

  return (
    <div>
      <LgtmPanel
        lgtms={lgtms}
        hasNextPage={hasNextPage}
        onLoadMore={handleLoadMore}
        onUploaded={handleUploaded}
      />
    </div>
  );
}
