import React, { useCallback } from "react";

export type FormProps = React.ComponentPropsWithoutRef<"form">;

export default function Form({ onSubmit, ...props }: FormProps) {
  const handleSubmit = useCallback(
    (event: React.FormEvent<HTMLFormElement>) => {
      event.preventDefault();
      onSubmit?.(event);
    },
    [onSubmit],
  );

  return <form onSubmit={handleSubmit} {...props} />;
}
