package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** An optional notification from the receiver to the requestor, informing them that a task's status has changed. Receivers are not required to send these notifications. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class TaskStatusNotification {
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public TaskStatusNotificationParams params;
}
