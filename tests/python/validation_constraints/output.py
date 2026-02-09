from __future__ import annotations

from typing import List
from pydantic import BaseModel, ConfigDict, EmailStr, Field




class User(BaseModel):
    model_config = ConfigDict(extra="forbid")

    age: int = Field(ge=0, le=150)
    email: EmailStr
    rating: float | None = Field(gt=0, lt=5, default=None)
    score: float | None = Field(ge=0, le=100, multiple_of=0.5, default=None)
    tags: List[str] | None = Field(min_length=1, max_length=10, default=None)
    username: str = Field(min_length=3, max_length=20, pattern=r"^[a-z_][a-z0-9_]*$")

