package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.HashMap;
import java.util.Map;

/** Parameters for a `notifications/message` notification. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class LoggingMessageNotificationParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /** The data to be logged, such as a string message or an object. Any JSON serializable type is allowed here. */
    @JsonProperty(value = "data", required = true)
    public Object data;
    /** The severity of this log message. */
    @JsonProperty(value = "level", required = true)
    public LoggingLevel level;
    /** An optional name of the logger issuing this message. */
    @JsonProperty(value = "logger")
    public String logger;
}
