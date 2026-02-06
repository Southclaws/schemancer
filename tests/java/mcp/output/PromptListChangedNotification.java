package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** An optional notification from the server to the client, informing it that the list of prompts it offers has changed. This may be issued by servers without any previous subscription from the client. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class PromptListChangedNotification {
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params")
    public NotificationParams params;

    public PromptListChangedNotification() {
        this.params = new NotificationParams();
    }
}
