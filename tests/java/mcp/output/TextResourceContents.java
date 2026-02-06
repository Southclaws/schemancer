package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.HashMap;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown = true)
public class TextResourceContents {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /** The MIME type of this resource, if known. */
    @JsonProperty(value = "mimeType")
    public String mimeType;
    /** The text of the item. This must only be set if the item can actually be represented as text (not binary data). */
    @JsonProperty(value = "text", required = true)
    public String text;
    /** The URI of this resource. */
    @JsonProperty(value = "uri", required = true)
    public URI uri;
}
