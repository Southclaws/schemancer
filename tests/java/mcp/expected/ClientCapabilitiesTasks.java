package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Present if the client supports task-augmented requests. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesTasks {
    /** Whether this client supports tasks/cancel. */
    @JsonProperty(value = "cancel")
    public ClientCapabilitiesTasksCancel cancel;
    /** Whether this client supports tasks/list. */
    @JsonProperty(value = "list")
    public ClientCapabilitiesTasksList list;
    /** Specifies which request types can be augmented with tasks. */
    @JsonProperty(value = "requests")
    public ClientCapabilitiesTasksRequests requests;

    public ClientCapabilitiesTasks() {
        this.cancel = new ClientCapabilitiesTasksCancel();
        this.list = new ClientCapabilitiesTasksList();
        this.requests = new ClientCapabilitiesTasksRequests();
    }
}
