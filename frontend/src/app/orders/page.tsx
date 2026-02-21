/* eslint-disable */
"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Edit, Eye } from "lucide-react";
import { toast } from "react-hot-toast";
import { useAuth } from "../context/AuthContext";

interface Order {
  id: string;
  date: string;
  status: "Pending" | "Processing" | "Shipped" | "Delivered" | "Cancelled";
  total: number;
  items: {
    name: string;
    quantity: number;
    price: number;
  }[];
  shippingAddress: {
    street: string;
    city: string;
    state: string;
    zipCode: string;
  };
}

export default function OrdersPage() {
  const { user } = useAuth();
  const [orders] = useState<Order[]>([
    {
      id: "001",
      date: "2024-11-15",
      status: "Pending",
      total: 299.97,
      items: [
        { name: "Wireless Headphones", quantity: 1, price: 199.99 },
        { name: "Organic T-Shirt", quantity: 2, price: 49.99 }
      ],
      shippingAddress: {
        street: "1234",
        city: "Bangalore",
        state: "Karnataka", 
        zipCode: "560123"
      }
    },
    {
      id: "002", 
      date: "2024-11-14",
      status: "Shipped",
      total: 15.99,
      items: [
        { name: "Gourmet Coffee Beans", quantity: 1, price: 15.99 }
      ],
      shippingAddress: {
        street: "15342",
        city: "Mumbai",
        state: "Maharastra",
        zipCode: "564321"
      }
    }
  ]);

  const [selectedOrder, setSelectedOrder] = useState<Order | null>(null);
  const [editMode, setEditMode] = useState(false);
  const [editedAddress, setEditedAddress] = useState({
    street: "",
    city: "",
    state: "",
    zipCode: ""
  });

  const handleEdit = (order: Order) => {
    setSelectedOrder(order);
    setEditedAddress(order.shippingAddress);
    setEditMode(true);
  };

  const handleSaveEdit = () => {
    // In real app, make API call to update order
    toast.success("Order updated successfully");
    setEditMode(false);
  };

  const getStatusColor = (status: Order["status"]) => {
    const colors = {
      Pending: "bg-yellow-100 text-yellow-800",
      Processing: "bg-blue-100 text-blue-800",
      Shipped: "bg-purple-100 text-purple-800",
      Delivered: "bg-green-100 text-green-800",
      Cancelled: "bg-red-100 text-red-800"
    };
    return colors[status];
  };

  return (
    <div className="min-h-screen bg-gray-50 p-6">
      <div className="max-w-6xl mx-auto">
        <h1 className="text-2xl font-semibold mb-6">My Orders</h1>
        <div className="grid gap-4">
          {orders.map((order) => (
            <Card key={order.id}>
              <CardContent className="p-6">
                <div className="flex justify-between items-start">
                  <div>
                    <p className="font-medium">Order #{order.id}</p>
                    <p className="text-sm text-gray-500">{order.date}</p>
                  </div>
                  <div className="flex items-center gap-2">
                    <Badge className={getStatusColor(order.status)}>
                      {order.status}
                    </Badge>
                    <Dialog>
                      <DialogTrigger asChild>
                        <Button variant="ghost" size="icon">
                          <Eye className="h-4 w-4" />
                        </Button>
                      </DialogTrigger>
                      <DialogContent>
                        <DialogHeader>
                          <DialogTitle>Order Details #{order.id}</DialogTitle>
                        </DialogHeader>
                        <div className="space-y-4">
                          <div>
                            <h3 className="font-medium mb-2">Items</h3>
                            {order.items.map((item, index) => (
                              <div key={index} className="flex justify-between text-sm">
                                <span>{item.name} x{item.quantity}</span>
                                <span>${item.price.toFixed(2)}</span>
                              </div>
                            ))}
                          </div>
                          <div>
                            <h3 className="font-medium mb-2">Shipping Address</h3>
                            <p className="text-sm">
                              {order.shippingAddress.street}<br />
                              {order.shippingAddress.city}, {order.shippingAddress.state} {order.shippingAddress.zipCode}
                            </p>
                          </div>
                          <div className="flex justify-between font-medium">
                            <span>Total</span>
                            <span>${order.total.toFixed(2)}</span>
                          </div>
                        </div>
                      </DialogContent>
                    </Dialog>
                    {order.status === "Pending" && (
                      <Dialog>
                        <DialogTrigger asChild>
                          <Button variant="ghost" size="icon" onClick={() => handleEdit(order)}>
                            <Edit className="h-4 w-4" />
                          </Button>
                        </DialogTrigger>
                        <DialogContent>
                          <DialogHeader>
                            <DialogTitle>Edit Order #{order.id}</DialogTitle>
                          </DialogHeader>
                          <div className="space-y-4">
                            <div>
                              <h3 className="font-medium mb-2">Update Shipping Address</h3>
                              <div className="space-y-2">
                                <Input
                                  placeholder="Street Address"
                                  value={editedAddress.street}
                                  onChange={(e) => setEditedAddress({
                                    ...editedAddress,
                                    street: e.target.value
                                  })}
                                />
                                <Input
                                  placeholder="City"
                                  value={editedAddress.city}
                                  onChange={(e) => setEditedAddress({
                                    ...editedAddress,
                                    city: e.target.value
                                  })}
                                />
                                <Input
                                  placeholder="State"
                                  value={editedAddress.state}
                                  onChange={(e) => setEditedAddress({
                                    ...editedAddress,
                                    state: e.target.value
                                  })}
                                />
                                <Input
                                  placeholder="ZIP Code"
                                  value={editedAddress.zipCode}
                                  onChange={(e) => setEditedAddress({
                                    ...editedAddress,
                                    zipCode: e.target.value
                                  })}
                                />
                              </div>
                            </div>
                            <Button className="w-full" onClick={handleSaveEdit}>
                              Save Changes
                            </Button>
                          </div>
                        </DialogContent>
                      </Dialog>
                    )}
                  </div>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    </div>
  );
}