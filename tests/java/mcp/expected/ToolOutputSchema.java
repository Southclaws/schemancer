package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;
import java.util.Map;


/**
 * An optional JSON Schema object defining the structure of the tool's output returned in
 * the structuredContent field of a CallToolResult.
 * 
 * Defaults to JSON Schema 2020-12 when no explicit $schema is provided.
 * Currently restricted to type: "object" at the root level.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ToolOutputSchema {
    @JsonProperty(value = "$schema")
    public String schema;
    @JsonProperty(value = "properties")
    public Map<String, Object> properties;
    @JsonProperty(value = "required")
    public List<String> required;
    @JsonProperty(value = "type", required = true)
    public String type;
}
