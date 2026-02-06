package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Sent from the client to request a list of resources the server has. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ListResourcesRequest {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params")
    public PaginatedRequestParams params;

    public ListResourcesRequest() {
        this.params = new PaginatedRequestParams();
    }
}
