package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** After receiving an initialize request from the client, the server sends this response. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class InitializeResult {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    @JsonProperty(value = "capabilities", required = true)
    public ServerCapabilities capabilities;
    /**
 * Instructions describing how to use the server and its features.
 * 
 * This can be used by clients to improve the LLM's understanding of available tools, resources, etc. It can be thought of like a "hint" to the model. For example, this information MAY be added to the system prompt.
 */
    @JsonProperty(value = "instructions")
    public String instructions;
    /** The version of the Model Context Protocol that the server wants to use. This may not match the version that the client requested. If the client cannot support this version, it MUST disconnect. */
    @JsonProperty(value = "protocolVersion", required = true)
    public String protocolVersion;
    @JsonProperty(value = "serverInfo", required = true)
    public Implementation serverInfo;
}
