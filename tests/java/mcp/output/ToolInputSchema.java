package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/** A JSON Schema object defining the expected parameters for the tool. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ToolInputSchema {
    @JsonProperty(value = "$schema")
    public String schema;
    @JsonProperty(value = "properties")
    public Map<String, ToolInputSchemaPropertiesValue> properties = new HashMap<>();
    @JsonProperty(value = "required")
    public List<String> required = new ArrayList<>();
    @JsonProperty(value = "type", required = true)
    public String type;
}
