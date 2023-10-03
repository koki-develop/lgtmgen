import { ModelsImage, ModelsLGTM, ServiceErrCode } from "@/lib/generated/api";
import { useI18n } from "@/providers/I18nProvider";
import { useCallback, useState } from "react";
import { MagnifyingGlassIcon } from "@heroicons/react/24/outline";
import Form from "@/components/util/Form";
import LgtmPreview from "./LgtmPreview";
import { useGenerateLgtm } from "@/lib/models/lgtm/lgtmHooks";
import ImageCard from "./ImageCard";
import { useSearchImages } from "@/lib/models/image/imageHooks";

export type SearchImagePanelProps = {
  images: ModelsImage[];
  query: string;

  onChangeQuery: (query: string) => void;
  onSearched: (images: ModelsImage[]) => void;
  onGenerated: (lgtm: ModelsLGTM) => void;
};

export default function SearchImagePanel({
  images,
  query,

  onChangeQuery,
  onSearched,
  onGenerated,
}: SearchImagePanelProps) {
  const { t } = useI18n();
  const { generateLgtm, generating } = useGenerateLgtm();
  const { searchImages, searching } = useSearchImages();

  const [url, setUrl] = useState<string | null>(null);

  const handleChangeQuery = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      onChangeQuery(event.target.value);
    },
    [onChangeQuery],
  );

  const handleClosePreview = useCallback(() => {
    setUrl(null);
  }, []);

  const handleSearch = useCallback(async () => {
    const trimmedQuery = query.trim();
    if (trimmedQuery === "") return;

    const images = await searchImages(trimmedQuery);
    onSearched(images);
  }, [query, searchImages, onSearched]);

  const handleClickImage = useCallback(async (image: ModelsImage) => {
    setUrl(image.url);
  }, []);

  const handleGenerate = useCallback(async () => {
    if (!url) return;

    const lgtm = await generateLgtm({ url });
    if (lgtm) {
      onGenerated(lgtm);
      handleClosePreview();
    }
  }, [url, generateLgtm, onGenerated, handleClosePreview]);

  return (
    <>
      <LgtmPreview
        src={url}
        onCancel={handleClosePreview}
        generating={generating}
        onGenerate={handleGenerate}
      />

      <div className="flex flex-col gap-4">
        <Form onSubmit={handleSearch}>
          <div className="flex overflow-hidden rounded bg-white shadow-md">
            <input
              className="flex-grow p-4 outline-none"
              disabled={searching}
              type="search"
              placeholder={t.keyword}
              value={query}
              onChange={handleChangeQuery}
            />
            <button
              className="button-primary px-4"
              type="submit"
              disabled={searching}
              onClick={handleSearch}
            >
              <MagnifyingGlassIcon className="h-6 w-6" />
            </button>
          </div>
        </Form>

        <div>
          <ul className="grid grid-cols-4 gap-4">
            {images.map((image) => (
              <li key={image.url}>
                <ImageCard
                  className="f-full"
                  src={image.url}
                  alt={image.title}
                  onClick={() => handleClickImage(image)}
                />
              </li>
            ))}
          </ul>
        </div>
      </div>
    </>
  );
}
