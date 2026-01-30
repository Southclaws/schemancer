package multi;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

public enum Status {
    ACTIVE("active"),
    INACTIVE("inactive"),
    PENDING("pending");

    private final String value;

    Status(String value) {
        this.value = value;
    }

    @JsonProperty
    public String getValue() {
        return value;
    }
}
