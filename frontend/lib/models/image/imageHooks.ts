import { api } from "@/lib/api";
import { useCallback, useState } from "react";

export const useSearchImages = () => {
  const [searching, setSearching] = useState<boolean>(false);

  const searchImages = useCallback(async (query: string) => {
    setSearching(true);
    return await api.v1
      .imagesList({ q: query })
      .then((response) => {
        if (!response.ok) throw response.error;
        return response.data;
      })
      .finally(() => {
        setSearching(false);
      });
  }, []);

  return {
    searching,
    searchImages,
  };
};
