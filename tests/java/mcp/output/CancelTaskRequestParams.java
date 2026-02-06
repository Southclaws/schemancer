package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class CancelTaskRequestParams {
    /** The task identifier to cancel. */
    @JsonProperty(value = "taskId", required = true)
    public String taskID;
}
