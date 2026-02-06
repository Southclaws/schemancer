package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Parameters for an `initialize` request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class InitializeRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public InitializeRequestParamsMeta meta;
    @JsonProperty(value = "capabilities", required = true)
    public ClientCapabilities capabilities;
    @JsonProperty(value = "clientInfo", required = true)
    public Implementation clientInfo;
    /** The latest version of the Model Context Protocol that the client supports. The client MAY decide to support older versions as well. */
    @JsonProperty(value = "protocolVersion", required = true)
    public String protocolVersion;

    public InitializeRequestParams() {
        this.meta = new InitializeRequestParamsMeta();
        this.capabilities = new ClientCapabilities();
        this.clientInfo = new Implementation();
    }
}
