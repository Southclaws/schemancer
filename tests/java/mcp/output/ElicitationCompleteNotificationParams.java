package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


@JsonIgnoreProperties(ignoreUnknown = true)
public class ElicitationCompleteNotificationParams {
    /** The ID of the elicitation that completed. */
    @JsonProperty(value = "elicitationId", required = true)
    public String elicitationID;
}
