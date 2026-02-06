package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class BooleanSchema {
    @JsonProperty(value = "default")
    public Boolean default;
    @JsonProperty(value = "description")
    public String description;
    @JsonProperty(value = "title")
    public String title;
    @JsonProperty(value = "type", required = true)
    public String type;
}
