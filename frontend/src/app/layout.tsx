import Image from "next/image";
import type { Metadata } from "next";
import localFont from "next/font/local";
import "./styles/globals.css";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "ChainWave",
  description: "Crowdsourced Supply Chain Management System",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
        <header className="bg-gray-800 text-white p-4">
          <nav className="container mx-auto flex justify-between items-center">
            <a href="/" className="flex items-center text-xl font-bold">
              <Image
                src="/images/icon.svg" 
                alt="Icon"
                width={60} 
                height={60} 
                className="mr-2"
              />&nbsp;
              ChainWave
            </a>
            <div>
              <a href="/auth/register" className="mr-4 hover:underline">Register</a>
              <a href="/auth/login" className="hover:underline">Login</a>
            </div>
          </nav>
        </header>
        <main className="min-h-screen">{children}</main>
      </body>
    </html>
  );
}
