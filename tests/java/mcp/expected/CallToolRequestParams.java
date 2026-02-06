package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.HashMap;
import java.util.Map;

/** Parameters for a `tools/call` request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CallToolRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public CallToolRequestParamsMeta meta;
    /** Arguments to use for the tool call. */
    @JsonProperty(value = "arguments")
    public Map<String, Object> arguments = new HashMap<>();
    /** The name of the tool. */
    @JsonProperty(value = "name", required = true)
    public String name;
    /**
 * If specified, the caller is requesting task-augmented execution for this request.
 * The request will return a CreateTaskResult immediately, and the actual result can be
 * retrieved later via tasks/result.
 * 
 * Task augmentation is subject to capability negotiation - receivers MUST declare support
 * for task augmentation of specific request types in their capabilities.
 */
    @JsonProperty(value = "task")
    public TaskMetadata task;

    public CallToolRequestParams() {
        this.meta = new CallToolRequestParamsMeta();
        this.task = new TaskMetadata();
    }
}
