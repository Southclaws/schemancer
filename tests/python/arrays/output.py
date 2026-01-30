from __future__ import annotations

from typing import List
from pydantic import BaseModel, ConfigDict, RootModel
class IntArray(RootModel[List[int]]):
    pass

class NestedArray(RootModel[List[List[float]]]):
    pass


class ObjectArrayItem(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: str
    value: int | None = None
class ObjectArray(RootModel[List[ObjectArrayItem]]):
    pass

class StringArray(RootModel[List[str]]):
    pass


class MixedContainer(BaseModel):
    model_config = ConfigDict(extra="forbid")

    nested: NestedArray | None = None
    numbers: IntArray | None = None
    objects: ObjectArray | None = None
    strings: StringArray | None = None
