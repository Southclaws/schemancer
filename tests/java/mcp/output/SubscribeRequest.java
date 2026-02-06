package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Sent from the client to request resources/updated notifications from the server whenever a particular resource changes. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class SubscribeRequest {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public SubscribeRequestParams params;

    public SubscribeRequest() {
        this.params = new SubscribeRequestParams();
    }
}
