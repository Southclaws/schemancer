package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.Map;


/** Represents a root directory or file that the server can operate on. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class Root {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /**
 * An optional name for the root. This can be used to provide a human-readable
 * identifier for the root, which may be useful for display purposes or for
 * referencing the root in other parts of the application.
 */
    @JsonProperty(value = "name")
    public String name;
    /**
 * The URI identifying the root. This *must* start with file:// for now.
 * This restriction may be relaxed in future versions of the protocol to allow
 * other URI schemes.
 */
    @JsonProperty(value = "uri", required = true)
    public URI uri;
}
