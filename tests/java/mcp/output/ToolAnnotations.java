package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Additional properties describing a Tool to clients.
 * 
 * NOTE: all properties in ToolAnnotations are **hints**.
 * They are not guaranteed to provide a faithful description of
 * tool behavior (including descriptive properties like `title`).
 * 
 * Clients should never make tool use decisions based on ToolAnnotations
 * received from untrusted servers.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ToolAnnotations {
    /**
 * If true, the tool may perform destructive updates to its environment.
 * If false, the tool performs only additive updates.
 * 
 * (This property is meaningful only when `readOnlyHint == false`)
 * 
 * Default: true
 */
    @JsonProperty(value = "destructiveHint")
    public Boolean destructiveHint;
    /**
 * If true, calling the tool repeatedly with the same arguments
 * will have no additional effect on its environment.
 * 
 * (This property is meaningful only when `readOnlyHint == false`)
 * 
 * Default: false
 */
    @JsonProperty(value = "idempotentHint")
    public Boolean idempotentHint;
    /**
 * If true, this tool may interact with an "open world" of external
 * entities. If false, the tool's domain of interaction is closed.
 * For example, the world of a web search tool is open, whereas that
 * of a memory tool is not.
 * 
 * Default: true
 */
    @JsonProperty(value = "openWorldHint")
    public Boolean openWorldHint;
    /**
 * If true, the tool does not modify its environment.
 * 
 * Default: false
 */
    @JsonProperty(value = "readOnlyHint")
    public Boolean readOnlyHint;
    /** A human-readable title for the tool. */
    @JsonProperty(value = "title")
    public String title;
}
