"use client";

import React, { createContext, useContext, useState, useEffect } from "react";
import { Product } from "../marketplace/types";

interface CartItem extends Product {
  quantity: number;
}

interface CartContextType {
  items: CartItem[];
  addToCart: (product: Product, quantity: number) => void;
  removeFromCart: (productId: string) => void;
  updateQuantity: (productId: string, quantity: number) => void;
  clearCart: () => void;
  total: number;
}

const CartContext = createContext<CartContextType>({
  items: [],
  addToCart: () => {},
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

  const addToCart = (product: Product, quantity: number = 1) => {
    setItems((current) => {
      const exists = current.find((item) => item.id === product.id);
      if (exists) {
        // Check if new total quantity exceeds available stock
        const newQuantity = exists.quantity + quantity;
        if (newQuantity > product.quantity) {
          alert(`Cannot add more items. Only ${product.quantity} units available in stock.`);
          return current;
        }
        return current.map((item) =>
          item.id === product.id
            ? { ...item, quantity: newQuantity }
            : item
        );
      }
      // Check if initial quantity exceeds available stock
      if (quantity > product.quantity) {
        alert(`Cannot add ${quantity} items. Only ${product.quantity} units available in stock.`);
        return current;
      }
      return [...current, { ...product, quantity }];
    });
  };

  const removeFromCart = (productId: string) => {
    setItems((current) => current.filter((item) => item.id !== productId));
  };

  const updateQuantity = (productId: string, quantity: number) => {
    if (quantity <= 0) {
      removeFromCart(productId);
      return;
    }
    
    setItems((current) => {
      const item = current.find((i) => i.id === productId);
      if (!item) return current;
      
      // Check if new quantity exceeds available stock
      if (quantity > item.quantity) {
        alert(`Cannot add more items. Only ${item.quantity} units available in stock.`);
        return current;
      }
      
      return current.map((item) =>
        item.id === productId ? { ...item, quantity } : item
      );
    });
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