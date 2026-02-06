package multi;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.UUID;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Person {
    @JsonProperty(value = "address")
    public Address address;
    @JsonProperty(value = "age")
    public Integer age;
    @JsonProperty(value = "id", required = true)
    public UUID id;
    @JsonProperty(value = "name", required = true)
    public String name;
    @JsonProperty(value = "status", required = true)
    public Status status;

    public Person() {
        this.address = new Address();
    }
}
