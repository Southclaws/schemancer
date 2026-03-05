package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** An error response that indicates that the server requires the client to provide additional information via an elicitation request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class URLElicitationRequiredError {
    @JsonProperty(value = "error", required = true)
    public Object error;
    @JsonProperty(value = "id")
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
}
