package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;


/** Optional metadata to pass through to the LLM provider. The format of this metadata is provider-specific. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CreateMessageRequestParamsMetadata {
}
