package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** The parameters for a request to elicit non-sensitive information from the user via a form in the client. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ElicitRequestFormParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public ElicitRequestFormParamsMeta meta;
    /** The message to present to the user describing what information is being requested. */
    @JsonProperty(value = "message", required = true)
    public String message;
    /** The elicitation mode. */
    @JsonProperty(value = "mode")
    public String mode;
    /**
 * A restricted subset of JSON Schema.
 * Only top-level properties are allowed, without nesting.
 */
    @JsonProperty(value = "requestedSchema", required = true)
    public ElicitRequestFormParamsRequestedSchema requestedSchema;
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
}
