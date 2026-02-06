package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Data associated with a task. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class Task {
    /** ISO 8601 timestamp when the task was created. */
    @JsonProperty(value = "createdAt", required = true)
    public String createdAt;
    /** ISO 8601 timestamp when the task was last updated. */
    @JsonProperty(value = "lastUpdatedAt", required = true)
    public String lastUpdatedAt;
    /** Suggested polling interval in milliseconds. */
    @JsonProperty(value = "pollInterval")
    public Integer pollInterval;
    /** Current task state. */
    @JsonProperty(value = "status", required = true)
    public TaskStatus status;
    /**
 * Optional human-readable message describing the current task state.
 * This can provide context for any status, including:
 * - Reasons for "cancelled" status
 * - Summaries for "completed" status
 * - Diagnostic information for "failed" status (e.g., error details, what went wrong)
 */
    @JsonProperty(value = "statusMessage")
    public String statusMessage;
    /** The task identifier. */
    @JsonProperty(value = "taskId", required = true)
    public String taskID;
    /** Actual retention duration from creation in milliseconds, null for unlimited. */
    @JsonProperty(value = "ttl", required = true)
    public int ttl;
}
