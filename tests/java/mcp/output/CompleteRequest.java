package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** A request from the client to the server, to ask for completion options. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CompleteRequest {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public CompleteRequestParams params;

    public CompleteRequest() {
        this.params = new CompleteRequestParams();
    }
}
