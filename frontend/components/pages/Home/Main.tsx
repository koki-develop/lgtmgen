"use client";

import React, { useCallback, useState } from "react";
import { api } from "@/lib/api";
import { ModelsImage, ModelsLGTM } from "@/lib/generated/api";
import LgtmPanel from "./LgtmPanel";
import SearchImagePanel from "./SearchImagePanel";
import { Tab } from "@headlessui/react";
import { i18n } from "@/lib/i18n";
import clsx from "clsx";

export type MainProps = {
  locale: string;
  initialData: ModelsLGTM[];
  perPage: number;
};

export default function Main({ locale, initialData, perPage }: MainProps) {
  const t = i18n(locale);

  /*
   * LGTM
   */

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

  const handleGenerated = useCallback((lgtm: ModelsLGTM) => {
    setLgtms((prev) => [lgtm, ...prev]);
  }, []);

  /*
   * SearchImage
   */

  const [images, setImages] = useState<ModelsImage[]>([]);

  const handleSearch = useCallback(async (query: string) => {
    const resp = await api.v1.imagesList({ q: query });
    if (!resp.ok) throw resp.error;
    setImages(resp.data);
  }, []);

  // Render
  return (
    <div>
      <Tab.Group>
        <Tab.List className="mb-4 flex rounded-t bg-white shadow-md">
          {[t.lgtm, t.searchImage, t.favorite].map((label) => (
            <Tab
              key={label}
              className={clsx(
                "flex-grow py-2 outline-none transition",
                "border-b-2 border-b-white",
                "hover:border-b-gray-100 hover:bg-gray-100",
                "ui-selected:border-b-primary-main ui-selected:font-semibold ui-selected:text-primary-main ui-not-selected:text-gray-400",
              )}
            >
              {label}
            </Tab>
          ))}
        </Tab.List>

        <Tab.Panels>
          <Tab.Panel>
            <LgtmPanel
              lgtms={lgtms}
              hasNextPage={hasNextPage}
              onLoadMore={handleLoadMore}
              onUploaded={handleGenerated}
            />
          </Tab.Panel>

          <Tab.Panel>
            <SearchImagePanel
              images={images}
              onSearch={handleSearch}
              onGenerated={handleGenerated}
            />
          </Tab.Panel>
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
}
