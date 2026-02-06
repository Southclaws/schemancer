package multi;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Address {
    @JsonProperty(value = "city", required = true)
    public String city;
    @JsonProperty(value = "country")
    public String country;
    @JsonProperty(value = "street", required = true)
    public String street;
}
