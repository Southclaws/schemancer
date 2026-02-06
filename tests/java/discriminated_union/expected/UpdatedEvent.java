package com.example.generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonTypeName;
import java.time.OffsetDateTime;
import java.util.Map;

@JsonTypeName("updated")
public record UpdatedEvent(
    @JsonProperty(value = "type") String type,
    @JsonProperty(value = "changes", required = true) Map<String, Object> changes,
    @JsonProperty(value = "id", required = true) String id,
    @JsonProperty(value = "timestamp", required = true) OffsetDateTime timestamp
) implements Event {
    @JsonCreator
    public UpdatedEvent {}
}
