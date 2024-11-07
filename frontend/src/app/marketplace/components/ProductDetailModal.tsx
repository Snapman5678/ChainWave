"use client";

import React, { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Product } from "../types";
import { useCart } from "@/app/context/CartContext";
import { Plus, Minus } from "lucide-react";

interface ProductDetailModalProps {
  isOpen: boolean;
  onClose: () => void;
  product: Product | null;
}

export default function ProductDetailModal({
  isOpen,
  onClose,
  product,
}: ProductDetailModalProps) {
  const { addToCart } = useCart();
  const [quantity, setQuantity] = useState(1);
  const [isAdded, setIsAdded] = useState(false);
  const [error, setError] = useState<string | null>(null);

  if (!product) return null;

  const handleQuantityChange = (newQuantity: number) => {
    if (newQuantity > product.quantity) {
      setError(`Only ${product.quantity} units available in stock`);
      return;
    }
    setError(null);
    setQuantity(newQuantity);
  };

  const handleAddToCart = () => {
    if (quantity > product.quantity) {
      setError(`Only ${product.quantity} units available in stock`);
      return;
    }
    addToCart(product, quantity);
    setIsAdded(true);
    setTimeout(() => setIsAdded(false), 2000);
  };

  const handleBuyNow = () => {
    addToCart(product, quantity);
    window.location.href = "/checkout";
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[800px]">
        <DialogHeader>
          <DialogTitle>{product.name}</DialogTitle>
        </DialogHeader>
        <div className="grid grid-cols-2 gap-6">
          <div>
            <img
              src={product.imageUrl}
              alt={product.name}
              className="w-full rounded-lg"
            />
          </div>
          <div className="space-y-4">
            <p className="text-2xl font-bold">${product.price.toFixed(2)}</p>
            <p className="text-gray-600">{product.description}</p>
            <div className="border-t pt-4">
              <p>Category: {product.category}</p>
              <p>Stock: {product.quantity}</p>
              <p>Seller: {product.businessName}</p>
              <p>Contact: {product.contactPhone}</p>
              <p>Email: {product.contactEmail}</p>
            </div>
            <div className="flex items-center gap-4 mb-4">
              <span>Quantity:</span>
              <div className="flex items-center gap-2">
                <Button
                  variant="outline"
                  size="icon"
                  onClick={() => handleQuantityChange(Math.max(1, quantity - 1))}
                >
                  <Minus className="h-4 w-4" />
                </Button>
                <span className="w-8 text-center">{quantity}</span>
                <Button
                  variant="outline"
                  size="icon"
                  onClick={() => handleQuantityChange(quantity + 1)}
                >
                  <Plus className="h-4 w-4" />
                </Button>
              </div>
            </div>
            {error && <p className="text-red-500 text-sm">{error}</p>}
            <div className="flex gap-4 pt-4">
              <Button onClick={handleAddToCart} disabled={isAdded}>
                {isAdded ? "Added to Cart" : "Add to Cart"}
              </Button>
              <Button variant="secondary" onClick={handleBuyNow}>
                Buy Now
              </Button>
            </div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
