package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


/** Schema for single-selection enumeration without display titles for options. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class UntitledSingleSelectEnumSchema {
    /** Optional default value. */
    @JsonProperty(value = "default")
    public String default;
    /** Optional description for the enum field. */
    @JsonProperty(value = "description")
    public String description;
    /** Array of enum values to choose from. */
    @JsonProperty(value = "enum", required = true)
    public List<String> enum;
    /** Optional title for the enum field. */
    @JsonProperty(value = "title")
    public String title;
    @JsonProperty(value = "type", required = true)
    public String type;
}
