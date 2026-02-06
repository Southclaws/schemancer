package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** A ping, issued by either the server or the client, to check that the other party is still alive. The receiver must promptly respond, or else may be disconnected. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class PingRequest {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params")
    public RequestParams params;

    public PingRequest() {
        this.params = new RequestParams();
    }
}
