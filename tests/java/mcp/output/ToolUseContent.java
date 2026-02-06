package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.HashMap;
import java.util.Map;

/** A request from the assistant to call a tool. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ToolUseContent {
    /**
 * Optional metadata about the tool use. Clients SHOULD preserve this field when
 * including tool uses in subsequent sampling requests to enable caching optimizations.
 * 
 * See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
 */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /**
 * A unique identifier for this tool use.
 * 
 * This ID is used to match tool results to their corresponding tool uses.
 */
    @JsonProperty(value = "id", required = true)
    public String id;
    /** The arguments to pass to the tool, conforming to the tool's input schema. */
    @JsonProperty(value = "input", required = true)
    public Map<String, Object> input = new HashMap<>();
    /** The name of the tool to call. */
    @JsonProperty(value = "name", required = true)
    public String name;
    @JsonProperty(value = "type", required = true)
    public String type;
}
