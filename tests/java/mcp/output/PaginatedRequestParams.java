package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Common parameters for paginated requests. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class PaginatedRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public PaginatedRequestParamsMeta meta;
    /**
 * An opaque token representing the current pagination position.
 * If provided, the server should return results starting after this cursor.
 */
    @JsonProperty(value = "cursor")
    public String cursor;

    public PaginatedRequestParams() {
        this.meta = new PaginatedRequestParamsMeta();
    }
}
