package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Present if the client supports elicitation from the server. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ClientCapabilitiesElicitation {
    @JsonProperty(value = "form")
    public ClientCapabilitiesElicitationForm form;
    @JsonProperty(value = "url")
    public ClientCapabilitiesElicitationURL url;

    public ClientCapabilitiesElicitation() {
        this.form = new ClientCapabilitiesElicitationForm();
        this.url = new ClientCapabilitiesElicitationURL();
    }
}
