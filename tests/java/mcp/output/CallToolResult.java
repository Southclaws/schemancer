package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/** The server's response to a tool call. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CallToolResult {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /** A list of content objects that represent the unstructured result of the tool call. */
    @JsonProperty(value = "content", required = true)
    public List<ContentBlock> content = new ArrayList<>();
    /**
 * Whether the tool call ended in an error.
 * 
 * If not set, this is assumed to be false (the call was successful).
 * 
 * Any errors that originate from the tool SHOULD be reported inside the result
 * object, with `isError` set to true, _not_ as an MCP protocol-level error
 * response. Otherwise, the LLM would not be able to see that an error occurred
 * and self-correct.
 * 
 * However, any errors in _finding_ the tool, an error indicating that the
 * server does not support tool calls, or any other exceptional conditions,
 * should be reported as an MCP error response.
 */
    @JsonProperty(value = "isError")
    public Boolean isError;
    /** An optional JSON object that represents the structured result of the tool call. */
    @JsonProperty(value = "structuredContent")
    public Map<String, Object> structuredContent = new HashMap<>();
}
