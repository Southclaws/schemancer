package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


/** Schema for the array items. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class UntitledMultiSelectEnumSchemaItems {
    /** Array of enum values to choose from. */
    @JsonProperty(value = "enum", required = true)
    public List<String> enum;
    @JsonProperty(value = "type", required = true)
    public String type;
}
