import { z } from "zod";

export const UserSchema = z.object({
  age: z.number().int().min(0).max(150),
  email: z.string().email(),
  rating: z.number().gt(0).lt(5).optional(),
  score: z.number().min(0).max(100).multipleOf(0.5).optional(),
  tags: z.array(z.string()).min(1).max(10).optional(),
  username: z.string().min(3).max(20).regex(/^[a-z_][a-z0-9_]*$/),
});
export type User = z.infer<typeof UserSchema>;
