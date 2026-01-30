package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


/** Parameters for a `sampling/createMessage` request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CreateMessageRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public CreateMessageRequestParamsMeta meta;
    /**
 * A request to include context from one or more MCP servers (including the caller), to be attached to the prompt.
 * The client MAY ignore this request.
 * 
 * Default is "none". Values "thisServer" and "allServers" are soft-deprecated. Servers SHOULD only use these values if the client
 * declares ClientCapabilities.sampling.context. These values may be removed in future spec releases.
 */
    @JsonProperty(value = "includeContext")
    public String includeContext;
    /**
 * The requested maximum number of tokens to sample (to prevent runaway completions).
 * 
 * The client MAY choose to sample fewer tokens than the requested maximum.
 */
    @JsonProperty(value = "maxTokens", required = true)
    public int maxTokens;
    @JsonProperty(value = "messages", required = true)
    public List<SamplingMessage> messages;
    /** Optional metadata to pass through to the LLM provider. The format of this metadata is provider-specific. */
    @JsonProperty(value = "metadata")
    public CreateMessageRequestParamsMetadata metadata;
    /** The server's preferences for which model to select. The client MAY ignore these preferences. */
    @JsonProperty(value = "modelPreferences")
    public ModelPreferences modelPreferences;
    @JsonProperty(value = "stopSequences")
    public List<String> stopSequences;
    /** An optional system prompt the server wants to use for sampling. The client MAY modify or omit this prompt. */
    @JsonProperty(value = "systemPrompt")
    public String systemPrompt;
    /**
 * If specified, the caller is requesting task-augmented execution for this request.
 * The request will return a CreateTaskResult immediately, and the actual result can be
 * retrieved later via tasks/result.
 * 
 * Task augmentation is subject to capability negotiation - receivers MUST declare support
 * for task augmentation of specific request types in their capabilities.
 */
    @JsonProperty(value = "task")
    public TaskMetadata task;
    @JsonProperty(value = "temperature")
    public Double temperature;
    /**
 * Controls how the model uses tools.
 * The client MUST return an error if this field is provided but ClientCapabilities.sampling.tools is not declared.
 * Default is `{ mode: "auto" }`.
 */
    @JsonProperty(value = "toolChoice")
    public ToolChoice toolChoice;
    /**
 * Tools that the model may use during generation.
 * The client MUST return an error if this field is provided but ClientCapabilities.sampling.tools is not declared.
 */
    @JsonProperty(value = "tools")
    public List<Tool> tools;
}
