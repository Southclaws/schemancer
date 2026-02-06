package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** Parameters for a `logging/setLevel` request. */
@JsonIgnoreProperties(ignoreUnknown = true)
public class SetLevelRequestParams {
    /** See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage. */
    @JsonProperty(value = "_meta")
    public SetLevelRequestParamsMeta meta;
    /** The level of logging that the client wants to receive from the server. The server should send all logs at this level and higher (i.e., more severe) to the client as notifications/message. */
    @JsonProperty(value = "level", required = true)
    public LoggingLevel level;

    public SetLevelRequestParams() {
        this.meta = new SetLevelRequestParamsMeta();
    }
}
