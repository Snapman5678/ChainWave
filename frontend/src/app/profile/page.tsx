/* eslint-disable */
"use client";

import { useState } from "react";
import { useAuth } from "../context/AuthContext";
import axios from "axios";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { toast } from "sonner";

export default function ProfilePage() {
  const { user, setUser } = useAuth();
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState({
    email: false,
    username: false,
    password: false,
  });

  const handleUpdateEmail = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!email) return;
    setIsLoading(prev => ({ ...prev, email: true }));
    try {
      const response = await axios.put(
        "http://localhost:8000/api/user/email",
        { email },
        {
          headers: { Authorization: `Bearer ${user?.token}` },
        }
      );
      
      // Update user state while preserving token and other properties
      if (user) {
        setUser({
          ...user,
          email: email
        });
      }
      
      toast.success("Email updated successfully!");
      setEmail("");
    } catch (error: any) {
      toast.error(error.response?.data?.error || "Failed to update email");
    } finally {
      setIsLoading(prev => ({ ...prev, email: false }));
    }
  };

  const handleUpdateUsername = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!username) return;
    setIsLoading(prev => ({ ...prev, username: true }));
    try {
      const response = await axios.put(
        "http://localhost:8000/api/user/username",
        { username },
        {
          headers: { Authorization: `Bearer ${user?.token}` },
        }
      );
      
      // Update user state while preserving token and other properties
      if (user) {
        setUser({
          ...user,
          username: username
        });
      }
      
      toast.success("Username updated successfully!");
      setUsername("");
    } catch (error: any) {
      toast.error(error.response?.data?.error || "Failed to update username");
    } finally {
      setIsLoading(prev => ({ ...prev, username: false }));
    }
  };

  const handleUpdatePassword = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!password) return;
    setIsLoading(prev => ({ ...prev, password: true }));
    try {
      const response = await axios.put(
        "http://localhost:8000/api/user/password",
        { password },
        {
          headers: { Authorization: `Bearer ${user?.token}` },
        }
      );
      
      // No need to update user state for password change
      toast.success("Password updated successfully!");
      setPassword("");
    } catch (error: any) {
      toast.error(error.response?.data?.error || "Failed to update password");
    } finally {
      setIsLoading(prev => ({ ...prev, password: false }));
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="max-w-2xl mx-auto space-y-6 px-4">
        <h1 className="text-2xl font-bold mb-6">Profile Settings</h1>

        <Card>
          <CardHeader>
            <CardTitle>Update Email</CardTitle>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleUpdateEmail} className="space-y-4">
              <Input
                type="email"
                placeholder="New email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
              <Button type="submit" disabled={isLoading.email}>
                {isLoading.email ? "Updating..." : "Update Email"}
              </Button>
            </form>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Update Username</CardTitle>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleUpdateUsername} className="space-y-4">
              <Input
                type="text"
                placeholder="New username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
              />
              <Button type="submit" disabled={isLoading.username}>
                {isLoading.username ? "Updating..." : "Update Username"}
              </Button>
            </form>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Update Password</CardTitle>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleUpdatePassword} className="space-y-4">
              <Input
                type="password"
                placeholder="New password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
              />
              <Button type="submit" disabled={isLoading.password}>
                {isLoading.password ? "Updating..." : "Update Password"}
              </Button>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
