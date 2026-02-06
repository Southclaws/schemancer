package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Whether the client supports context inclusion via includeContext parameter.
 * If not declared, servers SHOULD only use `includeContext: "none"` (or omit it).
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesSamplingContext {
}
