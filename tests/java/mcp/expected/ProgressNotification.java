package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** An out-of-band notification used to inform the receiver of a progress update for a long-running request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ProgressNotification {
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public ProgressNotificationParams params;

    public ProgressNotification() {
        this.params = new ProgressNotificationParams();
    }
}
