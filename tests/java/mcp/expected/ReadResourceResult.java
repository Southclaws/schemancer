package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;
import java.util.Map;


/** The server's response to a resources/read request from the client. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ReadResourceResult {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    @JsonProperty(value = "contents", required = true)
    public List<Object> contents;
}
