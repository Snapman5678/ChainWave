export interface Product {
  id: string;
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

export interface User {
  email: string;
  username: string;
  token: string;
}
