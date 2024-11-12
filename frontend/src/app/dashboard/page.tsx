/* eslint-disable */
"use client";

import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";
import { useAuth } from "../context/AuthContext";
import RoleSelectionForm from "../components/RoleSelectionForm"; // Add this import
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Package, TrendingUp, Truck, Factory, Search } from "lucide-react";

export default function DashboardPage() {
  const router = useRouter();
  const { user, setUser, updateToken } = useAuth();
  const [userType, setUserType] = useState<string>("customer");
  const [showRoleForm, setShowRoleForm] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [orderId, setOrderId] = useState("");
  const [orders] = useState([
    {
      id: "123",
      status: "In Transit",
      date: "2024-03-15",
      details: "Electronics Shipment",
    },
    {
      id: "124",
      status: "Pending",
      date: "2024-03-16",
      details: "Raw Materials",
    },
  ]);

  useEffect(() => {
    const checkUserRoles = async () => {
      if (!user) {
        router.push("/auth/login");
        return;
      }

      setIsLoading(true);
      try {
        const response = await axios.get('http://localhost:8000/api/role', {
          headers: { Authorization: `Bearer ${user.token}` }
        });

        // Update token if provided in response
        if (response.data.token) {
          updateToken(response.data.token);
        }

        if (response.data.roles && response.data.roles.length > 0) {
          // Extract role names as strings
          const roles = response.data.roles.map((role: any) => 
            typeof role === 'string' ? role : role.name || String(role)
          );

          // Only update state if roles have changed
          if (JSON.stringify(user.roles) !== JSON.stringify(roles)) {
            setUser({
              ...user,
              roles: roles
            });
          }
          setUserType(roles[0].toLowerCase());
          setShowRoleForm(false);
        } else {
          // No roles found, show role selection form
          setShowRoleForm(true);
        }
      } catch (error: any) {
        console.error("Error checking roles:", error);
        if (error.response?.status === 401) {
          // Token expired or invalid
          router.push("/auth/login");
        }
      } finally {
        setIsLoading(false);
      }
    };

    checkUserRoles();
  // Update the dependency array
  }, [user?.token]);

  // Show loading state
  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900 mx-auto"></div>
          <p className="mt-4 text-gray-600">Loading...</p>
        </div>
      </div>
    );
  }

  // Show role selection form when no roles exist
  if (showRoleForm) {
    return <RoleSelectionForm onSubmit={(roles: string[]) => {  // Add type annotation
      if (roles.length > 0) {
        setUser({
          ...user!,
          roles: roles
        });
        setUserType(roles[0]);
        setShowRoleForm(false);
      }
    }} />;
  }

  const renderCustomerDashboard = () => (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      <Card className="col-span-full">
        <CardHeader>
          <CardTitle>Order Tracking</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex gap-2">
            <Input
              placeholder="Enter Order ID"
              value={orderId}
              onChange={(e) => setOrderId(e.target.value)}
              className="max-w-xs"
            />
            <Button variant="outline">
              <Search className="h-4 w-4 mr-2" />
              Track Order
            </Button>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Active Orders</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            {orders.map((order) => (
              <div key={order.id} className="p-4 border rounded-lg">
                <div className="font-medium">Order #{order.id}</div>
                <div className="text-sm text-gray-500">{order.status}</div>
                <div className="text-sm text-gray-500">{order.date}</div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Order Statistics</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-2">
            <div className="flex justify-between">
              <span>Total Orders</span>
              <span className="font-medium">24</span>
            </div>
            <div className="flex justify-between">
              <span>In Transit</span>
              <span className="font-medium">3</span>
            </div>
            <div className="flex justify-between">
              <span>Completed</span>
              <span className="font-medium">20</span>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );

  const renderBusinessAdminDashboard = () => (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      <Card>
        <CardHeader>
          <CardTitle>Business Overview</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-2">
            <div className="flex justify-between">
              <span>Total Revenue</span>
              <span className="font-medium">$124,500</span>
            </div>
            <div className="flex justify-between">
              <span>Active Orders</span>
              <span className="font-medium">15</span>
            </div>
            <div className="flex justify-between">
              <span>Suppliers</span>
              <span className="font-medium">8</span>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Supply Chain Status</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="p-4 border rounded-lg bg-green-50">
              <div className="font-medium">Normal Operations</div>
              <div className="text-sm text-gray-500">
                All systems functioning
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );

  const renderTransporterDashboard = () => (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      <Card>
        <CardHeader>
          <CardTitle>Delivery Schedule</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            {orders.map((order) => (
              <div key={order.id} className="p-4 border rounded-lg">
                <div className="font-medium">Delivery #{order.id}</div>
                <div className="text-sm text-gray-500">{order.details}</div>
                <div className="text-sm text-gray-500">ETA: {order.date}</div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Vehicle Status</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-2">
            <div className="flex justify-between">
              <span>Available Vehicles</span>
              <span className="font-medium">5</span>
            </div>
            <div className="flex justify-between">
              <span>In Transit</span>
              <span className="font-medium">3</span>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );

  const renderSupplierDashboard = () => (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      <Card>
        <CardHeader>
          <CardTitle>Inventory Status</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-2">
            <div className="flex justify-between">
              <span>Total Items</span>
              <span className="font-medium">1,240</span>
            </div>
            <div className="flex justify-between">
              <span>Low Stock Alerts</span>
              <span className="font-medium text-red-500">3</span>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Purchase Orders</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="p-4 border rounded-lg">
              <div className="font-medium">New Orders</div>
              <div className="text-sm text-gray-500">5 pending approval</div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );

  const renderDashboard = () => {
    switch (userType) {
      case "customer":
        return renderCustomerDashboard();
      case "business":
        return renderBusinessAdminDashboard();
      case "transporter":
        return renderTransporterDashboard();
      case "supplier":
        return renderSupplierDashboard();
      default:
        return renderCustomerDashboard();
    }
  };

  const capitalizeRole = (role: string | undefined) => {
    if (typeof role !== 'string' || !role) return '';
    return role.charAt(0).toUpperCase() + role.slice(1);
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center mb-6 gap-4">
          <div>
            <h1 className="text-2xl font-semibold text-gray-900">Dashboard</h1>
            <p className="mt-1 text-sm text-gray-500">
              Welcome to your supply chain management dashboard
            </p>
          </div>

          {userType && (
            <Select 
              value={userType || "customer"} 
              onValueChange={(value) => setUserType(value)}
            >
              <SelectTrigger className="w-[200px] bg-white shadow-sm">
                <div className="flex items-center">
                  {userType === "customer" && <Package className="h-4 w-4 mr-2" />}
                  {userType === "business" && <TrendingUp className="h-4 w-4 mr-2" />}
                  {userType === "transporter" && <Truck className="h-4 w-4 mr-2" />}
                  {userType === "supplier" && <Factory className="h-4 w-4 mr-2" />}
                  {capitalizeRole(userType)}
                </div>
              </SelectTrigger>
              <SelectContent className="bg-white">
                <SelectItem value="customer" className="hover:bg-gray-100">
                  <div className="flex items-center">
                    <Package className="h-4 w-4 mr-2" />
                    Customer
                  </div>
                </SelectItem>
                <SelectItem value="business" className="hover:bg-gray-100">
                  <div className="flex items-center">
                    <TrendingUp className="h-4 w-4 mr-2" />
                    Business Admin
                  </div>
                </SelectItem>
                <SelectItem value="transporter" className="hover:bg-gray-100">
                  <div className="flex items-center">
                    <Truck className="h-4 w-4 mr-2" />
                    Transporter
                  </div>
                </SelectItem>
                <SelectItem value="supplier" className="hover:bg-gray-100">
                  <div className="flex items-center">
                    <Factory className="h-4 w-4 mr-2" />
                    Supplier
                  </div>
                </SelectItem>
              </SelectContent>
            </Select>
          )}
        </div>

        {userType && renderDashboard()}
      </div>
    </div>
  );
}
