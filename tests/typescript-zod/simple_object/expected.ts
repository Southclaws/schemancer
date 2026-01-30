import { z } from "zod";

export const UserSchema = z.object({
  active: z.boolean().optional(),
  age: z.number().int().optional(),
  email: z.string().email(),
  id: z.string().uuid(),
  name: z.string(),
});
export type User = z.infer<typeof UserSchema>;
