package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/**
 * The contents of a resource, embedded into a prompt or tool call result.
 * 
 * It is up to the client how best to render embedded resources for the benefit
 * of the LLM and/or the user.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class EmbeddedResource {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /** Optional annotations for the client. */
    @JsonProperty(value = "annotations")
    public Annotations annotations;
    @JsonProperty(value = "resource", required = true)
    public Object resource;
    @JsonProperty(value = "type", required = true)
    public String type;
}
