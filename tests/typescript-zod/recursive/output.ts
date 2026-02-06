import { z } from "zod";

export const BinaryTreeSchema = z.object({
  get left() { return BinaryTreeSchema.optional(); },
  get right() { return BinaryTreeSchema.optional(); },
  value: z.number(),
});
export type BinaryTree = z.infer<typeof BinaryTreeSchema>;

export const GraphEdgesItemSchema = z.object({
  get target() { return GraphSchema; },
  weight: z.number().optional(),
});
export type GraphEdgesItem = z.infer<typeof GraphEdgesItemSchema>;

export const GraphSchema = z.object({
  get edges() { return z.array(GraphEdgesItemSchema).optional(); },
  id: z.string().optional(),
});
export type Graph = z.infer<typeof GraphSchema>;

export const LinkedListNodeSchema = z.object({
  data: z.number().int(),
  get next() { return LinkedListNodeSchema.optional(); },
});
export type LinkedListNode = z.infer<typeof LinkedListNodeSchema>;

export const MutualBSchema = z.object({
  get a() { return MutualASchema.optional(); },
  name: z.string(),
});
export type MutualB = z.infer<typeof MutualBSchema>;

export const MutualASchema = z.object({
  get b() { return MutualBSchema.optional(); },
  name: z.string(),
});
export type MutualA = z.infer<typeof MutualASchema>;

export const TreeNodeSchema = z.object({
  get children() { return z.array(TreeNodeSchema).optional(); },
  value: z.string(),
});
export type TreeNode = z.infer<typeof TreeNodeSchema>;
