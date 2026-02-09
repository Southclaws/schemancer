from __future__ import annotations

from typing import Annotated, List, Literal, Union
from pydantic import BaseModel, ConfigDict, Field




class BaseField(BaseModel):
    model_config = ConfigDict(extra="forbid")

    name: str
    required: bool
    type: str


class TextField(BaseField):
    type: Literal["text"]


class ObjectField(BaseField):
    type: Literal["object"]
    fields: List[FieldSchema] | None = None


class ArrayField(BaseField):
    type: Literal["array"]
    fields: List[FieldSchema] | None = None

FieldSchema = Annotated[
    Union[TextField, ObjectField, ArrayField],
    Field(discriminator="type"),
]



class Root(BaseModel):
    model_config = ConfigDict(extra="forbid")

    fields: List[FieldSchema]

