package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Customer {
    @JsonProperty(value = "email")
    public String email;
    @JsonProperty(value = "id", required = true)
    public String id;
    @JsonProperty(value = "name", required = true)
    public String name;
}

@JsonIgnoreProperties(ignoreUnknown = true)
public class LineItem {
    @JsonProperty(value = "price", required = true)
    public double price;
    @JsonProperty(value = "productId", required = true)
    public String productId;
    @JsonProperty(value = "quantity", required = true)
    public int quantity;
}

@JsonIgnoreProperties(ignoreUnknown = true)
public class Order {
    @JsonProperty(value = "customer", required = true)
    public Customer customer;
    @JsonProperty(value = "id", required = true)
    public String id;
    @JsonProperty(value = "items", required = true)
    public List<LineItem> items;
    @JsonProperty(value = "total")
    public Double total;
}
