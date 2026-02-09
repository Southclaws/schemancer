from __future__ import annotations

from uuid import UUID
from pydantic import BaseModel, ConfigDict, EmailStr




class User(BaseModel):
    model_config = ConfigDict(extra="forbid")

    active: bool | None = None
    age: int | None = None
    email: EmailStr
    id: UUID
    name: str

