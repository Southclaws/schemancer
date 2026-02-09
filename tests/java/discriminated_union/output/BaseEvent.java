package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.time.OffsetDateTime;

@JsonIgnoreProperties(ignoreUnknown = true)
public class BaseEvent {
    @JsonProperty(value = "timestamp", required = true)
    public OffsetDateTime timestamp;
    @JsonProperty(value = "type", required = true)
    public String type;
}
