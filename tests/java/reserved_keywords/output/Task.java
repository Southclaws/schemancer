package com.example.generated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Task {
    @JsonProperty(value = "class", required = true)
    private String class_;
    @JsonProperty(value = "import")
    private String import_;
    @JsonProperty(value = "name", required = true)
    private String name;
    @JsonProperty(value = "new")
    private String new_;
    @JsonProperty(value = "type", required = true)
    private String type;

    public String getClass_() {
        return class_;
    }

    public void setClass_(String class_) {
        this.class_ = class_;
    }

    public String getImport() {
        return import_;
    }

    public void setImport(String import_) {
        this.import_ = import_;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getNew() {
        return new_;
    }

    public void setNew(String new_) {
        this.new_ = new_;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }
}
