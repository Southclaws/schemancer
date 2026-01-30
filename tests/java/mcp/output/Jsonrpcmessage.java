package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Refers to any valid JSON-RPC object that can be decoded off the wire, or encoded to be sent. */

@JsonIgnoreProperties(ignoreUnknown = true)
public class JSONRPCMessage extends Object {}
