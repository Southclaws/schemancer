package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;


/** Optional annotations for the client. The client can use annotations to inform how objects are used or displayed */
@JsonIgnoreProperties(ignoreUnknown = true)
public class Annotations {
    /**
 * Describes who the intended audience of this object or data is.
 * 
 * It can include multiple entries to indicate content useful for multiple audiences (e.g., `["user", "assistant"]`).
 */
    @JsonProperty(value = "audience")
    public List<Role> audience;
    /**
 * The moment the resource was last modified, as an ISO 8601 formatted string.
 * 
 * Should be an ISO 8601 formatted string (e.g., "2025-01-12T15:00:58Z").
 * 
 * Examples: last activity timestamp in an open file, timestamp when the resource
 * was attached, etc.
 */
    @JsonProperty(value = "lastModified")
    public String lastModified;
    /**
 * Describes how important this data is for operating the server.
 * 
 * A value of 1 means "most important," and indicates that the data is
 * effectively required, while 0 means "least important," and indicates that
 * the data is entirely optional.
 */
    @JsonProperty(value = "priority")
    public Double priority;
}
