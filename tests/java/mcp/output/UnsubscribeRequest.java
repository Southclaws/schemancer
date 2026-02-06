package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Sent from the client to request cancellation of resources/updated notifications from the server. This should follow a previous resources/subscribe request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class UnsubscribeRequest {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public UnsubscribeRequestParams params;

    public UnsubscribeRequest() {
        this.params = new UnsubscribeRequestParams();
    }
}
