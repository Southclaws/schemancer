package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class LineItem {
    @JsonProperty(value = "price", required = true)
    public double price;
    @JsonProperty(value = "productId", required = true)
    public String productID;
    @JsonProperty(value = "quantity", required = true)
    public int quantity;
}
