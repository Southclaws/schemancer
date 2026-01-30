package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/**
 * Metadata for associating messages with a task.
 * Include this in the `_meta` field under the key `io.modelcontextprotocol/related-task`.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class RelatedTaskMetadata {
    /** The task identifier this message is associated with. */
    @JsonProperty(value = "taskId", required = true)
    public String taskID;
}
