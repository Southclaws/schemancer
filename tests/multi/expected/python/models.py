from __future__ import annotations

from uuid import UUID
from enum import Enum
from pydantic import BaseModel, ConfigDict

class Address(BaseModel):
    model_config = ConfigDict(extra="forbid")

    city: str
    country: str | None = None
    street: str
class Status(str, Enum):
    ACTIVE = "active"
    INACTIVE = "inactive"
    PENDING = "pending"


class Person(BaseModel):
    model_config = ConfigDict(extra="forbid")

    address: Address | None = None
    age: int | None = None
    id: UUID
    name: str
    status: Status
