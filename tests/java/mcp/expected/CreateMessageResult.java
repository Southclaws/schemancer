package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/**
 * The client's response to a sampling/createMessage request from the server.
 * The client should inform the user before returning the sampled message, to allow them
 * to inspect the response (human in the loop) and decide whether to allow the server to see it.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CreateMessageResult {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    @JsonProperty(value = "content", required = true)
    public Object content;
    /** The name of the model that generated the message. */
    @JsonProperty(value = "model", required = true)
    public String model;
    @JsonProperty(value = "role", required = true)
    public Role role;
    /**
 * The reason why sampling stopped, if known.
 * 
 * Standard values:
 * - "endTurn": Natural end of the assistant's turn
 * - "stopSequence": A stop sequence was encountered
 * - "maxTokens": Maximum token limit was reached
 * - "toolUse": The model wants to use one or more tools
 * 
 * This field is an open string to allow for provider-specific stop reasons.
 */
    @JsonProperty(value = "stopReason")
    public String stopReason;
}
