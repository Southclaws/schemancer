from __future__ import annotations

from typing import List
from pydantic import BaseModel, ConfigDict




class BinaryTree(BaseModel):
    model_config = ConfigDict(extra="forbid")

    left: BinaryTree | None = None
    right: BinaryTree | None = None
    value: float


class GraphEdgesItem(BaseModel):
    model_config = ConfigDict(extra="forbid")

    target: Graph
    weight: float | None = None


class Graph(BaseModel):
    model_config = ConfigDict(extra="forbid")

    edges: List[GraphEdgesItem] | None = None
    id: str | None = None


class LinkedListNode(BaseModel):
    model_config = ConfigDict(extra="forbid")

    data: int
    next: LinkedListNode | None = None


class MutualB(BaseModel):
    model_config = ConfigDict(extra="forbid")

    a: MutualA | None = None
    name: str


class MutualA(BaseModel):
    model_config = ConfigDict(extra="forbid")

    b: MutualB | None = None
    name: str


class TreeNode(BaseModel):
    model_config = ConfigDict(extra="forbid")

    children: List[TreeNode] | None = None
    value: str

