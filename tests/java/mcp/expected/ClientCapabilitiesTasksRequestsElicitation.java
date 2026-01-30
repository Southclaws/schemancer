package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** Task support for elicitation-related requests. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesTasksRequestsElicitation {
    /** Whether the client supports task-augmented elicitation/create requests. */
    @JsonProperty(value = "create")
    public ClientCapabilitiesTasksRequestsElicitationCreate create;
}
