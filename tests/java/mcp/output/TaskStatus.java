package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

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

    @JsonProperty
    public String getValue() {
        return value;
    }
}
