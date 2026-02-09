import { z } from "zod";

export const BaseEventSchema = z.object({
  timestamp: z.iso.datetime(),
  type: z.string(),
});
export type BaseEvent = z.infer<typeof BaseEventSchema>;

export const CreatedEventSchema = z.object({
  type: z.literal("created"),
  id: z.string(),
  name: z.string(),
  timestamp: z.iso.datetime(),
});
export type CreatedEvent = z.infer<typeof CreatedEventSchema>;

export const UpdatedEventSchema = z.object({
  type: z.literal("updated"),
  changes: z.record(z.string(), z.unknown()),
  id: z.string(),
  timestamp: z.iso.datetime(),
});
export type UpdatedEvent = z.infer<typeof UpdatedEventSchema>;

export const DeletedEventSchema = z.object({
  type: z.literal("deleted"),
  id: z.string(),
  reason: z.string().optional(),
  timestamp: z.iso.datetime(),
});
export type DeletedEvent = z.infer<typeof DeletedEventSchema>;

export const EventSchema = z.discriminatedUnion("type", [
  CreatedEventSchema,
  UpdatedEventSchema,
  DeletedEventSchema,
]);
export type Event = z.infer<typeof EventSchema>;
