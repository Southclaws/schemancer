package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * This notification can be sent by either side to indicate that it is cancelling a previously-issued request.
 * 
 * The request SHOULD still be in-flight, but due to communication latency, it is always possible that this notification MAY arrive after the request has already finished.
 * 
 * This notification indicates that the result will be unused, so any associated processing SHOULD cease.
 * 
 * A client MUST NOT attempt to cancel its `initialize` request.
 * 
 * For task cancellation, use the `tasks/cancel` request instead of this notification.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CancelledNotification {
    @JsonProperty(value = "jsonrpc", required = true)
    public String jsonrpc;
    @JsonProperty(value = "method", required = true)
    public String method;
    @JsonProperty(value = "params", required = true)
    public CancelledNotificationParams params;

    public CancelledNotification() {
        this.params = new CancelledNotificationParams();
    }
}
