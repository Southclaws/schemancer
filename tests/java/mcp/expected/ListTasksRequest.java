package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** A request to retrieve a list of tasks. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ListTasksRequest {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params")
    public PaginatedRequestParams params;

    public ListTasksRequest() {
        this.params = new PaginatedRequestParams();
    }
}
