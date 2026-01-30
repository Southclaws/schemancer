package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** Parameters for a `completion/complete` request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CompleteRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public CompleteRequestParamsMeta meta;
    /** The argument's information */
    @JsonProperty(value = "argument", required = true)
    public CompleteRequestParamsArgument argument;
    /** Additional, optional context for completions */
    @JsonProperty(value = "context")
    public CompleteRequestParamsContext context;
    @JsonProperty(value = "ref", required = true)
    public Object ref;
}
