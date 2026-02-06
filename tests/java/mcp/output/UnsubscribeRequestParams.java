package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;

/** Parameters for a `resources/unsubscribe` request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class UnsubscribeRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public UnsubscribeRequestParamsMeta meta;
    /** The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it. */
    @JsonProperty(value = "uri", required = true)
    public URI uri;

    public UnsubscribeRequestParams() {
        this.meta = new UnsubscribeRequestParamsMeta();
    }
}
