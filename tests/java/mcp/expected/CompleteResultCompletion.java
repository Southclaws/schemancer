package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


@JsonIgnoreProperties(ignoreUnknown = true)
public class CompleteResultCompletion {
    /** Indicates whether there are additional completion options beyond those provided in the current response, even if the exact total is unknown. */
    @JsonProperty(value = "hasMore")
    public Boolean hasMore;
    /** The total number of completion options available. This can exceed the number of values actually sent in the response. */
    @JsonProperty(value = "total")
    public Integer total;
    /** An array of completion values. Must not exceed 100 items. */
    @JsonProperty(value = "values", required = true)
    public List<String> values;
}
