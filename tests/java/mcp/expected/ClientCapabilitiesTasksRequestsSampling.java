package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Task support for sampling-related requests. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesTasksRequestsSampling {
    /** Whether the client supports task-augmented sampling/createMessage requests. */
    @JsonProperty(value = "createMessage")
    public ClientCapabilitiesTasksRequestsSamplingCreateMessage createMessage;

    public ClientCapabilitiesTasksRequestsSampling() {
        this.createMessage = new ClientCapabilitiesTasksRequestsSamplingCreateMessage();
    }
}
