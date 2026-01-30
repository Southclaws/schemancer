package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** Text provided to or from an LLM. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class TextContent {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /** Optional annotations for the client. */
    @JsonProperty(value = "annotations")
    public Annotations annotations;
    /** The text content of the message. */
    @JsonProperty(value = "text", required = true)
    public String text;
    @JsonProperty(value = "type", required = true)
    public String type;
}
