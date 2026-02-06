package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.ArrayList;
import java.util.List;

/** Server configuration with default values */
@JsonIgnoreProperties(ignoreUnknown = true)
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
    public Integer maxRetries = 3;
    @JsonProperty(value = "port")
    @JsonSetter(nulls = Nulls.SKIP)
    public Integer port = 8080;
    @JsonProperty(value = "tags")
    public List<String> tags = new ArrayList<>();
    @JsonProperty(value = "timeout")
    @JsonSetter(nulls = Nulls.SKIP)
    public Double timeout = 30;
}
