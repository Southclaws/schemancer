import { z } from "zod";

export const BaseFieldSchema = z.object({
  name: z.string(),
  required: z.boolean(),
  type: z.string(),
});
export type BaseField = z.infer<typeof BaseFieldSchema>;

export const TextFieldSchema = z.object({
  type: z.literal("text"),
  name: z.string(),
  required: z.boolean(),
});
export type TextField = z.infer<typeof TextFieldSchema>;

export const ObjectFieldSchema = z.object({
  type: z.literal("object"),
  get fields(): z.ZodOptional<z.ZodArray<z.ZodType>> { return z.array(FieldSchemaSchema).optional(); },
  name: z.string(),
  required: z.boolean(),
});
export type ObjectField = z.infer<typeof ObjectFieldSchema>;

export const ArrayFieldSchema = z.object({
  type: z.literal("array"),
  get fields(): z.ZodOptional<z.ZodArray<z.ZodType>> { return z.array(FieldSchemaSchema).optional(); },
  name: z.string(),
  required: z.boolean(),
});
export type ArrayField = z.infer<typeof ArrayFieldSchema>;

export const FieldSchemaSchema = z.union([
  TextFieldSchema,
  ObjectFieldSchema,
  ArrayFieldSchema,
]);
export type FieldSchema = z.infer<typeof FieldSchemaSchema>;

export const RootSchema = z.object({
  fields: z.array(FieldSchemaSchema),
});
export type Root = z.infer<typeof RootSchema>;
