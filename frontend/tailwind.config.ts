import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./providers/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  plugins: [require("@headlessui/tailwindcss")],
  theme: {
    extend: {
      colors: {
        primary: {
          main: "#1E90FF",
          light: "#E8EEF2",
          dark: "#0070df",
        },
        favorite: {
          main: "#f48fb1",
          dark: "#C2185B",
          light: "#F8BBD0",
        },
        report: {
          main: "#FF9800",
          dark: "#F57B00",
        },
      },
    },
  },
};
export default config;
