import { z } from "zod";

export const BinaryTreeSchema = z.object({
  get left(): z.ZodOptional<typeof BinaryTreeSchema> { return BinaryTreeSchema.optional(); },
  get right(): z.ZodOptional<typeof BinaryTreeSchema> { return BinaryTreeSchema.optional(); },
  value: z.number(),
});
export type BinaryTree = z.infer<typeof BinaryTreeSchema>;

export const CategoryListSchema = z.lazy(() => z.array(CategorySchema));
export type CategoryList = z.infer<typeof CategoryListSchema>;

export const CategorySchema = z.object({
  get children(): z.ZodOptional<typeof CategoryListSchema> { return CategoryListSchema.optional(); },
  name: z.string(),
});
export type Category = z.infer<typeof CategorySchema>;

export const EmployeeSchema = z.object({
  get department(): z.ZodOptional<typeof DepartmentSchema> { return DepartmentSchema.optional(); },
  name: z.string(),
});
export type Employee = z.infer<typeof EmployeeSchema>;

export const TeamSchema = z.object({
  get members(): z.ZodOptional<z.ZodArray<typeof EmployeeSchema>> { return z.array(EmployeeSchema).optional(); },
  name: z.string(),
});
export type Team = z.infer<typeof TeamSchema>;

export const DepartmentSchema = z.object({
  name: z.string(),
  get teams(): z.ZodOptional<z.ZodArray<typeof TeamSchema>> { return z.array(TeamSchema).optional(); },
});
export type Department = z.infer<typeof DepartmentSchema>;

export const GraphEdgesItemSchema = z.object({
  get target(): typeof GraphSchema { return GraphSchema; },
  weight: z.number().optional(),
});
export type GraphEdgesItem = z.infer<typeof GraphEdgesItemSchema>;

export const GraphSchema = z.object({
  get edges(): z.ZodOptional<z.ZodArray<typeof GraphEdgesItemSchema>> { return z.array(GraphEdgesItemSchema).optional(); },
  id: z.string().optional(),
});
export type Graph = z.infer<typeof GraphSchema>;

export const LinkedListNodeSchema = z.object({
  data: z.number().int(),
  get next(): z.ZodOptional<typeof LinkedListNodeSchema> { return LinkedListNodeSchema.optional(); },
});
export type LinkedListNode = z.infer<typeof LinkedListNodeSchema>;

export const MutualBSchema = z.object({
  get a(): z.ZodOptional<typeof MutualASchema> { return MutualASchema.optional(); },
  name: z.string(),
});
export type MutualB = z.infer<typeof MutualBSchema>;

export const MutualASchema = z.object({
  get b(): z.ZodOptional<typeof MutualBSchema> { return MutualBSchema.optional(); },
  name: z.string(),
});
export type MutualA = z.infer<typeof MutualASchema>;

export const TreeNodeSchema = z.object({
  get children(): z.ZodOptional<z.ZodArray<typeof TreeNodeSchema>> { return z.array(TreeNodeSchema).optional(); },
  value: z.string(),
});
export type TreeNode = z.infer<typeof TreeNodeSchema>;
