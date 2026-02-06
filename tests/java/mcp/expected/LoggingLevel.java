package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonValue;

/**
 * The severity of a log message.
 * 
 * These map to syslog message severities, as specified in RFC-5424:
 * https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1
 */

public enum LoggingLevel {
    ALERT("alert"),
    CRITICAL("critical"),
    DEBUG("debug"),
    EMERGENCY("emergency"),
    ERROR("error"),
    INFO("info"),
    NOTICE("notice"),
    WARNING("warning");

    private final String value;

    LoggingLevel(String value) {
        this.value = value;
    }

    @JsonValue
    public String getValue() {
        return value;
    }
}
