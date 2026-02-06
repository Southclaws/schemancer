package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Common params for any task-augmented request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class TaskAugmentedRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public TaskAugmentedRequestParamsMeta meta;
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

    public TaskAugmentedRequestParams() {
        this.meta = new TaskAugmentedRequestParamsMeta();
        this.task = new TaskMetadata();
    }
}
