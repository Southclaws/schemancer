package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** Specifies which request types can be augmented with tasks. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ServerCapabilitiesTasksRequests {
    /** Task support for tool-related requests. */
    @JsonProperty(value = "tools")
    public ServerCapabilitiesTasksRequestsTools tools;
}
