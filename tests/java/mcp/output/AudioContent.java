package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Map;


/** Audio provided to or from an LLM. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class AudioContent {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /** Optional annotations for the client. */
    @JsonProperty(value = "annotations")
    public Annotations annotations;
    /** The base64-encoded audio data. */
    @JsonProperty(value = "data", required = true)
    public byte[] data;
    /** The MIME type of the audio. Different providers may support different audio types. */
    @JsonProperty(value = "mimeType", required = true)
    public String mimeType;
    @JsonProperty(value = "type", required = true)
    public String type;
}
