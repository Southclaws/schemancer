package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** The parameters for a request to elicit additional information from the user via the client. */

@JsonIgnoreProperties(ignoreUnknown = true)
public class ElicitRequestParams extends Object {}
