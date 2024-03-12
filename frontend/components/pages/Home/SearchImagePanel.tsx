import { useI18n } from "@/providers/I18nProvider";

export type SearchImagePanelProps = {};

export default function SearchImagePanel() {
  const { t } = useI18n();

  return (
    <div className="flex justify-center">
      <p className="text-gray-500">{t.searchImagesNotAvailable}</p>
    </div>
  );
}
