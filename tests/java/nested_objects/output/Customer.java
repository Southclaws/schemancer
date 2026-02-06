package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Customer {
    @JsonProperty(value = "email")
    public String email;
    @JsonProperty(value = "id", required = true)
    public String id;
    @JsonProperty(value = "name", required = true)
    public String name;
}
