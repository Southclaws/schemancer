package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Task support for tool-related requests. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ServerCapabilitiesTasksRequestsTools {
    /** Whether the server supports task-augmented tools/call requests. */
    @JsonProperty(value = "call")
    public ServerCapabilitiesTasksRequestsToolsCall call;

    public ServerCapabilitiesTasksRequestsTools() {
        this.call = new ServerCapabilitiesTasksRequestsToolsCall();
    }
}
