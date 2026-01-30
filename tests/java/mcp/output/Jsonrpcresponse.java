package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** A response to a request, containing either the result or error. */

@JsonIgnoreProperties(ignoreUnknown = true)
public class JSONRPCResponse extends Object {}
