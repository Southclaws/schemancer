package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Present if the server supports sending log messages to the client. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ServerCapabilitiesLogging {
}
