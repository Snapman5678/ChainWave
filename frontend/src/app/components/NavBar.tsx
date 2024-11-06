"use client";

import Image from "next/image";
import { useState, useRef, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuth } from "../context/AuthContext";

export default function NavBar() {
  const { isAuthenticated, user, logout } = useAuth();
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);
  const router = useRouter();

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
      <div>
        {isAuthenticated && user ? (
          <div className="relative flex items-center" ref={dropdownRef}>
            <span className="text-white mr-2">{user.username}</span>
            <button
              onClick={() => setDropdownOpen(!dropdownOpen)}
              className="flex items-center text-white hover:text-gray-300"
            >
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
                  <a
                    href="/profile/change-username"
                    className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    Change Username
                  </a>
                  <a
                    href="/profile/change-email"
                    className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    Change Email
                  </a>
                  <a
                    href="/profile/change-password"
                    className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    Change Password
                  </a>
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
          <>
            <a href="/auth/register" className="mr-4 hover:underline">
              Register
            </a>
            <a href="/auth/login" className="hover:underline">
              Login
            </a>
          </>
        )}
      </div>
    </nav>
  );
}
