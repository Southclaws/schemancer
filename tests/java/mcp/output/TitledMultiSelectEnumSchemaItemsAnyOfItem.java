package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


@JsonIgnoreProperties(ignoreUnknown = true)
public class TitledMultiSelectEnumSchemaItemsAnyOfItem {
    /** The constant enum value. */
    @JsonProperty(value = "const", required = true)
    public String const;
    /** Display title for this option. */
    @JsonProperty(value = "title", required = true)
    public String title;
}
