package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;

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
