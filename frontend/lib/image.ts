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

export const resizeDataUrl = async (
  dataUrl: string,
  sideLength: number,
  type: "image/jpeg" | "image/png",
): Promise<string> => {
  // load image
  const image = new Image();
  image.src = dataUrl;
  await new Promise<void>((resolve) => {
    image.onload = () => resolve();
  });

  // create canvas
  const canvas = document.createElement("canvas");
  const ctx = canvas.getContext("2d");
  if (!ctx) {
    throw new Error("Failed to get canvas context");
  }

  // keep aspect ratio
  const ratio = Math.max(sideLength / image.width, sideLength / image.height);
  const width = image.width * ratio;
  const height = image.height * ratio;

  // resize
  canvas.width = width;
  canvas.height = height;
  ctx.drawImage(image, 0, 0, width, height);

  // convert to data url
  return canvas.toDataURL(type);
};
