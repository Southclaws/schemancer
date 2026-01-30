from __future__ import annotations

from pydantic import BaseModel, ConfigDict, RootModel

class AllPrimitives(BaseModel):
    model_config = ConfigDict(extra="forbid")

    bool_field: bool
    int_field: int
    number_field: float
    string_field: str
class Amount(RootModel[float]):
    pass


class MixedRequired(BaseModel):
    model_config = ConfigDict(extra="forbid")

    optional_int: int | None = None
    optional_string: str | None = None
    required_int: int
    required_string: str

class OptionalPrimitives(BaseModel):
    model_config = ConfigDict(extra="forbid")

    maybe_bool: bool | None = None
    maybe_int: int | None = None
    maybe_number: float | None = None
    maybe_string: str | None = None
class Timestamp(RootModel[int]):
    pass

class UserId(RootModel[str]):
    pass


class TypeAliases(BaseModel):
    model_config = ConfigDict(extra="forbid")

    amount: Amount | None = None
    timestamp: Timestamp | None = None
    user_id: UserId | None = None
