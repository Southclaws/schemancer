package com.example.generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonTypeName;

@JsonTypeName("a")
public record ObjectUnionA(
    @JsonProperty(value = "kind") String kind,
    @JsonProperty(value = "aField", required = true) String afield
) implements ObjectUnion {
    @JsonCreator
    public ObjectUnionA {}
}
