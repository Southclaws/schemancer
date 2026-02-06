package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonValue;

/** The sender or recipient of messages and data in a conversation. */

public enum Role {
    ASSISTANT("assistant"),
    USER("user");

    private final String value;

    Role(String value) {
        this.value = value;
    }

    @JsonValue
    public String getValue() {
        return value;
    }
}
