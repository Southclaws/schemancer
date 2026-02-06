package com.example.generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonTypeName;
import java.time.OffsetDateTime;

@JsonTypeName("created")
public record CreatedEvent(
    @JsonProperty(value = "type") String type,
    @JsonProperty(value = "id", required = true) String id,
    @JsonProperty(value = "name", required = true) String name,
    @JsonProperty(value = "timestamp", required = true) OffsetDateTime timestamp
) implements Event {
    @JsonCreator
    public CreatedEvent {}
}
