package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;

/** Schema for array items with enum options and display labels. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class TitledMultiSelectEnumSchemaItems {
    /** Array of enum options with values and display labels. */
    @JsonProperty(value = "anyOf", required = true)
    public List<TitledMultiSelectEnumSchemaItemsAnyOfItem> anyOf = new ArrayList<>();
}
