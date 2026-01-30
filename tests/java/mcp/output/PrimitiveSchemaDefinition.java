package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Restricted schema definitions that only allow primitive types
 * without nested objects or arrays.
 */

@JsonIgnoreProperties(ignoreUnknown = true)
public class PrimitiveSchemaDefinition extends Object {}
