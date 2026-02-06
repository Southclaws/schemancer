package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Used by the client to invoke a tool provided by the server. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CallToolRequest {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public CallToolRequestParams params;

    public CallToolRequest() {
        this.params = new CallToolRequestParams();
    }
}
