package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** An optional notification from the server to the client, informing it that the list of resources it can read from has changed. This may be issued by servers without any previous subscription from the client. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ResourceListChangedNotification {
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params")
    public NotificationParams params;

    public ResourceListChangedNotification() {
        this.params = new NotificationParams();
    }
}
