package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Controls tool selection behavior for sampling requests. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ToolChoice {
    /**
 * Controls the tool use ability of the model:
 * - "auto": Model decides whether to use tools (default)
 * - "required": Model MUST use at least one tool before completing
 * - "none": Model MUST NOT use any tools
 */
    @JsonProperty(value = "mode")
    public String mode;
}
