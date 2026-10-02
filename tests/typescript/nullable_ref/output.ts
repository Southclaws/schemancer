export interface Address {
  city: string;
  street: string;
}

export interface Container {
  address: Address | null;
  name: string;
}
