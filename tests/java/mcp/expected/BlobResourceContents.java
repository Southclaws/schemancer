package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.HashMap;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown = true)
public class BlobResourceContents {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /** A base64-encoded string representing the binary data of the item. */
    @JsonProperty(value = "blob", required = true)
    public byte[] blob;
    /** The MIME type of this resource, if known. */
    @JsonProperty(value = "mimeType")
    public String mimeType;
    /** The URI of this resource. */
    @JsonProperty(value = "uri", required = true)
    public URI uri;
}
