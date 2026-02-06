package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/** The server's response to a prompts/get request from the client. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class GetPromptResult {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /** An optional description for the prompt. */
    @JsonProperty(value = "description")
    public String description;
    @JsonProperty(value = "messages", required = true)
    public List<PromptMessage> messages = new ArrayList<>();
}
