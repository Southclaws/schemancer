package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** Execution-related properties for a tool. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ToolExecution {
    /**
 * Indicates whether this tool supports task-augmented execution.
 * This allows clients to handle long-running operations through polling
 * the task system.
 * 
 * - "forbidden": Tool does not support task-augmented execution (default when absent)
 * - "optional": Tool may support task-augmented execution
 * - "required": Tool requires task-augmented execution
 * 
 * Default: "forbidden"
 */
    @JsonProperty(value = "taskSupport")
    public String taskSupport;
}
