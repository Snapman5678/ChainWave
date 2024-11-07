"use client";

import { createContext, useContext, useState, useEffect } from "react";
import axios from "axios";

interface User {
  username: string;
  email: string;
  roles?: string[];
}

interface AuthContextType {
  isAuthenticated: boolean;
  user: User | null;
  setUser: (user: User | null) => void;
  logout: () => void;
  updateUserRoles: (roles: string[]) => Promise<void>;
}

const AuthContext = createContext<AuthContextType>({
  isAuthenticated: false,
  user: null,
  setUser: () => {},
  logout: () => {},
  updateUserRoles: async () => {},
});

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    // Check localStorage for user data on mount
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
  }, []);

  const logout = () => {
    setUser(null);
    localStorage.removeItem("user");
    localStorage.removeItem("authToken");
  };

  const updateUserRoles = async (roles: string[]) => {
    try {
      const response = await axios.post('http://localhost:8000/api/user/roles', { roles });
      if (response.data.success) {
        setUser(prev => ({ ...prev!, roles }));
        localStorage.setItem('user', JSON.stringify({ ...user, roles }));
      }
    } catch (error) {
      console.error('Error updating roles:', error);
    }
  };

  return (
    <AuthContext.Provider
      value={{
        isAuthenticated: !!user,
        user,
        setUser,
        logout,
        updateUserRoles,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => useContext(AuthContext);
