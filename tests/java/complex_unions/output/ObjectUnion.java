package com.example.generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import com.fasterxml.jackson.annotation.JsonTypeName;

/** Test complex anyOf/oneOf scenarios */

@JsonIgnoreProperties(ignoreUnknown = true)
@JsonTypeInfo(
    use = JsonTypeInfo.Id.NAME,
    include = JsonTypeInfo.As.PROPERTY,
    property = "kind",
    visible = true
)
@JsonSubTypes({
    @JsonSubTypes.Type(value = ObjectUnionA.class, name = "a"),
    @JsonSubTypes.Type(value = ObjectUnionB.class, name = "b"),
    @JsonSubTypes.Type(value = ObjectUnionC.class, name = "c")
})
public sealed interface ObjectUnion permits ObjectUnionA, ObjectUnionB, ObjectUnionC {
    String kind();
}


@JsonTypeName("a")
public record ObjectUnionA(
    @JsonProperty(value = "kind") String kind,
    @JsonProperty(value = "aField", required = true) String afield
) implements ObjectUnion {
    @JsonCreator
    public ObjectUnionA {}
}

@JsonTypeName("b")
public record ObjectUnionB(
    @JsonProperty(value = "kind") String kind,
    @JsonProperty(value = "bField", required = true) int bfield
) implements ObjectUnion {
    @JsonCreator
    public ObjectUnionB {}
}

@JsonTypeName("c")
public record ObjectUnionC(
    @JsonProperty(value = "kind") String kind,
    @JsonProperty(value = "cField", required = true) boolean cfield
) implements ObjectUnion {
    @JsonCreator
    public ObjectUnionC {}
}
