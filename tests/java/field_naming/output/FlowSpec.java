package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class FlowSpec {
    @JsonProperty(value = "api-key")
    public String apiKey;
    @JsonProperty(value = "nodes", required = true)
    public List<String> tasks = new ArrayList<>();
    @JsonProperty(value = "regular_field")
    public String regularField;
}
