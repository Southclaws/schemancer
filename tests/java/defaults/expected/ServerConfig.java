package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

/** Server configuration with default values */
@JsonIgnoreProperties(ignoreUnknown = true)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class ServerConfig {
    @JsonProperty(value = "debug")
    @JsonSetter(nulls = Nulls.SKIP)
    public Boolean debug = false;
    /** The hostname to bind to */
    @JsonProperty(value = "host", required = true)
    @JsonSetter(nulls = Nulls.SKIP)
    public String host = "localhost";
    @JsonProperty(value = "maxRetries")
    @JsonSetter(nulls = Nulls.SKIP)
    public Long maxRetries = 3L;
    @JsonProperty(value = "port")
    @JsonSetter(nulls = Nulls.SKIP)
    public Long port = 8080L;
    @JsonProperty(value = "tags")
    public List<String> tags;
    @JsonProperty(value = "timeout")
    @JsonSetter(nulls = Nulls.SKIP)
    public Double timeout = 30;
}
