package recursive_test

type BinaryTree struct {
	Left  *BinaryTree `json:"left,omitempty"`
	Right *BinaryTree `json:"right,omitempty"`
	Value float64     `json:"value"`
}

type GraphEdgesItem struct {
	Target Graph    `json:"target"`
	Weight *float64 `json:"weight,omitempty"`
}

type Graph struct {
	Edges []GraphEdgesItem `json:"edges,omitempty"`
	ID    *string          `json:"id,omitempty"`
}

type LinkedListNode struct {
	Data int             `json:"data"`
	Next *LinkedListNode `json:"next,omitempty"`
}

type MutualB struct {
	A    *MutualA `json:"a,omitempty"`
	Name string   `json:"name"`
}

type MutualA struct {
	B    *MutualB `json:"b,omitempty"`
	Name string   `json:"name"`
}

type TreeNode struct {
	Children []TreeNode `json:"children,omitempty"`
	Value    string     `json:"value"`
}
