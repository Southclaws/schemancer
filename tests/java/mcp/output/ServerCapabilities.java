package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.HashMap;
import java.util.Map;

/** Capabilities that a server may support. Known capabilities are defined here, in this schema, but this is not a closed set: any server can define its own, additional capabilities. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ServerCapabilities {
    /** Present if the server supports argument autocompletion suggestions. */
    @JsonProperty(value = "completions")
    public ServerCapabilitiesCompletions completions;
    /** Experimental, non-standard capabilities that the server supports. */
    @JsonProperty(value = "experimental")
    public Map<String, ServerCapabilitiesExperimentalValue> experimental = new HashMap<>();
    /** Present if the server supports sending log messages to the client. */
    @JsonProperty(value = "logging")
    public ServerCapabilitiesLogging logging;
    /** Present if the server offers any prompt templates. */
    @JsonProperty(value = "prompts")
    public ServerCapabilitiesPrompts prompts;
    /** Present if the server offers any resources to read. */
    @JsonProperty(value = "resources")
    public ServerCapabilitiesResources resources;
    /** Present if the server supports task-augmented requests. */
    @JsonProperty(value = "tasks")
    public ServerCapabilitiesTasks tasks;
    /** Present if the server offers any tools to call. */
    @JsonProperty(value = "tools")
    public ServerCapabilitiesTools tools;

    public ServerCapabilities() {
        this.completions = new ServerCapabilitiesCompletions();
        this.logging = new ServerCapabilitiesLogging();
        this.prompts = new ServerCapabilitiesPrompts();
        this.resources = new ServerCapabilitiesResources();
        this.tasks = new ServerCapabilitiesTasks();
        this.tools = new ServerCapabilitiesTools();
    }
}
