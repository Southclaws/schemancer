package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.HashMap;
import java.util.Map;

/** The contents of a specific resource or sub-resource. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ResourceContents {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /** The MIME type of this resource, if known. */
    @JsonProperty(value = "mimeType")
    public String mimeType;
    /** The URI of this resource. */
    @JsonProperty(value = "uri", required = true)
    public URI uri;
}
