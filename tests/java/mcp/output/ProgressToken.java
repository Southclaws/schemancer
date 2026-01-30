package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** A progress token, used to associate progress notifications with the original request. */

@JsonIgnoreProperties(ignoreUnknown = true)
public class ProgressToken extends Object {}
