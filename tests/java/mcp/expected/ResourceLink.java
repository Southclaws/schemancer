package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * A resource that the server is capable of reading, included in a prompt or tool call result.
 * 
 * Note: resource links returned by tools are not guaranteed to appear in the results of `resources/list` requests.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ResourceLink {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /** Optional annotations for the client. */
    @JsonProperty(value = "annotations")
    public Annotations annotations;
    /**
 * A description of what this resource represents.
 * 
 * This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
 */
    @JsonProperty(value = "description")
    public String description;
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
    public List<Icon> icons = new ArrayList<>();
    /** The MIME type of this resource, if known. */
    @JsonProperty(value = "mimeType")
    public String mimeType;
    /** Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present). */
    @JsonProperty(value = "name", required = true)
    public String name;
    /**
 * The size of the raw resource content, in bytes (i.e., before base64 encoding or any tokenization), if known.
 * 
 * This can be used by Hosts to display file sizes and estimate context window usage.
 */
    @JsonProperty(value = "size")
    public Integer size;
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
    @JsonProperty(value = "type", required = true)
    public String type;
    /** The URI of this resource. */
    @JsonProperty(value = "uri", required = true)
    public URI uri;

    public ResourceLink() {
        this.annotations = new Annotations();
    }
}
