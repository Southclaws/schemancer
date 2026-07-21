package com.example.generated;

import com.example.AccountId;
import com.example.UserId;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class User {
    @JsonProperty(value = "accountId", required = true)
    public AccountId accountID;
    @JsonProperty(value = "age", required = true)
    public long age;
    @JsonProperty(value = "id", required = true)
    public UserId id;
    @JsonProperty(value = "name", required = true)
    public String name;
}
