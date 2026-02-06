package com.example.generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonTypeName;
import java.time.OffsetDateTime;

@JsonTypeName("deleted")
public record DeletedEvent(
    @JsonProperty(value = "type") String type,
    @JsonProperty(value = "id", required = true) String id,
    @JsonProperty(value = "reason") String reason,
    @JsonProperty(value = "timestamp", required = true) OffsetDateTime timestamp
) implements Event {
    @JsonCreator
    public DeletedEvent {}
}
