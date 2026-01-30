package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


/**
 * Use TitledSingleSelectEnumSchema instead.
 * This interface will be removed in a future version.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class LegacyTitledEnumSchema {
    @JsonProperty(value = "default")
    public String default;
    @JsonProperty(value = "description")
    public String description;
    @JsonProperty(value = "enum", required = true)
    public List<String> enum;
    /**
 * (Legacy) Display names for enum values.
 * Non-standard according to JSON schema 2020-12.
 */
    @JsonProperty(value = "enumNames")
    public List<String> enumNames;
    @JsonProperty(value = "title")
    public String title;
    @JsonProperty(value = "type", required = true)
    public String type;
}
