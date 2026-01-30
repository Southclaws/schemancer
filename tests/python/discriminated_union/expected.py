from __future__ import annotations

from typing import Annotated, Any, Dict, Literal, Union
from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field
class CreatedEvent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    type: Literal["created"]
    id: str
    name: str
    timestamp: datetime


class UpdatedEvent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    type: Literal["updated"]
    changes: Dict[str, Any]
    id: str
    timestamp: datetime


class DeletedEvent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    type: Literal["deleted"]
    id: str
    reason: str | None = None
    timestamp: datetime


Event = Annotated[
    Union[CreatedEvent, UpdatedEvent, DeletedEvent],
    Field(discriminator="type"),
]

