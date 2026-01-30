package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** Present if the client supports sampling from an LLM. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesSampling {
    /**
 * Whether the client supports context inclusion via includeContext parameter.
 * If not declared, servers SHOULD only use `includeContext: "none"` (or omit it).
 */
    @JsonProperty(value = "context")
    public ClientCapabilitiesSamplingContext context;
    /** Whether the client supports tool use via tools and toolChoice parameters. */
    @JsonProperty(value = "tools")
    public ClientCapabilitiesSamplingTools tools;
}
