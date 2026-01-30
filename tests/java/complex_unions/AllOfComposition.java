package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class AllOfComposition {
    @JsonProperty(value = "id", required = true)
    public String id;
    @JsonProperty(value = "name", required = true)
    public String name;
    @JsonProperty(value = "timestamp")
    public String timestamp;
}

@JsonIgnoreProperties(ignoreUnknown = true)
public class MultiType extends Object {
}

@JsonIgnoreProperties(ignoreUnknown = true)
public class NestedUnion {
    @JsonProperty(value = "data")
    public Object data;
}

@JsonIgnoreProperties(ignoreUnknown = true)
public class ObjectUnion extends Object {
}

@JsonIgnoreProperties(ignoreUnknown = true)
public class StringOrNull extends Object {
}

@JsonIgnoreProperties(ignoreUnknown = true)
public class StringOrNumber extends Object {
}
