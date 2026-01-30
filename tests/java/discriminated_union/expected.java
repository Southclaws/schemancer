package com.example.generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import com.fasterxml.jackson.annotation.JsonTypeName;
import java.time.OffsetDateTime;
import java.util.Map;

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
