export interface Address {
  city: string;
  country?: string;
  street: string;
}

export type Status =
  | "active"
  | "inactive"
  | "pending";

export interface Person {
  address?: Address;
  age?: number;
  id: string;
  name: string;
  status: Status;
}
