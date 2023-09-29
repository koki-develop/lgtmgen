import { api } from "@/lib/api";
import { ModelsImage, ModelsLGTM, ServiceErrCode } from "@/lib/generated/api";
import { useCallback, useState } from "react";

export type SearchImagePanelProps = {
  images: ModelsImage[];

  onSearch: (query: string) => void;
  onGenerated: (lgtm: ModelsLGTM) => void;
};

export default function SearchImagePanel({
  images,
  onSearch,
  onGenerated,
}: SearchImagePanelProps) {
  const [query, setQuery] = useState<string>("");

  const handleChangeQuery = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      setQuery(event.currentTarget.value);
    },
    [],
  );

  const handleSearch = useCallback(() => {
    const trimmedQuery = query.trim();
    if (trimmedQuery === "") return;
    onSearch(trimmedQuery);
  }, [query, onSearch]);

  const handleClickImage = useCallback(
    async (image: ModelsImage) => {
      const response = await api.v1.lgtmsCreate({ url: image.url });
      if (response.ok) {
        onGenerated(response.data);
        return;
      }

      switch (response.error.code) {
        case ServiceErrCode.ErrCodeUnsupportedImageFormat:
          alert("Unsupported image format");
          break;
        case ServiceErrCode.ErrCodeInternalServerError:
          alert("Internal server error");
          break;
        default:
          throw response.error;
      }
    },
    [onGenerated],
  );

  return (
    <div>
      <div>
        <input type="search" value={query} onChange={handleChangeQuery} />
        <button onClick={handleSearch}>Search</button>
      </div>
      <div>
        <ul>
          {images.map((image) => (
            <li key={image.url}>
              <img
                src={image.url}
                alt={image.title}
                style={{ width: 400, maxHeight: 400, objectFit: "contain" }}
                onClick={() => handleClickImage(image)}
              />
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
