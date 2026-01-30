package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


@JsonIgnoreProperties(ignoreUnknown = true)
public class NumberSchema {
    @JsonProperty(value = "default")
    public Integer default;
    @JsonProperty(value = "description")
    public String description;
    @JsonProperty(value = "maximum")
    public Integer maximum;
    @JsonProperty(value = "minimum")
    public Integer minimum;
    @JsonProperty(value = "title")
    public String title;
    @JsonProperty(value = "type", required = true)
    public String type;
}
