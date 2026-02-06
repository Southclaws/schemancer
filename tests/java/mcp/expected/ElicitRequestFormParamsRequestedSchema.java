package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * A restricted subset of JSON Schema.
 * Only top-level properties are allowed, without nesting.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ElicitRequestFormParamsRequestedSchema {
    @JsonProperty(value = "$schema")
    public String schema;
    @JsonProperty(value = "properties", required = true)
    public Map<String, Object> properties = new HashMap<>();
    @JsonProperty(value = "required")
    public List<String> required = new ArrayList<>();
    @JsonProperty(value = "type", required = true)
    public String type;
}
