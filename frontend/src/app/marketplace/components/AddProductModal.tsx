"use client";

import React, { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useAuth } from "@/app/context/AuthContext";

interface Product {
  id?: string;
  name: string;
  description: string;
  price: number;
  category: string;
  quantity: number;
  imageUrl: string;
  businessName: string;
  contactEmail: string;
  contactPhone: string;
}

interface AddProductModalProps {
  isOpen: boolean;
  onClose: () => void;
  onAddProduct: (productData: Omit<Product, "id">) => void;
}

export default function AddProductModal({
  isOpen,
  onClose,
  onAddProduct,
}: AddProductModalProps) {
  const { user } = useAuth();
  const [productData, setProductData] = useState<Product>({
    name: "",
    description: "",
    price: 0,
    category: "",
    quantity: 0,
    imageUrl: "",
    businessName: user?.username || "",
    contactEmail: user?.email || "",
    contactPhone: "",
  });

  const categories = [
    "Electronics",
    "Clothing",
    "Food",
    "Home",
    "Office",
    "Other",
  ];

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    // Basic validation
    if (
      !productData.name ||
      !productData.category ||
      !productData.price ||
      !productData.quantity ||
      !productData.imageUrl ||
      !productData.contactPhone
    ) {
      alert("Please fill in all required fields");
      return;
    }

    // Convert price and quantity to numbers and validate
    const price = Number(productData.price);
    const quantity = Number(productData.quantity);

    if (isNaN(price) || price <= 0) {
      alert("Please enter a valid price");
      return;
    }

    if (isNaN(quantity) || quantity < 0) {
      alert("Please enter a valid quantity");
      return;
    }

    onAddProduct({
      ...productData,
      price,
      quantity,
    });

    // Reset form
    setProductData({
      name: "",
      description: "",
      price: 0,
      category: "",
      quantity: 0,
      imageUrl: "",
      businessName: user?.username || "",
      contactEmail: user?.email || "",
      contactPhone: "",
    });
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[600px]">
        <DialogHeader>
          <DialogTitle>Add New Product</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">Product Name</label>
              <Input
                placeholder="Product Name"
                value={productData.name}
                onChange={(e) =>
                  setProductData({ ...productData, name: e.target.value })
                }
                required
              />
            </div>
            <div className="space-y-2">
              <label className="text-sm font-medium">Category</label>
              <Select
                value={productData.category}
                onValueChange={(value) =>
                  setProductData({ ...productData, category: value })
                }
              >
                <SelectTrigger>
                  <SelectValue placeholder="Select category" />
                </SelectTrigger>
                <SelectContent>
                  {categories.map((category) => (
                    <SelectItem key={category} value={category.toLowerCase()}>
                      {category}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium">Description</label>
            <Textarea
              placeholder="Product description"
              value={productData.description}
              onChange={(e) =>
                setProductData({ ...productData, description: e.target.value })
              }
              required
            />
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">Price</label>
              <Input
                type="number"
                placeholder="Price"
                value={productData.price}
                onChange={(e) =>
                  setProductData({ ...productData, price: Number(e.target.value) })
                }
                required
              />
            </div>
            <div className="space-y-2">
              <label className="text-sm font-medium">Quantity</label>
              <Input
                type="number"
                placeholder="Quantity"
                value={productData.quantity}
                onChange={(e) =>
                  setProductData({ ...productData, quantity: Number(e.target.value) })
                }
                required
              />
            </div>
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium">Image URL</label>
            <Input
              placeholder="Image URL"
              value={productData.imageUrl}
              onChange={(e) =>
                setProductData({ ...productData, imageUrl: e.target.value })
              }
              required
            />
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium">Contact Phone</label>
            <Input
              placeholder="Contact Phone"
              value={productData.contactPhone}
              onChange={(e) =>
                setProductData({ ...productData, contactPhone: e.target.value })
              }
              required
            />
          </div>

          <Button type="submit" className="w-full">
            Add Product
          </Button>
        </form>
      </DialogContent>
    </Dialog>
  );
}
