package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** A reference to a resource or resource template definition. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ResourceTemplateReference {
    @JsonProperty(value = "type", required = true)
    public String type;
    /** The URI or URI template of the resource. */
    @JsonProperty(value = "uri", required = true)
    public String uri;
}
