from __future__ import annotations

from pydantic import BaseModel, ConfigDict




class Address(BaseModel):
    model_config = ConfigDict(extra="forbid")

    city: str
    street: str


class Container(BaseModel):
    model_config = ConfigDict(extra="forbid")

    address: Address | None
    name: str

