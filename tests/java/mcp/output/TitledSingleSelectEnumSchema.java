package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


/** Schema for single-selection enumeration with display titles for each option. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class TitledSingleSelectEnumSchema {
    /** Optional default value. */
    @JsonProperty(value = "default")
    public String default;
    /** Optional description for the enum field. */
    @JsonProperty(value = "description")
    public String description;
    /** Array of enum options with values and display labels. */
    @JsonProperty(value = "oneOf", required = true)
    public List<TitledSingleSelectEnumSchemaOneOfItem> oneOf;
    /** Optional title for the enum field. */
    @JsonProperty(value = "title")
    public String title;
    @JsonProperty(value = "type", required = true)
    public String type;
}
