import { z } from "zod";

export const AddressSchema = z.object({
  city: z.string(),
  country: z.string().optional(),
  street: z.string(),
});
export type Address = z.infer<typeof AddressSchema>;

export const StatusSchema = z.enum(["active", "inactive", "pending"]);
export type Status = z.infer<typeof StatusSchema>;

export const PersonSchema = z.object({
  address: AddressSchema.optional(),
  age: z.number().int().optional(),
  id: z.string().uuid(),
  name: z.string(),
  status: StatusSchema,
});
export type Person = z.infer<typeof PersonSchema>;
