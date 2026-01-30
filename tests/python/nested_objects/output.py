from __future__ import annotations

from typing import List
from pydantic import BaseModel, ConfigDict

class Customer(BaseModel):
    model_config = ConfigDict(extra="forbid")

    email: str | None = None
    id: str
    name: str

class LineItem(BaseModel):
    model_config = ConfigDict(extra="forbid")

    price: float
    product_id: str
    quantity: int

class Order(BaseModel):
    model_config = ConfigDict(extra="forbid")

    customer: Customer
    id: str
    items: List[LineItem]
    total: float | None = None
