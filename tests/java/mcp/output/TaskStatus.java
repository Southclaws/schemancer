package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonValue;

/** The status of a task. */

public enum TaskStatus {
    CANCELLED("cancelled"),
    COMPLETED("completed"),
    FAILED("failed"),
    INPUT_REQUIRED("input_required"),
    WORKING("working");

    private final String value;

    TaskStatus(String value) {
        this.value = value;
    }

    @JsonValue
    public String getValue() {
        return value;
    }
}
