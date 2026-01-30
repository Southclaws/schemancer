package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


@JsonIgnoreProperties(ignoreUnknown = true)
public class Person {
    @JsonProperty(value = "name", required = true)
    public String name;
    @JsonProperty(value = "scores")
    public List<Double> scores;
    @JsonProperty(value = "tags", required = true)
    public List<String> tags;
}
