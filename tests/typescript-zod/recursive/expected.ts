import { z } from "zod";

export const BinaryTreeSchema = z.object({
  get left(): z.ZodOptional<typeof BinaryTreeSchema> { return BinaryTreeSchema.optional(); },
  get right(): z.ZodOptional<typeof BinaryTreeSchema> { return BinaryTreeSchema.optional(); },
  value: z.number(),
});
export interface BinaryTree {
  left?: BinaryTree;
  right?: BinaryTree;
  value: number;
}

export const CategoryListSchema = z.lazy(() => z.array(CategorySchema));
export type CategoryList = Category[];

export const CategorySchema = z.object({
  get children(): z.ZodOptional<typeof CategoryListSchema> { return CategoryListSchema.optional(); },
  name: z.string(),
});
export interface Category {
  children?: CategoryList;
  name: string;
}

export const EmployeeSchema = z.object({
  get department(): z.ZodOptional<typeof DepartmentSchema> { return DepartmentSchema.optional(); },
  name: z.string(),
});
export interface Employee {
  department?: Department;
  name: string;
}

export const TeamSchema = z.object({
  get members(): z.ZodOptional<z.ZodArray<typeof EmployeeSchema>> { return z.array(EmployeeSchema).optional(); },
  name: z.string(),
});
export interface Team {
  members?: Employee[];
  name: string;
}

export const DepartmentSchema = z.object({
  name: z.string(),
  get teams(): z.ZodOptional<z.ZodArray<typeof TeamSchema>> { return z.array(TeamSchema).optional(); },
});
export interface Department {
  name: string;
  teams?: Team[];
}

export const GraphEdgesItemSchema = z.object({
  get target(): typeof GraphSchema { return GraphSchema; },
  weight: z.number().optional(),
});
export interface GraphEdgesItem {
  target: Graph;
  weight?: number;
}

export const GraphSchema = z.object({
  get edges(): z.ZodOptional<z.ZodArray<typeof GraphEdgesItemSchema>> { return z.array(GraphEdgesItemSchema).optional(); },
  id: z.string().optional(),
});
export interface Graph {
  edges?: GraphEdgesItem[];
  id?: string;
}

export const LinkedListNodeSchema = z.object({
  data: z.number().int(),
  get next(): z.ZodOptional<typeof LinkedListNodeSchema> { return LinkedListNodeSchema.optional(); },
});
export interface LinkedListNode {
  data: number;
  next?: LinkedListNode;
}

export const MutualBSchema = z.object({
  get a(): z.ZodOptional<typeof MutualASchema> { return MutualASchema.optional(); },
  name: z.string(),
});
export interface MutualB {
  a?: MutualA;
  name: string;
}

export const MutualASchema = z.object({
  get b(): z.ZodOptional<typeof MutualBSchema> { return MutualBSchema.optional(); },
  name: z.string(),
});
export interface MutualA {
  b?: MutualB;
  name: string;
}

export const TreeNodeSchema = z.object({
  get children(): z.ZodOptional<z.ZodArray<typeof TreeNodeSchema>> { return z.array(TreeNodeSchema).optional(); },
  value: z.string(),
});
export interface TreeNode {
  children?: TreeNode[];
  value: string;
}
