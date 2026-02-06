package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/** The result of a tool use, provided by the user back to the assistant. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ToolResultContent {
    /**
 * Optional metadata about the tool result. Clients SHOULD preserve this field when
 * including tool results in subsequent sampling requests to enable caching optimizations.
 * 
 * See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
 */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /**
 * The unstructured result content of the tool use.
 * 
 * This has the same format as CallToolResult.content and can include text, images,
 * audio, resource links, and embedded resources.
 */
    @JsonProperty(value = "content", required = true)
    public List<ContentBlock> content = new ArrayList<>();
    /**
 * Whether the tool use resulted in an error.
 * 
 * If true, the content typically describes the error that occurred.
 * Default: false
 */
    @JsonProperty(value = "isError")
    public Boolean isError;
    /**
 * An optional structured result object.
 * 
 * If the tool defined an outputSchema, this SHOULD conform to that schema.
 */
    @JsonProperty(value = "structuredContent")
    public Map<String, Object> structuredContent = new HashMap<>();
    /**
 * The ID of the tool use this result corresponds to.
 * 
 * This MUST match the ID from a previous ToolUseContent.
 */
    @JsonProperty(value = "toolUseId", required = true)
    public String toolUseID;
    @JsonProperty(value = "type", required = true)
    public String type;
}
