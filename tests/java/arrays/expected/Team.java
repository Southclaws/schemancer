package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Team {
    @JsonProperty(value = "members", required = true)
    public List<Person> members = new ArrayList<>();
    @JsonProperty(value = "name", required = true)
    public String name;
}
