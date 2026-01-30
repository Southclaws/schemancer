package com.example.mcp;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/** An opaque token used to represent a cursor for pagination. */

@JsonIgnoreProperties(ignoreUnknown = true)
public class Cursor extends String {}
