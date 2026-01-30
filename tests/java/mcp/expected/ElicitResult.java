package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** The client's response to an elicitation request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ElicitResult {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /**
 * The user action in response to the elicitation.
 * - "accept": User submitted the form/confirmed the action
 * - "decline": User explicitly decline the action
 * - "cancel": User dismissed without making an explicit choice
 */
    @JsonProperty(value = "action", required = true)
    public String action;
    /**
 * The submitted form data, only present when action is "accept" and mode was "form".
 * Contains values matching the requested schema.
 * Omitted for out-of-band mode responses.
 */
    @JsonProperty(value = "content")
    public Map<String, Object> content;
}
