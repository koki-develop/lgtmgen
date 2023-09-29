export const lgtmUrl = (id: string): string => {
  return `${process.env.NEXT_PUBLIC_IMAGES_BASE_URL}/${id}`;
};

export const fileToDataUrl = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = (error) => reject(error);
    reader.readAsDataURL(file);
  });
};

export const dataUrlToBase64 = (dataUrl: string): string => {
  return dataUrl.split(",")[1];
};
