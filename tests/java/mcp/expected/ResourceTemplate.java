package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/** A template description for resources available on the server. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class ResourceTemplate {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public Map<String, Object> meta = new HashMap<>();
    /** Optional annotations for the client. */
    @JsonProperty(value = "annotations")
    public Annotations annotations;
    /**
 * A description of what this template is for.
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
    /** The MIME type for all resources that match this template. This should only be included if all resources matching this template have the same type. */
    @JsonProperty(value = "mimeType")
    public String mimeType;
    /** Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present). */
    @JsonProperty(value = "name", required = true)
    public String name;
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
    /** A URI template (according to RFC 6570) that can be used to construct resource URIs. */
    @JsonProperty(value = "uriTemplate", required = true)
    public String uritemplate;

    public ResourceTemplate() {
        this.annotations = new Annotations();
    }
}
