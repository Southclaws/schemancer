package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** An optional notification from the server to the client, informing it of a completion of a out-of-band elicitation request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ElicitationCompleteNotification {
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public ElicitationCompleteNotificationParams params;

    public ElicitationCompleteNotification() {
        this.params = new ElicitationCompleteNotificationParams();
    }
}
