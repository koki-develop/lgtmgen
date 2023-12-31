declare namespace NodeJS {
  interface ProcessEnv {
    readonly NEXT_PUBLIC_STAGE: string;
    readonly NEXT_PUBLIC_API_BASE_URL: string;
    readonly NEXT_PUBLIC_IMAGES_BASE_URL: string;
    readonly NEXT_PUBLIC_GA_MEASUREMENT_ID: string;
    // Sentry
    readonly NEXT_PUBLIC_SENTRY_ORG: string;
    readonly NEXT_PUBLIC_SENTRY_PROJECT: string;
    readonly NEXT_PUBLIC_SENTRY_DSN: string;
  }
}
