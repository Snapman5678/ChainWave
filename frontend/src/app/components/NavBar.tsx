"use client";

import Image from "next/image";
import { useState, useRef, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuth } from "../context/AuthContext";
import { useCart } from "../context/CartContext";
import { ShoppingCart, Store } from "lucide-react";
import Link from "next/link";

export default function NavBar() {
  const { isAuthenticated, user, logout } = useAuth();
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);
  const router = useRouter();
  const { items } = useCart();
  const totalItems = items.reduce((sum, item) => sum + item.quantity, 0);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setDropdownOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  const handleLogout = () => {
    logout();
    router.push("/auth/login");
    setDropdownOpen(false);
  };

  return (
    <nav className="container mx-auto flex justify-between items-center">
      <a href="/" className="flex items-center text-xl font-bold">
        <Image
          src="/images/icon.svg"
          alt="Icon"
          width={60}
          height={60}
          className="mr-2"
        />
        &nbsp; ChainWave
      </a>
      <div className="flex items-center space-x-6">
        <Link 
          href="/marketplace" 
          className="flex items-center text-white hover:text-gray-300 transition-colors"
        >
          <Store className="w-6 h-6" />
        </Link>
        {isAuthenticated && user ? (
          <div className="relative flex items-center space-x-6" ref={dropdownRef}>
            <Link href="/cart" className="flex items-center text-white">
              <div className="relative">
                <ShoppingCart className="w-6 h-6" />
                {totalItems > 0 && (
                  <span className="absolute -top-2 -right-2 bg-red-500 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center">
                    {totalItems}
                  </span>
                )}
              </div>
            </Link>
            <button
              onClick={() => setDropdownOpen(!dropdownOpen)}
              className="flex items-center space-x-1 text-white hover:text-gray-300"
            >
              <span className="text-lg">{user.username}</span>
              <svg
                className="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M19 9l-7 7-7-7"
                />
              </svg>
            </button>
            {dropdownOpen && (
              <div className="absolute right-0 top-full mt-2 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-50">
                <div className="py-1" role="menu">
                  <Link
                    href="/profile"
                    className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    Profile Settings
                  </Link>
                  <Link
                    href="/orders"
                    className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    My Orders
                  </Link>
                  <button
                    onClick={handleLogout}
                    className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    Logout
                  </button>
                </div>
              </div>
            )}
          </div>
        ) : (
          <div className="flex items-center space-x-6">
            <button 
              onClick={() => router.push('/auth/register')} 
              className="px-6 py-2 rounded-md border border-white/20 text-white hover:bg-white/10 transition-colors"
            >
              Register
            </button>
            <button 
              onClick={() => router.push('/auth/login')} 
              className="px-6 py-2 rounded-md border border-white/20 text-white hover:bg-white/10 transition-colors"
            >
              Login
            </button>
          </div>
        )}
      </div>
    </nav>
  );
}
