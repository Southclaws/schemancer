import { z } from "zod";

export const HttpMethodSchema = z.enum(["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"]);
export type HttpMethod = z.infer<typeof HttpMethodSchema>;

export const ApiRequestSchema = z.object({
  body: z.string().optional(),
  method: HttpMethodSchema,
  url: z.string(),
});
export type ApiRequest = z.infer<typeof ApiRequestSchema>;

export const ColorSchema = z.enum(["red", "green", "blue", "yellow"]);
export type Color = z.infer<typeof ColorSchema>;

export const PrioritySchema = z.enum(["low", "medium", "high", "critical"]);
export type Priority = z.infer<typeof PrioritySchema>;

export const StatusSchema = z.enum(["pending", "in_progress", "completed", "failed", "cancelled"]);
export type Status = z.infer<typeof StatusSchema>;

export const TaskSchema = z.object({
  color: ColorSchema.optional(),
  id: z.string(),
  priority: PrioritySchema.optional(),
  status: StatusSchema,
  title: z.string(),
});
export type Task = z.infer<typeof TaskSchema>;
