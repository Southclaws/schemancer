package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Whether this server supports tasks/list. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ServerCapabilitiesTasksList {
}
