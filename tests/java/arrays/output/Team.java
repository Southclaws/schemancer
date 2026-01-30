package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


@JsonIgnoreProperties(ignoreUnknown = true)
public class Team {
    @JsonProperty(value = "members", required = true)
    public List<Person> members;
    @JsonProperty(value = "name", required = true)
    public String name;
}
