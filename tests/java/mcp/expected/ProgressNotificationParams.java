package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** Parameters for a `notifications/progress` notification. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ProgressNotificationParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /** An optional message describing the current progress. */
    @JsonProperty(value = "message")
    public String message;
    /** The progress thus far. This should increase every time progress is made, even if the total is unknown. */
    @JsonProperty(value = "progress", required = true)
    public double progress;
    /** The progress token which was given in the initial request, used to associate this notification with the request that is proceeding. */
    @JsonProperty(value = "progressToken", required = true)
    public ProgressToken progressToken;
    /** Total number of items to process (or total progress required), if known. */
    @JsonProperty(value = "total")
    public Double total;
}
