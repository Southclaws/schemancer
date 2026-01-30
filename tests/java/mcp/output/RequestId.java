package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** A uniquely identifying ID for a request in JSON-RPC. */

@JsonIgnoreProperties(ignoreUnknown = true)
public class RequestId extends Object {}
