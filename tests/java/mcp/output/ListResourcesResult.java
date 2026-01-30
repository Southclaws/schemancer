package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;
import java.util.Map;


/** The server's response to a resources/list request from the client. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ListResourcesResult {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /**
 * An opaque token representing the pagination position after the last returned result.
 * If present, there may be more results available.
 */
    @JsonProperty(value = "nextCursor")
    public String nextCursor;
    @JsonProperty(value = "resources", required = true)
    public List<Resource> resources;
}
