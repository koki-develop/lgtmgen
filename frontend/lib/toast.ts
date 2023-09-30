import { useCallback } from "react";
import toast from "react-hot-toast";

export const useToast = () => {
  const enqueueToast = useCallback(
    (message: string, type: "success" | "error" = "success") => {
      switch (type) {
        case "success":
          toast.success(message);
          break;
        case "error":
          toast.error(message);
          break;
      }
    },
    [],
  );

  return {
    enqueueToast,
  };
};
