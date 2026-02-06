package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Whether this client supports tasks/list. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesTasksList {
}
