package com.example.generated;

import com.fasterxml.jackson.annotation.JsonValue;

public enum Color {
    RED("red"),
    GREEN("green"),
    BLUE("blue"),
    YELLOW("yellow");

    private final String value;

    Color(String value) {
        this.value = value;
    }

    @JsonValue
    public String getValue() {
        return value;
    }
}
