package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;
import java.util.Map;


/** Definition for a tool the client can call. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class Tool {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta;
    /**
 * Optional additional tool information.
 * 
 * Display name precedence order is: title, annotations.title, then name.
 */
    @JsonProperty(value = "annotations")
    public ToolAnnotations annotations;
    /**
 * A human-readable description of the tool.
 * 
 * This can be used by clients to improve the LLM's understanding of available tools. It can be thought of like a "hint" to the model.
 */
    @JsonProperty(value = "description")
    public String description;
    /** Execution-related properties for this tool. */
    @JsonProperty(value = "execution")
    public ToolExecution execution;
    /**
 * Optional set of sized icons that the client can display in a user interface.
 * 
 * Clients that support rendering icons MUST support at least the following MIME types:
 * - `image/png` - PNG images (safe, universal compatibility)
 * - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
 * 
 * Clients that support rendering icons SHOULD also support:
 * - `image/svg+xml` - SVG images (scalable but requires security precautions)
 * - `image/webp` - WebP images (modern, efficient format)
 */
    @JsonProperty(value = "icons")
    public List<Icon> icons;
    /** A JSON Schema object defining the expected parameters for the tool. */
    @JsonProperty(value = "inputSchema", required = true)
    public ToolInputSchema inputSchema;
    /** Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present). */
    @JsonProperty(value = "name", required = true)
    public String name;
    /**
 * An optional JSON Schema object defining the structure of the tool's output returned in
 * the structuredContent field of a CallToolResult.
 * 
 * Defaults to JSON Schema 2020-12 when no explicit $schema is provided.
 * Currently restricted to type: "object" at the root level.
 */
    @JsonProperty(value = "outputSchema")
    public ToolOutputSchema outputSchema;
    /**
 * Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
 * even by those unfamiliar with domain-specific terminology.
 * 
 * If not provided, the name should be used for display (except for Tool,
 * where `annotations.title` should be given precedence over using `name`,
 * if present).
 */
    @JsonProperty(value = "title")
    public String title;
}
