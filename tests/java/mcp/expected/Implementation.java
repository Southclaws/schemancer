package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.List;


/** Describes the MCP implementation. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class Implementation {
    /**
 * An optional human-readable description of what this implementation does.
 * 
 * This can be used by clients or servers to provide context about their purpose
 * and capabilities. For example, a server might describe the types of resources
 * or tools it provides, while a client might describe its intended use case.
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
    public List<Icon> icons;
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
    @JsonProperty(value = "version", required = true)
    public String version;
    /** An optional URL of the website for this implementation. */
    @JsonProperty(value = "websiteUrl")
    public URI websiteURL;
}
