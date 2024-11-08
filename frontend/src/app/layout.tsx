import type { Metadata } from "next";
import localFont from "next/font/local";
import "./styles/globals.css";
import { AuthProvider } from "./context/AuthContext";
import NavBar from './components/NavBar';
import { Toaster } from "react-hot-toast";
import { CartProvider } from "./context/CartContext";
import { ApiProvider } from "./context/ApiContext";

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
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <AuthProvider>
          <ApiProvider>
            <CartProvider>
              <header className="bg-gray-800 text-white p-4">
                <NavBar />
              </header>
              <main className="min-h-screen">{children}</main>
              <Toaster position="bottom-right" />
            </CartProvider>
          </ApiProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
