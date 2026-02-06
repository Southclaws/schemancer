package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Present if the client supports listing roots. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesRoots {
    /** Whether the client supports notifications for changes to the roots list. */
    @JsonProperty(value = "listChanged")
    public Boolean listChanged;
}
