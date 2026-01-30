package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** Parameters for a `notifications/cancelled` notification. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CancelledNotificationParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /** An optional string describing the reason for the cancellation. This MAY be logged or presented to the user. */
    @JsonProperty(value = "reason")
    public String reason;
    /**
 * The ID of the request to cancel.
 * 
 * This MUST correspond to the ID of a request previously issued in the same direction.
 * This MUST be provided for cancelling non-task requests.
 * This MUST NOT be used for cancelling tasks (use the `tasks/cancel` request instead).
 */
    @JsonProperty(value = "requestId")
    public RequestId requestID;
}
