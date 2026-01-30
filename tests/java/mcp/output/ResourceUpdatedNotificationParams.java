package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.Map;


/** Parameters for a `notifications/resources/updated` notification. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ResourceUpdatedNotificationParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /** The URI of the resource that has been updated. This might be a sub-resource of the one that the client actually subscribed to. */
    @JsonProperty(value = "uri", required = true)
    public URI uri;
}
