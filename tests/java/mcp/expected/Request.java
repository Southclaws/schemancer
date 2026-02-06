package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.HashMap;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Request {
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params")
    public Map<String, Object> params = new HashMap<>();
}
