"use client";

import { useEffect, useState } from "react";
import { useCart } from "../context/CartContext";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useAuth } from "../context/AuthContext";
import { useRouter } from "next/navigation";

// Add Product interface
interface Product {
  id: string;
  name: string;
  price: number;
}

// Declare Razorpay on window object
declare global {
  interface Window {
    Razorpay: any;
  }
}

interface CheckoutItem extends Product {
  quantity: number;
}

interface DirectPurchase {
  items: CheckoutItem[];
  isDirectPurchase: boolean;
}

export default function CheckoutPage() {
  const { items, total, clearCart } = useCart();
  const { user } = useAuth();
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [checkoutItems, setCheckoutItems] = useState<CheckoutItem[]>([]);
  const [isDirectPurchase, setIsDirectPurchase] = useState(false);
  const [checkoutTotal, setCheckoutTotal] = useState(0);
  const [address, setAddress] = useState({
    street: "",
    city: "",
    state: "",
    zipCode: "",
  });

  useEffect(() => {
    // Load Razorpay SDK
    const script = document.createElement("script");
    script.src = "https://checkout.razorpay.com/v1/checkout.js";
    script.async = true;
    document.body.appendChild(script);

    // Check if this is a direct purchase from "Buy Now"
    const directPurchaseData = sessionStorage.getItem('directPurchase');
    if (directPurchaseData) {
      const { items: directItems, isDirectPurchase }: DirectPurchase = JSON.parse(directPurchaseData);
      setCheckoutItems(directItems);
      setIsDirectPurchase(isDirectPurchase);
      setCheckoutTotal(directItems.reduce((sum, item) => sum + item.price * item.quantity, 0));
      // Clear the session storage
      sessionStorage.removeItem('directPurchase');
    } else {
      // Use cart items if no direct purchase
      setCheckoutItems(items);
      setCheckoutTotal(total);
      setIsDirectPurchase(false);
    }
  }, [items, total]);

  const handlePayment = async () => {
    if (!checkoutItems.length) {
      alert("No items to checkout");
      return;
    }

    if (!address.street || !address.city || !address.state || !address.zipCode) {
      alert("Please fill in all address fields");
      return;
    }

    setLoading(true);
    try {
      // TODO: Replace with your actual API call to create order
      const orderData = {
        amount: checkoutTotal * 100, // Razorpay expects amount in paise
        currency: "INR",
        receipt: "order_" + Math.random().toString(36).substr(2, 9),
      };

      // Initialize Razorpay options
      const options = {
        key: process.env.NEXT_PUBLIC_RAZORPAY_KEY_ID, // Your Razorpay Key ID
        amount: orderData.amount,
        currency: orderData.currency,
        name: "ChainWave",
        description: "Purchase from ChainWave",
        order_id: orderData.receipt,
        handler: function (response: any) {
          // Handle successful payment
          console.log(response);
          if (!isDirectPurchase) {
            clearCart();
          }
          alert("Payment successful!");
          router.push("/marketplace"); // Redirect to marketplace after successful payment
        },
        prefill: {
          name: user?.username,
          email: user?.email,
        },
        theme: {
          color: "#000000",
        },
      };

      const razorpay = new window.Razorpay(options);
      razorpay.open();
    } catch (error) {
      console.error("Payment failed:", error);
      alert("Payment failed. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 p-6">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-2xl font-semibold mb-6">Checkout</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="space-y-4">
            <h2 className="text-xl font-semibold">Shipping Address</h2>
            <Input
              placeholder="Street Address"
              value={address.street}
              onChange={(e) =>
                setAddress({ ...address, street: e.target.value })
              }
            />
            <Input
              placeholder="City"
              value={address.city}
              onChange={(e) => setAddress({ ...address, city: e.target.value })}
            />
            <Input
              placeholder="State"
              value={address.state}
              onChange={(e) => setAddress({ ...address, state: e.target.value })}
            />
            <Input
              placeholder="ZIP Code"
              value={address.zipCode}
              onChange={(e) =>
                setAddress({ ...address, zipCode: e.target.value })
              }
            />
          </div>
          <div>
            <h2 className="text-xl font-semibold mb-4">Order Summary</h2>
            <div className="bg-white p-4 rounded-lg shadow-sm">
              {checkoutItems.map((item) => (
                <div
                  key={item.id}
                  className="flex justify-between items-center mb-2"
                >
                  <span>
                    {item.name} x {item.quantity}
                  </span>
                  <span>${(item.price * item.quantity).toFixed(2)}</span>
                </div>
              ))}
              <div className="border-t mt-4 pt-4">
                <div className="flex justify-between items-center font-bold">
                  <span>Total</span>
                  <span>${checkoutTotal.toFixed(2)}</span>
                </div>
              </div>
              <Button
                className="w-full mt-4"
                onClick={handlePayment}
                disabled={loading}
              >
                {loading ? "Processing..." : "Pay Now"}
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
