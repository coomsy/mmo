import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

import { NavLinks } from '@/app/_components/ui/nav-links'

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Based titte",
  description: "Based desc",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <NavLinks />
        <main>{children}</main>
      </body>
    </html>
  );
}
