package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** Parameters for a `prompts/get` request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class GetPromptRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public GetPromptRequestParamsMeta meta;
    /** Arguments to use for templating the prompt. */
    @JsonProperty(value = "arguments")
    public Map<String, Object> arguments;
    /** The name of the prompt or prompt template. */
    @JsonProperty(value = "name", required = true)
    public String name;
}
