package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;

/** Schema for multiple-selection enumeration with display titles for each option. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class TitledMultiSelectEnumSchema {
    /** Optional default value. */
    @JsonProperty(value = "default")
    public List<String> default = new ArrayList<>();
    /** Optional description for the enum field. */
    @JsonProperty(value = "description")
    public String description;
    /** Schema for array items with enum options and display labels. */
    @JsonProperty(value = "items", required = true)
    public TitledMultiSelectEnumSchemaItems items;
    /** Maximum number of items to select. */
    @JsonProperty(value = "maxItems")
    public Integer maxItems;
    /** Minimum number of items to select. */
    @JsonProperty(value = "minItems")
    public Integer minItems;
    /** Optional title for the enum field. */
    @JsonProperty(value = "title")
    public String title;
    @JsonProperty(value = "type", required = true)
    public String type;

    public TitledMultiSelectEnumSchema() {
        this.items = new TitledMultiSelectEnumSchemaItems();
    }
}
