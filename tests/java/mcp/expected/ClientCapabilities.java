package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** Capabilities a client may support. Known capabilities are defined here, in this schema, but this is not a closed set: any client can define its own, additional capabilities. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilities {
    /** Present if the client supports elicitation from the server. */
    @JsonProperty(value = "elicitation")
    public ClientCapabilitiesElicitation elicitation;
    /** Experimental, non-standard capabilities that the client supports. */
    @JsonProperty(value = "experimental")
    public Map<String, Object> experimental;
    /** Present if the client supports listing roots. */
    @JsonProperty(value = "roots")
    public ClientCapabilitiesRoots roots;
    /** Present if the client supports sampling from an LLM. */
    @JsonProperty(value = "sampling")
    public ClientCapabilitiesSampling sampling;
    /** Present if the client supports task-augmented requests. */
    @JsonProperty(value = "tasks")
    public ClientCapabilitiesTasks tasks;
}
