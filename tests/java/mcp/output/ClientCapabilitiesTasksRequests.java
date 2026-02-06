package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Specifies which request types can be augmented with tasks. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesTasksRequests {
    /** Task support for elicitation-related requests. */
    @JsonProperty(value = "elicitation")
    public ClientCapabilitiesTasksRequestsElicitation elicitation;
    /** Task support for sampling-related requests. */
    @JsonProperty(value = "sampling")
    public ClientCapabilitiesTasksRequestsSampling sampling;

    public ClientCapabilitiesTasksRequests() {
        this.elicitation = new ClientCapabilitiesTasksRequestsElicitation();
        this.sampling = new ClientCapabilitiesTasksRequestsSampling();
    }
}
