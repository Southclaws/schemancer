package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Describes a message returned as part of a prompt.
 * 
 * This is similar to `SamplingMessage`, but also supports the embedding of
 * resources from the MCP server.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class PromptMessage {
    @JsonProperty(value = "content", required = true)
    public ContentBlock content;
    @JsonProperty(value = "role", required = true)
    public Role role;
}
