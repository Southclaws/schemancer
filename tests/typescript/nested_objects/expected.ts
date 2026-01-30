export interface Customer {
  email?: string;
  id: string;
  name: string;
}

export interface LineItem {
  price: number;
  productId: string;
  quantity: number;
}

export interface Order {
  customer: Customer;
  id: string;
  items: LineItem[];
  total?: number;
}
