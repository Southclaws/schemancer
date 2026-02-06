from __future__ import annotations

from pydantic import BaseModel, ConfigDict, Field

class Task(BaseModel):
    model_config = ConfigDict(extra="forbid")

    class_: str = Field(alias="class")
    from_: str | None = Field(alias="from", default=None)
    import_: str | None = Field(alias="import", default=None)
    name: str
    type: str
