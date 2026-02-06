package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Present if the server offers any resources to read. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ServerCapabilitiesResources {
    /** Whether this server supports notifications for changes to the resource list. */
    @JsonProperty(value = "listChanged")
    public Boolean listChanged;
    /** Whether this server supports subscribing to resource updates. */
    @JsonProperty(value = "subscribe")
    public Boolean subscribe;
}
