package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** The argument's information */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CompleteRequestParamsArgument {
    /** The name of the argument */
    @JsonProperty(value = "name", required = true)
    public String name;
    /** The value of the argument to use for completion matching. */
    @JsonProperty(value = "value", required = true)
    public String value;
}
