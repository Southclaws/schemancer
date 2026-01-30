package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


@JsonIgnoreProperties(ignoreUnknown = true)
public class GetTaskPayloadRequestParams {
    /** The task identifier to retrieve results for. */
    @JsonProperty(value = "taskId", required = true)
    public String taskID;
}
