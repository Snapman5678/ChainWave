"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import axios from "axios";
import { Eye, EyeOff } from "lucide-react";
import { useAuth } from "../../context/AuthContext";

interface AuthFormProps {
  mode: "login" | "register";
}

interface ValidationErrors {
  email: string;
  password: string[];
}

interface PasswordCriterion {
  regex: RegExp;
  message: string;
  met: boolean;
}

const setAuthToken = (token: string, userId?: string, username?: string) => {
  if (token) {
    localStorage.setItem("authToken", token);
    if (userId) {
      localStorage.setItem("userId", userId);
    }
    if (username) {
      localStorage.setItem("username", username);
    }
    axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  } else {
    localStorage.removeItem("authToken");
    localStorage.removeItem("userId");
    localStorage.removeItem("username");
    delete axios.defaults.headers.common["Authorization"];
  }
};

// Configure axios interceptor to add token to all requests
axios.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("authToken");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default function AuthForm({ mode }: AuthFormProps) {
  const { setUser } = useAuth();
  const router = useRouter();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
    username: "",
  });
  const [errors, setErrors] = useState<ValidationErrors>({
    email: "",
    password: [],
  });
  const [loading, setLoading] = useState(false);
  const [showPassword, setShowPassword] = useState(false);
  const [passwordCriteria, setPasswordCriteria] = useState<PasswordCriterion[]>(
    [
      { regex: /.{8,}/, message: "At least 8 characters long", met: false },
      {
        regex: /[A-Z]/,
        message: "Contains at least one uppercase letter",
        met: false,
      },
      {
        regex: /[a-z]/,
        message: "Contains at least one lowercase letter",
        met: false,
      },
      { regex: /[0-9]/, message: "Contains at least one number", met: false },
      {
        regex: /[^A-Za-z0-9]/,
        message: "Contains at least one special character",
        met: false,
      },
    ]
  );
  const [loadingMessage, setLoadingMessage] = useState("");

  // Check for existing token on component mount
  useEffect(() => {
    const token = localStorage.getItem("authToken");
    if (token) {
      router.push("/dashboard");
    }
  }, [router]);

  const usernameRegex = /^(?![0-9])[A-Za-z0-9_]+$/;

  const validateUsername = (username: string) => {
    if (!usernameRegex.test(username)) {
      setErrors((prev) => ({
        ...prev,
        username:
          "Username must not start with a number, must not contain spaces, and only underscores are allowed as special characters.",
      }));
      return false;
    }
    setErrors((prev) => ({ ...prev, username: "" }));
    return true;
  };

  const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;

  const validateEmail = (email: string) => {
    if (!emailRegex.test(email)) {
      setErrors((prev) => ({
        ...prev,
        email: "Please enter a valid email address",
      }));
      return false;
    }
    setErrors((prev) => ({ ...prev, email: "" }));
    return true;
  };

  const updatePasswordCriteria = (password: string) => {
    const updatedCriteria = passwordCriteria.map((criterion) => ({
      ...criterion,
      met: criterion.regex.test(password),
    }));
    setPasswordCriteria(updatedCriteria);

    const failedCriteria = updatedCriteria
      .filter((criterion) => !criterion.met)
      .map((criterion) => criterion.message);

    setErrors((prev) => ({ ...prev, password: failedCriteria }));
    return failedCriteria.length === 0;
  };

  const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    let newUsername = e.target.value;
    newUsername = newUsername.replace(/\s/g, "");
    setFormData((prev) => ({ ...prev, username: newUsername }));
    validateUsername(newUsername);
  };

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newEmail = e.target.value;
    setFormData((prev) => ({ ...prev, email: newEmail }));
    validateEmail(newEmail);
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newPassword = e.target.value;
    setFormData((prev) => ({ ...prev, password: newPassword }));
    if (mode === "register") {
      updatePasswordCriteria(newPassword);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setLoadingMessage("");

    const isEmailValid = validateEmail(formData.email);
    const isUsernameValid =
      mode === "register" ? validateUsername(formData.username) : true;
    const isPasswordValid =
      mode === "register" ? updatePasswordCriteria(formData.password) : true;

    if (
      !isEmailValid ||
      !isUsernameValid ||
      (mode === "register" && !isPasswordValid)
    ) {
      setLoading(false);
      return;
    }

    try {
      setLoadingMessage(
        mode === "login" ? "Signing in..." : "Creating account..."
      );

      const endpoint =
        mode === "login" ? "/api/user/login" : "/api/user/register";
      const response = await axios.post(`http://localhost:8000${endpoint}`, {
        username: formData.username,
        email: formData.email,
        password: formData.password,
      });

      console.log("Server response:", response.data); // For debugging

      if (mode === "register") {
        // Registration successful, navigate to login page
        setLoadingMessage("Account created successfully!");
        router.push("/auth/login");
      } else if (response.data.token) {
        // Login successful, set auth token and navigate to dashboard
        const { token, userId, username, email } = response.data;
        setAuthToken(token, userId, username);

        const userData = { username, email };
        setUser(userData);
        localStorage.setItem("user", JSON.stringify(userData));

        setLoadingMessage("Login successful!");
        router.push("/dashboard");
      }
    } catch (err: any) {
      console.error("Error during submission:", err); // For debugging
      const errorMessage =
        err.response?.data?.error || "An error occurred. Please try again.";
      setErrors((prev) => ({
        ...prev,
        email: errorMessage,
      }));
      setAuthToken("");
      setLoadingMessage("");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    setPasswordCriteria((prev) =>
      prev.map((criterion) => ({ ...criterion, met: false }))
    );
  }, [mode]);

  return (
    <div className="space-y-6">
      <div className="text-center">
        <h2 className="text-3xl font-extrabold text-gray-900">
          {mode === "login"
            ? "Sign in to your account"
            : "Create a new account"}
        </h2>
        <p className="mt-2 text-sm text-gray-600">
          {mode === "login" ? (
            <>
              Don&apos;t have an account?{" "}
              <Link
                href="/auth/register"
                className="text-indigo-600 hover:text-indigo-500"
              >
                Register here
              </Link>
            </>
          ) : (
            <>
              Already have an account?{" "}
              <Link
                href="/auth/login"
                className="text-indigo-600 hover:text-indigo-500"
              >
                Sign in
              </Link>
            </>
          )}
        </p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">
        {mode === "register" && (
          <div>
            <label
              htmlFor="username"
              className="block text-sm font-medium text-gray-700"
            >
              Username
            </label>
            <input
              id="username"
              type="text"
              required
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              value={formData.username}
              onChange={handleUsernameChange}
            />
          </div>
        )}

        <div>
          <label
            htmlFor="email"
            className="block text-sm font-medium text-gray-700"
          >
            Email address
          </label>
          <input
            id="email"
            type="email"
            required
            className={`mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
              errors.email ? "border-red-500" : "border-gray-300"
            }`}
            value={formData.email}
            onChange={handleEmailChange}
          />
          {errors.email && (
            <p className="mt-1 text-sm text-red-600">{errors.email}</p>
          )}
        </div>

        <div>
          <label
            htmlFor="password"
            className="block text-sm font-medium text-gray-700"
          >
            Password
          </label>
          <div className="relative">
            <input
              id="password"
              type={showPassword ? "text" : "password"}
              required
              className={`mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                errors.password.length > 0
                  ? "border-red-500"
                  : "border-gray-300"
              }`}
              value={formData.password}
              onChange={handlePasswordChange}
            />
            <button
              type="button"
              className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-500"
              onMouseDown={() => setShowPassword(true)}
              onMouseUp={() => setShowPassword(false)}
              onMouseLeave={() => setShowPassword(false)}
            >
              {showPassword ? (
                <EyeOff className="h-5 w-5" />
              ) : (
                <Eye className="h-5 w-5" />
              )}
            </button>
          </div>

          {mode === "register" && (
            <div className="mt-2 space-y-1">
              {passwordCriteria.map((criterion, index) => (
                <p
                  key={index}
                  className={`text-sm ${
                    formData.password.length === 0
                      ? "text-gray-500"
                      : criterion.met
                      ? "text-green-600"
                      : "text-red-600"
                  }`}
                >
                  • {criterion.message}
                </p>
              ))}
            </div>
          )}
        </div>

        <div>
          <button
            type="submit"
            disabled={loading}
            className={`w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 ${
              loading ? "opacity-50 cursor-not-allowed" : ""
            }`}
          >
            {loadingMessage || (mode === "login" ? "Sign in" : "Register")}
          </button>
          {loading && (
            <div className="mt-4 text-center text-sm text-indigo-600">
              {loadingMessage}
            </div>
          )}
        </div>
      </form>
    </div>
  );
}
