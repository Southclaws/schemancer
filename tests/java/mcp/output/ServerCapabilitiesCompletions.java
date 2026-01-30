package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** Present if the server supports argument autocompletion suggestions. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ServerCapabilitiesCompletions {
}
