package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Present if the server supports task-augmented requests. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ServerCapabilitiesTasks {
    /** Whether this server supports tasks/cancel. */
    @JsonProperty(value = "cancel")
    public ServerCapabilitiesTasksCancel cancel;
    /** Whether this server supports tasks/list. */
    @JsonProperty(value = "list")
    public ServerCapabilitiesTasksList list;
    /** Specifies which request types can be augmented with tasks. */
    @JsonProperty(value = "requests")
    public ServerCapabilitiesTasksRequests requests;

    public ServerCapabilitiesTasks() {
        this.cancel = new ServerCapabilitiesTasksCancel();
        this.list = new ServerCapabilitiesTasksList();
        this.requests = new ServerCapabilitiesTasksRequests();
    }
}
