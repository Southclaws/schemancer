package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class StringSchema {
    @JsonProperty(value = "default")
    public String default;
    @JsonProperty(value = "description")
    public String description;
    @JsonProperty(value = "format")
    public String format;
    @JsonProperty(value = "maxLength")
    public Integer maxLength;
    @JsonProperty(value = "minLength")
    public Integer minLength;
    @JsonProperty(value = "title")
    public String title;
    @JsonProperty(value = "type", required = true)
    public String type;
}
