import { z } from "zod";

export const BinaryTreeSchema = z.object({
  get left() { return BinaryTreeSchema.optional(); },
  get right() { return BinaryTreeSchema.optional(); },
  value: z.number(),
});
export type BinaryTree = z.infer<typeof BinaryTreeSchema>;

export const CategoryListSchema = z.lazy(() => z.array(CategorySchema));
export type CategoryList = z.infer<typeof CategoryListSchema>;

export const CategorySchema = z.object({
  get children() { return CategoryListSchema.optional(); },
  name: z.string(),
});
export type Category = z.infer<typeof CategorySchema>;

export const EmployeeSchema = z.object({
  get department() { return DepartmentSchema.optional(); },
  name: z.string(),
});
export type Employee = z.infer<typeof EmployeeSchema>;

export const TeamSchema = z.object({
  get members() { return z.array(EmployeeSchema).optional(); },
  name: z.string(),
});
export type Team = z.infer<typeof TeamSchema>;

export const DepartmentSchema = z.object({
  name: z.string(),
  get teams() { return z.array(TeamSchema).optional(); },
});
export type Department = z.infer<typeof DepartmentSchema>;

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
