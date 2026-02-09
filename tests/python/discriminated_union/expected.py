from __future__ import annotations

from typing import Annotated, Any, Dict, Literal, Union
from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field




class BaseEvent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    timestamp: datetime
    type: str


class CreatedEvent(BaseEvent):
    type: Literal["created"]
    id: str
    name: str


class UpdatedEvent(BaseEvent):
    type: Literal["updated"]
    changes: Dict[str, Any]
    id: str


class DeletedEvent(BaseEvent):
    type: Literal["deleted"]
    id: str
    reason: str | None = None

Event = Annotated[
    Union[CreatedEvent, UpdatedEvent, DeletedEvent],
    Field(discriminator="type"),
]


