package com.example.generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonTypeName;

@JsonTypeName("b")
public record ObjectUnionB(
    @JsonProperty(value = "kind") String kind,
    @JsonProperty(value = "bField", required = true) int bfield
) implements ObjectUnion {
    @JsonCreator
    public ObjectUnionB {}
}
