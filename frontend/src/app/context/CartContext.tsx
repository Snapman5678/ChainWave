"use client";

import React, { createContext, useContext, useState, useEffect } from "react";
import { Product } from "../marketplace/types";

interface CartItem extends Product {
  quantity: number;
  availableStock: number;
}

interface CartContextType {
  items: CartItem[];
  addToCart: (product: Product, quantity: number) => boolean;
  removeFromCart: (productId: string) => void;
  updateQuantity: (productId: string, quantity: number) => void;
  clearCart: () => void;
  total: number;
}

const CartContext = createContext<CartContextType>({
  items: [],
  addToCart: () => false,
  removeFromCart: () => {},
  updateQuantity: () => {},
  clearCart: () => {},
  total: 0,
});

export function CartProvider({ children }: { children: React.ReactNode }) {
  const [items, setItems] = useState<CartItem[]>([]);

  useEffect(() => {
    try {
      const saved = localStorage.getItem("cart");
      if (saved) {
        setItems(JSON.parse(saved));
      }
    } catch (error) {
      console.error("Error loading cart from localStorage:", error);
    }
  }, []);

  useEffect(() => {
    try {
      localStorage.setItem("cart", JSON.stringify(items));
    } catch (error) {
      console.error("Error saving cart to localStorage:", error);
    }
  }, [items]);

  const addToCart = (product: Product, quantity: number = 1): boolean => {
    let added = false;
    setItems((current) => {
      const exists = current.find((item) => item.id === product.id);
      const maxQuantity = product.quantity;
      if (exists) {
        if (exists.quantity >= maxQuantity) {
          // Cannot add more than available stock
          return current;
        }
        const newQuantity = Math.min(exists.quantity + quantity, maxQuantity);
        added = true;
        return current.map((item) =>
          item.id === product.id
            ? { ...item, quantity: newQuantity }
            : item
        );
      }
      const initialQuantity = Math.min(quantity, maxQuantity);
      added = true;
      return [...current, { ...product, quantity: initialQuantity, availableStock: maxQuantity }];
    });
    return added;
  };

  const removeFromCart = (productId: string) => {
    setItems((current) => current.filter((item) => item.id !== productId));
  };

  const updateQuantity = (productId: string, quantity: number) => {
    if (quantity <= 0) {
      removeFromCart(productId);
      return;
    }
    setItems((current) =>
      current.map((item) =>
        item.id === productId ? { ...item, quantity } : item
      )
    );
  };

  const clearCart = () => {
    setItems([]);
  };

  const total = items.reduce((sum, item) => sum + item.price * item.quantity, 0);

  return (
    <CartContext.Provider
      value={{
        items,
        addToCart,
        removeFromCart,
        updateQuantity,
        clearCart,
        total,
      }}
    >
      {children}
    </CartContext.Provider>
  );
}

export const useCart = () => useContext(CartContext);