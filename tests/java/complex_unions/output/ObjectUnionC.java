package com.example.generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonTypeName;

@JsonTypeName("c")
public record ObjectUnionC(
    @JsonProperty(value = "kind") String kind,
    @JsonProperty(value = "cField", required = true) boolean cfield
) implements ObjectUnion {
    @JsonCreator
    public ObjectUnionC {}
}
