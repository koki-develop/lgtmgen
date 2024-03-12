import { ModelsImage, ModelsLGTM } from "@/lib/generated/api";
import { useI18n } from "@/providers/I18nProvider";

export type SearchImagePanelProps = {
  images: ModelsImage[];
  query: string;

  onChangeQuery: (query: string) => void;
  onSearched: (images: ModelsImage[]) => void;
  onGenerated: (lgtm: ModelsLGTM) => void;
};

export default function SearchImagePanel() {
  const { t } = useI18n();

  return (
    <div className="flex justify-center">
      <p className="text-gray-500">{t.searchImagesNotAvailable}</p>
    </div>
  );
}
