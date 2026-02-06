package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** A request to retrieve the state of a task. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class GetTaskRequest {
    @JsonProperty(value = "id", required = true)
    public RequestId id;
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public GetTaskRequestParams params;

    public GetTaskRequest() {
        this.params = new GetTaskRequestParams();
    }
}
