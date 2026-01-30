declare const __brand: unique symbol;
type Brand<B> = {[__brand]: B};
export type Branded<T, B> = T & Brand<B>;

// Age in years.
export type Age = Branded<number, "Age">;

// An email address.
export type Email = Branded<string, "Email">;

// A boolean flag.
export type Flag = Branded<boolean, "Flag">;

// A monetary amount.
export type Money = Branded<number, "Money">;

// A person's name.
export type Name = Branded<string, "Name">;

// Unique identifier for a user.
export type UserID = Branded<string, "UserID">;

export interface User {
  age?: Age;
  balance?: Money;
  email: Email;
  id: UserID;
  isActive?: Flag;
  name: Name;
}
