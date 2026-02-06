package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Task {
    @JsonProperty(value = "color")
    public Color color;
    @JsonProperty(value = "id", required = true)
    public String id;
    @JsonProperty(value = "priority")
    public Priority priority;
    @JsonProperty(value = "status", required = true)
    public Status status;
    @JsonProperty(value = "title", required = true)
    public String title;
}
