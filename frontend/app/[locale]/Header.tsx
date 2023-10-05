"use client";

import { useI18n } from "@/providers/I18nProvider";
import { Menu } from "@headlessui/react";
import { ChevronDownIcon, LanguageIcon } from "@heroicons/react/24/outline";
import clsx from "clsx";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

export type HeaderProps = {
  locale: string;
};

export default function Header() {
  const { t, locale } = useI18n();

  const pathname = usePathname();

  return (
    <header className="flex justify-center bg-primary-main shadow">
      <div className="container flex px-4 py-2">
        <h1
          className="flex-grow text-2xl text-white sm:text-3xl"
          style={{ fontFamily: "Archivo Black" }}
        >
          <Link href={`/${locale}`}>{t.app}</Link>
        </h1>
        <div className="relative flex justify-center px-2">
          <Menu>
            <Menu.Button className="flex items-center text-white">
              <ChevronDownIcon className="h-3 w-3" />
              <LanguageIcon className="h-6 w-6" />
            </Menu.Button>
            <Menu.Items className="absolute -left-8 top-4 flex flex-col divide-y overflow-hidden rounded bg-white text-gray-600 shadow-md">
              <Menu.Item>
                {({ active }) => (
                  <Link
                    href={pathname.replace(/^\/en/, "/ja")}
                    className={clsx("px-4 py-2 transition", {
                      "bg-gray-200": active,
                    })}
                  >
                    日本語
                  </Link>
                )}
              </Menu.Item>
              <Menu.Item>
                {({ active }) => (
                  <Link
                    href={pathname.replace(/^\/ja/, "/en")}
                    className={clsx("px-4 py-2 transition", {
                      "bg-gray-200": active,
                    })}
                  >
                    English
                  </Link>
                )}
              </Menu.Item>
            </Menu.Items>
          </Menu>
        </div>
      </div>
    </header>
  );
}
