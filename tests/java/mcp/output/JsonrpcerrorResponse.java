package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** A response to a request that indicates an error occurred. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class JSONRPCErrorResponse {
    @JsonProperty(value = "error", required = true)
    public Error error;
    @JsonProperty(value = "id")
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;

    public JSONRPCErrorResponse() {
        this.error = new Error();
    }
}
