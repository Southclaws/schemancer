from __future__ import annotations

from enum import Enum
from pydantic import BaseModel, ConfigDict




class HttpMethod(str, Enum):
    GET = "GET"
    POST = "POST"
    PUT = "PUT"
    DELETE = "DELETE"
    PATCH = "PATCH"
    HEAD = "HEAD"
    OPTIONS = "OPTIONS"


class ApiRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    body: str | None = None
    method: HttpMethod
    url: str


class Color(str, Enum):
    RED = "red"
    GREEN = "green"
    BLUE = "blue"
    YELLOW = "yellow"


class Priority(str, Enum):
    LOW = "low"
    MEDIUM = "medium"
    HIGH = "high"
    CRITICAL = "critical"


class Status(str, Enum):
    PENDING = "pending"
    IN_PROGRESS = "in_progress"
    COMPLETED = "completed"
    FAILED = "failed"
    CANCELLED = "cancelled"


class Task(BaseModel):
    model_config = ConfigDict(extra="forbid")

    color: Color | None = None
    id: str
    priority: Priority | None = None
    status: Status
    title: str

