package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** Describes a message issued to or received from an LLM API. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class SamplingMessage {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    @JsonProperty(value = "content", required = true)
    public Object content;
    @JsonProperty(value = "role", required = true)
    public Role role;
}
