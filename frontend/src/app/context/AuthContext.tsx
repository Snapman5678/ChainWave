"use client";

import { createContext, useContext, useEffect, useState, ReactNode } from "react";
import axios from "axios";

type User = {
  id: string;
  email: string;
  token: string;
  roles: string[];  // Added roles from the other context
  username: string;
};

export interface AuthContextType {
  user: User | null;
  setUser: (user: User | null) => void;
  isAuthenticated: boolean;
  login: (userData: User) => void;
  logout: () => void;
  updateToken: (newToken: string) => void; // Add this new method
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      const userData = JSON.parse(storedUser);
      // Fetch fresh roles when restoring the session
      const fetchRoles = async () => {
        try {
          const response = await axios.get('http://localhost:8000/api/role', {
            headers: { Authorization: `Bearer ${userData.token}` }
          });
          
          if (response.data.roles && response.data.roles.length > 0) {
            const roles = response.data.roles.map((role: any) => 
              typeof role === 'string' ? role : role.name || String(role)
            );
            
            // Update user with fresh roles
            const updatedUser = {
              ...userData,
              roles: roles
            };
            setUser(updatedUser);
            localStorage.setItem('user', JSON.stringify(updatedUser));
          }
        } catch (error) {
          console.error("Error fetching roles:", error);
          // If there's an error (like expired token), clear the session
          if (axios.isAxiosError(error) && error.response?.status === 401) {
            logout();
          }
        }
      };

      setUser(userData);
      fetchRoles();
    }
  }, []);

  useEffect(() => {
    const user = JSON.parse(localStorage.getItem("user") || "{}");
    if (user.token) {
      axios.defaults.headers.common["Authorization"] = `Bearer ${user.token}`;
    }
  }, []);

  // Add new updateToken function
  const updateToken = (newToken: string) => {
    if (user) {
      const updatedUser = {
        ...user,
        token: newToken
      };
      setUser(updatedUser);
      localStorage.setItem('user', JSON.stringify(updatedUser));
      axios.defaults.headers.common["Authorization"] = `Bearer ${newToken}`;
    }
  };

  // Configure axios interceptors to automatically handle token updates
  useEffect(() => {
    const requestInterceptor = axios.interceptors.request.use(
      (config) => {
        const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
        if (currentUser.token) {
          config.headers["Authorization"] = `Bearer ${currentUser.token}`;
        }
        return config;
      },
      (error) => Promise.reject(error)
    );

    const responseInterceptor = axios.interceptors.response.use(
      (response) => {
        // Check for token in response
        if (response.data?.token) {
          updateToken(response.data.token);
        }
        return response;
      },
      (error) => Promise.reject(error)
    );

    // Cleanup interceptors on unmount
    return () => {
      axios.interceptors.request.eject(requestInterceptor);
      axios.interceptors.response.eject(responseInterceptor);
    };
  }, [user]); // Add user as dependency

  const login = (userData: User) => {
    setUser(userData);
    localStorage.setItem('user', JSON.stringify(userData));
  };

  const logout = () => {
    setUser(null);
    localStorage.removeItem('user');
    localStorage.removeItem('authToken'); // Add this line
    localStorage.removeItem('userId');    // Add this line
    localStorage.removeItem('username');  // Add this line
    delete axios.defaults.headers.common["Authorization"]; // Add this line
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        setUser,
        isAuthenticated: !!user,
        login,
        logout,
        updateToken, // Add this to the context value
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

