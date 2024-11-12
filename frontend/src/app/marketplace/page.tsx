"use client";

import React, { useState } from "react";
import { PlusCircle, Search } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import AddProductModal from "./components/AddProductModal";
import { useAuth } from "../context/AuthContext";
import ProductDetailModal from "./components/ProductDetailModal";
import { useCart } from "../context/CartContext";
import { Product } from "./types";
import { useRouter } from "next/navigation";
import Image from "next/image";
import { toast } from "react-hot-toast";

// Add this interface near the top of the file
interface Role {
  name?: string;
  [key: string]: any;
}

export default function MarketplacePage() {
  const { user } = useAuth();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [searchQuery, setSearchQuery] = useState("");
  const [selectedCategory, setSelectedCategory] = useState<string>("all");
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);
  const [addedToCart, setAddedToCart] = useState<{[key: string]: boolean}>({});
  const router = useRouter();

  // TODO: Replace with API call to fetch products
  const [products, setProducts] = useState<Product[]>([
    {
      id: "1",
      name: "Sample Product",
      description: "This is a sample product description",
      price: 99.99,
      category: "electronics",
      quantity: 10,
      imageUrl: "/images/placeholder.jpg", // Updated path
      businessName: "Sample Business",
      contactEmail: "contact@example.com",
      contactPhone: "123-456-7890",
    },
    {
      id: "2",
      name: "Wireless Headphones",
      description: "High-quality wireless headphones with noise cancellation",
      price: 199.99,
      category: "electronics",
      quantity: 5,
      imageUrl: "/images/placeholder.jpg", // Updated path
      businessName: "AudioHub",
      contactEmail: "support@audiohub.com",
      contactPhone: "987-654-3210",
    },
    {
      id: "3",
      name: "Organic T-Shirt",
      description:
        "Comfortable organic cotton t-shirt available in various sizes",
      price: 29.99,
      category: "clothing",
      quantity: 20,
      imageUrl: "/images/placeholder.jpg", // Updated path
      businessName: "EcoWear",
      contactEmail: "info@ecowear.com",
      contactPhone: "555-123-4567",
    },
    {
      id: "4",
      name: "Gourmet Coffee Beans",
      description: "Premium roasted coffee beans sourced from the finest farms",
      price: 15.99,
      category: "food",
      quantity: 50,
      imageUrl: "/images/placeholder.jpg", // Updated path
      businessName: "CoffeeElite",
      contactEmail: "contact@coffeeelite.com",
      contactPhone: "444-555-6666",
    },
    {
      id: "5",
      name: "Smart LED Lamp",
      description: "Adjustable smart LED lamp with multiple color settings",
      price: 49.99,
      category: "home",
      quantity: 15,
      imageUrl: "/images/placeholder.jpg", // Updated path
      businessName: "BrightHome",
      contactEmail: "sales@brighthouse.com",
      contactPhone: "222-333-4444",
    },
  ]);

  const categories = [
    "Electronics",
    "Clothing",
    "Food",
    "Home",
    "Office",
    "Other",
  ];

  const handleAddProduct = async (newProduct: Omit<Product, "id">) => {
    try {
      // TODO: Replace with actual API call to create product
      const productWithId = {
        ...newProduct,
        id: Math.random().toString(36).substr(2, 9), // This will be replaced by server-generated ID
      };

      // TODO: Add error handling for API response
      setProducts([...products, productWithId]);
      setIsModalOpen(false);
    } catch (error) {
      console.error("Failed to add product:", error);
      alert("Failed to add product. Please try again.");
    }
  };

  // TODO: Add useEffect to fetch products when component mounts
  // useEffect(() => {
  //   const fetchProducts = async () => {
  //     try {
  //       const response = await fetch('/api/products');
  //       const data = await response.json();
  //       setProducts(data);
  //     } catch (error) {
  //       console.error('Failed to fetch products:', error);
  //     }
  //   };
  //   fetchProducts();
  // }, []);

  const filteredProducts = products.filter((product) => {
    const matchesSearch = product.name
      .toLowerCase()
      .includes(searchQuery.toLowerCase());
    const matchesCategory =
      selectedCategory === "all"
        ? true
        : product.category.toLowerCase() === selectedCategory.toLowerCase();
    return matchesSearch && matchesCategory;
  });

  const { addToCart } = useCart();

  const handleQuickAddToCart = (product: Product) => {
    if (!user) {
      router.push("/auth/login");
      return;
    }
    const success = addToCart(product, 1);
    if (success) {
      toast.success("Product added to cart.");
      setAddedToCart({ ...addedToCart, [product.id]: true });
      setTimeout(() => {
        setAddedToCart({ ...addedToCart, [product.id]: false });
      }, 2000);
    } else {
      toast.error(`Only ${product.quantity} units available in stock.`);
    }
  };

  const handleQuickBuyNow = (product: Product) => {
    if (!user) {
      router.push("/auth/login");
      return;
    }
    sessionStorage.setItem('directPurchase', JSON.stringify({
      items: [{...product, quantity: 1}],
      isDirectPurchase: true
    }));
    router.push("/checkout");
  };

  // Update the business role check with proper typing
  const isBusinessAdmin = user?.roles?.some((role: string | Role) => {
    if (typeof role === 'string') {
      return role.toLowerCase() === 'business';
    }
    return (role.name || '').toLowerCase() === 'business';
  });

  return (
    <div className="min-h-screen bg-gray-50 p-6">
      <div className="max-w-7xl mx-auto">
        <div className="flex flex-col gap-6">
          {/* Header */}
          <div className="flex justify-between items-center">
            <h1 className="text-2xl font-semibold text-gray-900">
              Marketplace
            </h1>
            {isBusinessAdmin && (  // Use the new check
              <Button onClick={() => setIsModalOpen(true)}>
                <PlusCircle className="h-4 w-4 mr-2" />
                Add Product
              </Button>
            )}
          </div>

          {/* Filters */}
          <div className="flex gap-4">
            <div className="flex-1">
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" />
                <Input
                  placeholder="Search products..."
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                  className="pl-10"
                />
              </div>
            </div>
            <Select
              value={selectedCategory}
              onValueChange={setSelectedCategory}
            >
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="All Categories" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">All Categories</SelectItem>
                {categories.map((category) => (
                  <SelectItem key={category} value={category.toLowerCase()}>
                    {category}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          {/* Products Grid */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {filteredProducts.map((product) => (
              <div
                key={product.id}
                className="bg-white rounded-lg shadow-sm overflow-hidden border border-gray-200"
              >
                <div
                  className="cursor-pointer"
                  onClick={() => {
                    setSelectedProduct(product);
                    setIsDetailModalOpen(true);
                  }}
                >
                  <Image
                    src={product.imageUrl.startsWith('http') ? product.imageUrl : product.imageUrl}
                    alt={product.name}
                    width={200}
                    height={200}
                    className="w-full h-48 object-cover rounded-t-lg"
                    onError={(e: any) => {
                      e.target.src = "/images/placeholder.jpg" // Updated fallback image path
                    }}
                    unoptimized={product.imageUrl.startsWith('http')} // Only use unoptimized for external URLs
                  />
                  <div className="p-4">
                    <h3 className="font-semibold text-lg mb-2">
                      {product.name}
                    </h3>
                    <p className="text-gray-600 text-sm mb-2">
                      {product.description.substring(0, 100)}...
                    </p>
                    <div className="flex justify-between items-center">
                      <span className="font-bold text-lg">
                        ${product.price.toFixed(2)}
                      </span>
                      <span className="text-sm text-gray-500">
                        Stock: {product.quantity}
                      </span>
                    </div>
                  </div>
                </div>
                <div className="p-4 border-t flex gap-2">
                  <Button
                    variant="outline"
                    className="flex-1"
                    onClick={(e) => {
                      e.stopPropagation();
                      handleQuickAddToCart(product);
                    }}
                    disabled={addedToCart[product.id]}
                  >
                    {addedToCart[product.id] ? "Added to Cart" : "Add to Cart"}
                  </Button>
                  <Button
                    className="flex-1"
                    onClick={(e) => {
                      e.stopPropagation();
                      handleQuickBuyNow(product);
                    }}
                  >
                    Buy Now
                  </Button>
                </div>
              </div>
            ))}
          </div>
        </div>

        <AddProductModal
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          onAddProduct={handleAddProduct}
        />

        <ProductDetailModal
          isOpen={isDetailModalOpen}
          onClose={() => setIsDetailModalOpen(false)}
          product={selectedProduct}
        />
      </div>
    </div>
  );
}
