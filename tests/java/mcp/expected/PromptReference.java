package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Identifies a prompt. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class PromptReference {
    /** Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present). */
    @JsonProperty(value = "name", required = true)
    public String name;
    /**
 * Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
 * even by those unfamiliar with domain-specific terminology.
 * 
 * If not provided, the name should be used for display (except for Tool,
 * where `annotations.title` should be given precedence over using `name`,
 * if present).
 */
    @JsonProperty(value = "title")
    public String title;
    @JsonProperty(value = "type", required = true)
    public String type;
}
