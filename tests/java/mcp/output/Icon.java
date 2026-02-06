package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.util.ArrayList;
import java.util.List;

/** An optionally-sized icon that can be displayed in a user interface. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class Icon {
    /**
 * Optional MIME type override if the source MIME type is missing or generic.
 * For example: `"image/png"`, `"image/jpeg"`, or `"image/svg+xml"`.
 */
    @JsonProperty(value = "mimeType")
    public String mimeType;
    /**
 * Optional array of strings that specify sizes at which the icon can be used.
 * Each string should be in WxH format (e.g., `"48x48"`, `"96x96"`) or `"any"` for scalable formats like SVG.
 * 
 * If not provided, the client should assume that the icon can be used at any size.
 */
    @JsonProperty(value = "sizes")
    public List<String> sizes = new ArrayList<>();
    /**
 * A standard URI pointing to an icon resource. May be an HTTP/HTTPS URL or a
 * `data:` URI with Base64-encoded image data.
 * 
 * Consumers SHOULD takes steps to ensure URLs serving icons are from the
 * same domain as the client/server or a trusted domain.
 * 
 * Consumers SHOULD take appropriate precautions when consuming SVGs as they can contain
 * executable JavaScript.
 */
    @JsonProperty(value = "src", required = true)
    public URI src;
    /**
 * Optional specifier for the theme this icon is designed for. `light` indicates
 * the icon is designed to be used with a light background, and `dark` indicates
 * the icon is designed to be used with a dark background.
 * 
 * If not provided, the client should assume the icon can be used with any theme.
 */
    @JsonProperty(value = "theme")
    public String theme;
}
