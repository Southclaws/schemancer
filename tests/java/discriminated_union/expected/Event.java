package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;

@JsonIgnoreProperties(ignoreUnknown = true)
@JsonTypeInfo(
    use = JsonTypeInfo.Id.NAME,
    include = JsonTypeInfo.As.PROPERTY,
    property = "type",
    visible = true
)
@JsonSubTypes({
    @JsonSubTypes.Type(value = CreatedEvent.class, name = "created"),
    @JsonSubTypes.Type(value = UpdatedEvent.class, name = "updated"),
    @JsonSubTypes.Type(value = DeletedEvent.class, name = "deleted")
})
public sealed interface Event permits CreatedEvent, UpdatedEvent, DeletedEvent {
    String type();
}
