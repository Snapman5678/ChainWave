/* eslint-disable */
"use client";

import React, { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Package, TrendingUp, Truck, Factory } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useAuth } from "../context/AuthContext";

interface RoleSelectionFormProps {
  onSubmit: (roles: string[]) => void;
}

interface LocationData {
  address: string;
  city: string;
  state: string;
  country: string;
  postal_code: string;
  latitude: number;
  longitude: number;
}

interface VehicleData {
  make: string;
  model: string;
  year: number;
  latitude: number;
  longitude: number;
  max_distance: number;
  max_capacity: number;
  current_capacity: number;
}

interface RoleFormData {
  customer?: {
    customer_name: string;
    contact_info: string;
  };
  business?: {
    company_name: string;
    contact_info: string;
  };
  transporter?: {
    driver_name: string;
    contact_info: string;
    vehicle: VehicleData;
  };
  supplier?: {
    supplier_name: string;
    contact_info: string;
    address: string;
  };
  location: LocationData;
}

// Add interceptor setup
axios.interceptors.response.use(
  (response) => {
    // If response contains new token, update it
    if (response.data?.token) {
      const user = JSON.parse(localStorage.getItem("user") || "{}");
      const updatedUser = {
        ...user,
        token: response.data.token
      };
      localStorage.setItem("user", JSON.stringify(updatedUser));
      axios.defaults.headers.common["Authorization"] = `Bearer ${response.data.token}`;
    }
    return response;
  },
  (error) => {
    if (error.response?.status === 404) {
      console.error("Endpoint not found. Check the API URL.");
    } else if (error.response?.status === 409) {
      console.error("Conflict: Role already exists for this user.");
    }
    return Promise.reject(error);
  }
);

export default function RoleSelectionForm({ onSubmit }: RoleSelectionFormProps) {
  const { updateToken } = useAuth();
  const router = useRouter();
  const [selectedRoles, setSelectedRoles] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);
  const [currentStep, setCurrentStep] = useState<'selection' | 'details'>('selection');
  const [formData, setFormData] = useState<Record<string, RoleFormData>>({});
  const [submitted, setSubmitted] = useState(false);

  const roles = [
    { id: "customer", icon: Package, label: "Customer" },
    { id: "business", icon: TrendingUp, label: "Business Admin" },
    { id: "transporter", icon: Truck, label: "Transporter" },
    { id: "supplier", icon: Factory, label: "Supplier" },
  ];

  const handleNextStep = () => {
    if (selectedRoles.length > 0) {
      const initialFormData: Record<string, RoleFormData> = {};
      selectedRoles.forEach(role => {
        initialFormData[role] = {
          location: {
            address: "",
            city: "",
            state: "",
            country: "",
            postal_code: "",
            latitude: 0,
            longitude: 0
          }
        };
        // Initialize role-specific data
        switch(role) {
          case 'customer':
            initialFormData[role].customer = { customer_name: '', contact_info: '' };
            break;
          case 'business':
            initialFormData[role].business = { company_name: '', contact_info: '' };
            break;
          case 'transporter':
            initialFormData[role].transporter = {
              driver_name: '',
              contact_info: '',
              vehicle: {
                make: '',
                model: '',
                year: new Date().getFullYear(),
                latitude: 0,
                longitude: 0,
                max_distance: 500,
                max_capacity: 1000,
                current_capacity: 0
              }
            };
            break;
          case 'supplier':
            initialFormData[role].supplier = { supplier_name: '', contact_info: '', address: '' };
            break;
        }
      });
      setFormData(initialFormData);
      setCurrentStep('details');
    }
  };

  const toggleRole = (roleId: string) => {
    setSelectedRoles(prev =>
      prev.includes(roleId)
        ? prev.filter(r => r !== roleId)
        : [...prev, roleId]
    );
  };

  const handleFormChange = (role: string, field: string, value: string) => {
    setFormData(prev => ({
      ...prev,
      [role]: {
        ...prev[role],
        [role]: {
          ...prev[role][role as keyof RoleFormData],
          [field]: value
        }
      }
    }));
  };

  const handleLocationChange = (role: string, field: string, value: string) => {
    setFormData(prev => ({
      ...prev,
      [role]: {
        ...prev[role],
        location: {
          ...prev[role].location,
          [field]: value
        }
      }
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (loading || submitted) return;
    setLoading(true);

    try {
      const user = JSON.parse(localStorage.getItem("user") || "{}");
      let hasError = false;

      for (const role of selectedRoles) {
        const roleData = formData[role];
        const endpoint = role === 'business' ? 'business-admin' : role;
        
        try {
          const response = await axios.post(`http://localhost:8000/api/${endpoint}`, roleData);
          if (response.data?.token) {
            updateToken(response.data.token);
          }
        } catch (error: any) {
          if (error.response?.status === 409) {
            console.warn(`Role ${role} already exists for this user`);
            continue;
          }
          hasError = true;
          throw error;
        }
      }

      if (!hasError) {
        // Get updated roles
        const rolesResponse = await axios.get('http://localhost:8000/api/role');
        if (rolesResponse.data?.token) {
          updateToken(rolesResponse.data.token);
        }

        // Update local storage with new roles
        const updatedUser = {
          ...user,
          roles: rolesResponse.data.roles || [],
        };
        localStorage.setItem("user", JSON.stringify(updatedUser));

        // Call onSubmit with the selected roles and redirect
        onSubmit(selectedRoles);
      }

    } catch (error: any) {
      console.error("Error submitting roles:", error.message);
      setSubmitted(false);
    } finally {
      setLoading(false);
    }
  };

  const renderRoleDetails = (role: string) => {
    const commonLocationFields = (
      <div className="space-y-4 mt-4">
        <h3 className="font-medium">Location Details</h3>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <Label>Address</Label>
            <Input
              placeholder="Address"
              value={formData[role]?.location?.address || ''}
              onChange={(e) => handleLocationChange(role, 'address', e.target.value)}
            />
          </div>
          <div>
            <Label>City</Label>
            <Input
              placeholder="City"
              value={formData[role]?.location?.city || ''}
              onChange={(e) => handleLocationChange(role, 'city', e.target.value)}
            />
          </div>
          <div>
            <Label>State</Label>
            <Input
              placeholder="State"
              value={formData[role]?.location?.state || ''}
              onChange={(e) => handleLocationChange(role, 'state', e.target.value)}
            />
          </div>
          <div>
            <Label>Country</Label>
            <Input
              placeholder="Country"
              value={formData[role]?.location?.country || ''}
              onChange={(e) => handleLocationChange(role, 'country', e.target.value)}
            />
          </div>
          <div>
            <Label>Postal Code</Label>
            <Input
              placeholder="Postal Code"
              value={formData[role]?.location?.postal_code || ''}
              onChange={(e) => handleLocationChange(role, 'postal_code', e.target.value)}
            />
          </div>
        </div>
      </div>
    );

    switch (role) {
      case 'customer':
        return (
          <div key={role} className="border p-4 rounded-lg mt-4">
            <h2 className="text-lg font-semibold mb-4">Customer Details</h2>
            <div className="space-y-4">
              <div>
                <Label>Customer Name</Label>
                <Input
                  placeholder="Customer Name"
                  value={formData[role]?.customer?.customer_name || ''}
                  onChange={(e) => handleFormChange(role, 'customer_name', e.target.value)}
                />
              </div>
              <div>
                <Label>Contact Info</Label>
                <Input
                  placeholder="Contact Info"
                  value={formData[role]?.customer?.contact_info || ''}
                  onChange={(e) => handleFormChange(role, 'contact_info', e.target.value)}
                />
              </div>
            </div>
            {commonLocationFields}
          </div>
        );

      case 'business':
        return (
          <div key={role} className="border p-4 rounded-lg mt-4">
            <h2 className="text-lg font-semibold mb-4">Business Details</h2>
            <div className="space-y-4">
              <div>
                <Label>Company Name</Label>
                <Input
                  placeholder="Company Name"
                  value={formData[role]?.business?.company_name || ''}
                  onChange={(e) => handleFormChange(role, 'company_name', e.target.value)}
                />
              </div>
              <div>
                <Label>Contact Info</Label>
                <Input
                  placeholder="Contact Info"
                  value={formData[role]?.business?.contact_info || ''}
                  onChange={(e) => handleFormChange(role, 'contact_info', e.target.value)}
                />
              </div>
            </div>
            {commonLocationFields}
          </div>
        );

      case 'transporter':
        return (
          <div key={role} className="border p-4 rounded-lg mt-4">
            <h2 className="text-lg font-semibold mb-4">Transporter Details</h2>
            <div className="space-y-4">
              <div>
                <Label>Driver Name</Label>
                <Input
                  placeholder="Driver Name"
                  value={formData[role]?.transporter?.driver_name || ''}
                  onChange={(e) => handleFormChange(role, 'driver_name', e.target.value)}
                />
              </div>
              <div>
                <Label>Contact Info</Label>
                <Input
                  placeholder="Contact Info"
                  value={formData[role]?.transporter?.contact_info || ''}
                  onChange={(e) => handleFormChange(role, 'contact_info', e.target.value)}
                />
              </div>
              <h3 className="font-medium">Vehicle Details</h3>
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <Label>Make</Label>
                  <Input
                    placeholder="Vehicle Make"
                    value={formData[role]?.transporter?.vehicle?.make || ''}
                    onChange={(e) => handleFormChange(role, 'vehicle_make', e.target.value)}
                  />
                </div>
                <div>
                  <Label>Model</Label>
                  <Input
                    placeholder="Vehicle Model"
                    value={formData[role]?.transporter?.vehicle?.model || ''}
                    onChange={(e) => handleFormChange(role, 'vehicle_model', e.target.value)}
                  />
                </div>
                <div>
                  <Label>Max Capacity (kg)</Label>
                  <Input
                    type="number"
                    placeholder="Max Capacity"
                    value={formData[role]?.transporter?.vehicle?.max_capacity || ''}
                    onChange={(e) => handleFormChange(role, 'max_capacity', e.target.value)}
                  />
                </div>
                <div>
                  <Label>Max Distance (km)</Label>
                  <Input
                    type="number"
                    placeholder="Max Distance"
                    value={formData[role]?.transporter?.vehicle?.max_distance || ''}
                    onChange={(e) => handleFormChange(role, 'max_distance', e.target.value)}
                  />
                </div>
              </div>
            </div>
            {commonLocationFields}
          </div>
        );

      case 'supplier':
        return (
          <div key={role} className="border p-4 rounded-lg mt-4">
            <h2 className="text-lg font-semibold mb-4">Supplier Details</h2>
            <div className="space-y-4">
              <div>
                <Label>Supplier Name</Label>
                <Input
                  placeholder="Supplier Name"
                  value={formData[role]?.supplier?.supplier_name || ''}
                  onChange={(e) => handleFormChange(role, 'supplier_name', e.target.value)}
                />
              </div>
              <div>
                <Label>Contact Info</Label>
                <Input
                  placeholder="Contact Info"
                  value={formData[role]?.supplier?.contact_info || ''}
                  onChange={(e) => handleFormChange(role, 'contact_info', e.target.value)}
                />
              </div>
              <div>
                <Label>Address</Label>
                <Input
                  placeholder="Address"
                  value={formData[role]?.supplier?.address || ''}
                  onChange={(e) => handleFormChange(role, 'address', e.target.value)}
                />
              </div>
            </div>
            {commonLocationFields}
          </div>
        );
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center">
      <Card className="w-full max-w-2xl">
        <CardHeader>
          <CardTitle>
            {currentStep === 'selection' ? 'Select Your Roles' : 'Enter Role Details'}
          </CardTitle>
        </CardHeader>
        <CardContent>
          {currentStep === 'selection' ? (
            <div className="space-y-6">
              <div className="grid grid-cols-2 gap-4">
                {roles.map(({ id, icon: Icon, label }) => (
                  <button
                    key={id}
                    type="button"
                    onClick={() => toggleRole(id)}
                    className={`p-4 border rounded-lg flex items-center space-x-3 hover:bg-gray-50 transition-colors ${
                      selectedRoles.includes(id) ? 'border-indigo-500 bg-indigo-50' : 'border-gray-200'
                    }`}
                  >
                    <Icon className="h-5 w-5" />
                    <span>{label}</span>
                  </button>
                ))}
              </div>
              <Button 
                onClick={handleNextStep}
                disabled={selectedRoles.length === 0}
                className="w-full"
              >
                Next
              </Button>
            </div>
          ) : (
            <form onSubmit={handleSubmit} className="space-y-6">
              {selectedRoles.map(role => renderRoleDetails(role))}
              <Button 
                type="submit" 
                disabled={loading || submitted}
                className="w-full"
              >
                {loading ? "Processing..." : submitted ? "Submitted" : "Submit"}
              </Button>
            </form>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
