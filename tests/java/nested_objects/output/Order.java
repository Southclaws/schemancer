package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.ArrayList;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Order {
    @JsonProperty(value = "customer", required = true)
    public Customer customer;
    @JsonProperty(value = "id", required = true)
    public String id;
    @JsonProperty(value = "items", required = true)
    public List<LineItem> items = new ArrayList<>();
    @JsonProperty(value = "total")
    public Double total;

    public Order() {
        this.customer = new Customer();
    }
}
