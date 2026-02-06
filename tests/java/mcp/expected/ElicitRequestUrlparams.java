package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;

/** The parameters for a request to elicit information from the user via a URL in the client. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ElicitRequestURLParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public ElicitRequestURLParamsMeta meta;
    /**
 * The ID of the elicitation, which must be unique within the context of the server.
 * The client MUST treat this ID as an opaque value.
 */
    @JsonProperty(value = "elicitationId", required = true)
    public String elicitationID;
    /** The message to present to the user explaining why the interaction is needed. */
    @JsonProperty(value = "message", required = true)
    public String message;
    /** The elicitation mode. */
    @JsonProperty(value = "mode", required = true)
    public String mode;
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
    /** The URL that the user should navigate to. */
    @JsonProperty(value = "url", required = true)
    public URI url;

    public ElicitRequestURLParams() {
        this.meta = new ElicitRequestURLParamsMeta();
        this.task = new TaskMetadata();
    }
}
