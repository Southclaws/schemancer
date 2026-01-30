package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** A successful (non-error) response to a request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class JSONRPCResultResponse {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "result", required = true)
    public Result result;
}
