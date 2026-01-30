import { z } from "zod";

export const HttpStatusSchema = z.union([z.literal(200), z.literal(201), z.literal(400), z.literal(404), z.literal(500)]);
export type HttpStatus = z.infer<typeof HttpStatusSchema>;

export const PrioritySchema = z.union([z.literal(1), z.literal(2), z.literal(3)]);
export type Priority = z.infer<typeof PrioritySchema>;

export const ResponseSchema = z.object({
  priority: PrioritySchema.optional(),
  status: HttpStatusSchema,
});
export type Response = z.infer<typeof ResponseSchema>;
