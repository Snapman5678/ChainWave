"use client";

import { createContext, useContext } from 'react';
import { useAuth } from './AuthContext';

interface ApiContextType {
  get: (url: string) => Promise<any>;
  post: (url: string, data: any) => Promise<any>;
  put: (url: string, data: any) => Promise<any>;
  delete: (url: string) => Promise<any>;
}

const ApiContext = createContext<ApiContextType | undefined>(undefined);

export function ApiProvider({ children }: { children: React.ReactNode }) {
  const { user } = useAuth();
  const baseUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8000';

  const headers = {
    'Content-Type': 'application/json',
    ...(user?.token && { 'Authorization': `Bearer ${user.token}` })
  };

  const api = {
    get: async (url: string) => {
      const response = await fetch(`${baseUrl}${url}`, { headers });
      if (!response.ok) throw new Error('API request failed');
      return response.json();
    },

    post: async (url: string, data: any) => {
      const response = await fetch(`${baseUrl}${url}`, {
        method: 'POST',
        headers,
        body: JSON.stringify(data)
      });
      if (!response.ok) throw new Error('API request failed');
      return response.json();
    },

    put: async (url: string, data: any) => {
      const response = await fetch(`${baseUrl}${url}`, {
        method: 'PUT',
        headers,
        body: JSON.stringify(data)
      });
      if (!response.ok) throw new Error('API request failed');
      return response.json();
    },

    delete: async (url: string) => {
      const response = await fetch(`${baseUrl}${url}`, {
        method: 'DELETE',
        headers
      });
      if (!response.ok) throw new Error('API request failed');
      return response.json();
    }
  };

  return (
    <ApiContext.Provider value={api}>
      {children}
    </ApiContext.Provider>
  );
}

export const useApi = () => {
  const context = useContext(ApiContext);
  if (context === undefined) {
    throw new Error('useApi must be used within an ApiProvider');
  }
  return context;
};
