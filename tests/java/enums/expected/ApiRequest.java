package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class ApiRequest {
    @JsonProperty(value = "body")
    public String body;
    @JsonProperty(value = "method", required = true)
    public HttpMethod method;
    @JsonProperty(value = "url", required = true)
    public String url;
}
