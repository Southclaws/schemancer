package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.URI;
import java.time.LocalDate;
import java.time.OffsetDateTime;
import java.util.UUID;


@JsonIgnoreProperties(ignoreUnknown = true)
public class FormattedTypes {
    @JsonProperty(value = "birthDate")
    public LocalDate birthDate;
    @JsonProperty(value = "createdAt", required = true)
    public OffsetDateTime createdAt;
    @JsonProperty(value = "email", required = true)
    public String email;
    @JsonProperty(value = "id", required = true)
    public UUID id;
    @JsonProperty(value = "website")
    public URI website;
}
