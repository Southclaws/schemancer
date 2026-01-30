package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


@JsonIgnoreProperties(ignoreUnknown = true)
public class Error {
    /** The error type that occurred. */
    @JsonProperty(value = "code", required = true)
    public int code;
    /** Additional information about the error. The value of this member is defined by the sender (e.g. detailed error information, nested errors etc.). */
    @JsonProperty(value = "data")
    public Object data;
    /** A short description of the error. The message SHOULD be limited to a concise single sentence. */
    @JsonProperty(value = "message", required = true)
    public String message;
}
