"use client";

import React, { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Package, TrendingUp, Truck, Factory } from "lucide-react";

interface RoleSelectionFormProps {
  onSubmit: (roles: string[]) => void;
}

export default function RoleSelectionForm({ onSubmit }: RoleSelectionFormProps) {
  const [selectedRoles, setSelectedRoles] = useState<string[]>([]);

  const roles = [
    { id: "customer", icon: Package, label: "Customer" },
    { id: "business", icon: TrendingUp, label: "Business Admin" },
    { id: "transporter", icon: Truck, label: "Transporter" },
    { id: "supplier", icon: Factory, label: "Supplier" },
  ];

  const toggleRole = (roleId: string) => {
    setSelectedRoles(prev =>
      prev.includes(roleId)
        ? prev.filter(r => r !== roleId)
        : [...prev, roleId]
    );
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (selectedRoles.length > 0) {
      onSubmit(selectedRoles);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center">
      <Card className="w-full max-w-2xl">
        <CardHeader>
          <CardTitle>Select Your Roles</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-6">
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
              type="submit" 
              disabled={selectedRoles.length === 0}
              className="w-full"
            >
              Continue
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
