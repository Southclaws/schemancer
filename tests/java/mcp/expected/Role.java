package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** The sender or recipient of messages and data in a conversation. */

public enum Role {
    ASSISTANT("assistant"),
    USER("user");

    private final String value;

    Role(String value) {
        this.value = value;
    }

    @JsonProperty
    public String getValue() {
        return value;
    }
}
