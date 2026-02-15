# schemancer

> Used in production at [Storyden](https://github.com/Southclaws/storyden) - an open source forum+wiki written in Go!

**JSON Schema Code Generator** — Generate type-safe code from JSON Schema definitions.

## Features

- **Multi-language support**: Go, TypeScript, (Types, Zod), Java, Python (Pydantic v2)
- **Discriminated unions**: First-class support for tagged unions with type guards and pattern matching
- **Format mappings**: Configurable type mappings for `uuid`, `date-time`, `email`, and other formats
- **Config file**: Generate multiple languages from a single schema with `schemancer.yaml`

## Why?

I use a ton of code-generation in my workflows, OpenAPI, JSONSchema, Protocol Buffers, etc. I found a lot of the JSONSchema tooling to lack support for a tool I reach for a lot while defining data models: discriminated unions. I also wanted a consistent cross-language tool and output style, Schemancer supports 4 languages.

[More details on why, and how this codebase is maintained.](#more-on-why)

## Installation

```bash
go install github.com/Southclaws/schemancer@latest
```

## Usage

### Single Language

```bash
# Generate Go types
schemancer schema.yaml golang output.go --package=models

# Generate TypeScript types
schemancer schema.yaml typescript output.ts

# Generate Java classes
schemancer schema.yaml java output.java --package=com.example

# Generate Python Pydantic models
schemancer schema.yaml python output.py

# Generate TypeScript Zod schemas
schemancer schema.yaml typescript-zod output.ts

# Output to stdout
schemancer schema.yaml typescript -
```

### Multi-Language (Config File)

Create a `schemancer.yaml`:

```yaml
golang:
  output: "./generated"
  package: "models"

typescript:
  output: "./generated"

typescript-zod:
  output: "./generated"

java:
  output: "./generated"
  package: "com.example.models"

python:
  output: "./generated"
```

Then run:

```bash
schemancer schema.yaml
```

This generates all configured languages in one command.

## Discriminated Unions

schemancer has first-class support for discriminated unions (tagged unions). Given a schema like:

```yaml
$defs:
  Event:
    oneOf:
      - $ref: "#/$defs/CreatedEvent"
      - $ref: "#/$defs/UpdatedEvent"
      - $ref: "#/$defs/DeletedEvent"

  CreatedEvent:
    type: object
    required: [type, id, name]
    properties:
      type: { const: "created" }
      id: { type: string }
      name: { type: string }

  UpdatedEvent:
    type: object
    required: [type, id, changes]
    properties:
      type: { const: "updated" }
      id: { type: string }
      changes: { type: object }

  DeletedEvent:
    type: object
    required: [type, id]
    properties:
      type: { const: "deleted" }
      id: { type: string }
      reason: { type: string }
```

### Generated TypeScript

```typescript
export interface CreatedEvent {
  type: "created";
  id: string;
  name: string;
}

export interface UpdatedEvent {
  type: "updated";
  id: string;
  changes: Record<string, unknown>;
}

export interface DeletedEvent {
  type: "deleted";
  id: string;
  reason?: string;
}

export type Event = CreatedEvent | UpdatedEvent | DeletedEvent;

export function isCreatedEvent(value: Event): value is CreatedEvent {
  return value.type === "created";
}
// ... type guards for each variant
```

### Generated TypeScript Zod

```typescript
import { z } from "zod";

export const CreatedEventSchema = z.object({
  type: z.literal("created"),
  id: z.string(),
  name: z.string(),
});
export type CreatedEvent = z.infer<typeof CreatedEventSchema>;

export const UpdatedEventSchema = z.object({
  type: z.literal("updated"),
  id: z.string(),
  changes: z.record(z.string(), z.unknown()),
});
export type UpdatedEvent = z.infer<typeof UpdatedEventSchema>;

export const DeletedEventSchema = z.object({
  type: z.literal("deleted"),
  id: z.string(),
  reason: z.string().optional(),
});
export type DeletedEvent = z.infer<typeof DeletedEventSchema>;

export const EventSchema = z.discriminatedUnion("type", [
  CreatedEventSchema,
  UpdatedEventSchema,
  DeletedEventSchema,
]);
export type Event = z.infer<typeof EventSchema>;
```

### Generated Python (Pydantic v2)

```python
from typing import Annotated, Literal, Union
from pydantic import BaseModel, Field

class CreatedEvent(BaseModel):
    type: Literal["created"]
    id: str
    name: str

class UpdatedEvent(BaseModel):
    type: Literal["updated"]
    id: str
    changes: dict[str, Any]

class DeletedEvent(BaseModel):
    type: Literal["deleted"]
    id: str
    reason: str | None = None

Event = Annotated[
    Union[CreatedEvent, UpdatedEvent, DeletedEvent],
    Field(discriminator="type"),
]
```

### Generated Go

```go
type Event interface {
    EventType() string
    isEvent()
}

type EventWrapper struct { Event }

func (w *EventWrapper) UnmarshalJSON(data []byte) error {
    // Automatic unmarshaling based on discriminator
}

type CreatedEvent struct {
    Type string `json:"type"`
    ID   string `json:"id"`
    Name string `json:"name"`
}

func (CreatedEvent) EventType() string { return "created" }
func (CreatedEvent) isEvent() {}
// ... other variants
```

### Generated Java

```java
@JsonTypeInfo(use = JsonTypeInfo.Id.NAME, property = "type")
@JsonSubTypes({
    @JsonSubTypes.Type(value = CreatedEvent.class, name = "created"),
    @JsonSubTypes.Type(value = UpdatedEvent.class, name = "updated"),
    @JsonSubTypes.Type(value = DeletedEvent.class, name = "deleted")
})
public sealed interface Event permits CreatedEvent, UpdatedEvent, DeletedEvent {
    String type();
}
```

## Configuration Options

### Go

| Option            | Description                                           |
| ----------------- | ----------------------------------------------------- |
| `package`         | Package name for generated code                       |
| `optional_style`  | `pointer` (default) or `opt` (uses `opt.Optional[T]`) |
| `format_mappings` | Custom type mappings                                  |

### TypeScript

| Option               | Description                                           |
| -------------------- | ----------------------------------------------------- |
| `null_optional`      | Use `null` instead of `undefined` for optional fields |
| `branded_primitives` | Use branded types for nominal typing                  |
| `format_mappings`    | Custom type mappings                                  |

### TypeScript Zod

| Option            | Description          |
| ----------------- | -------------------- |
| `format_mappings` | Custom type mappings |

### Java

| Option            | Description                     |
| ----------------- | ------------------------------- |
| `package`         | Package name for generated code |
| `format_mappings` | Custom type mappings            |

### Python

| Option            | Description          |
| ----------------- | -------------------- |
| `format_mappings` | Custom type mappings |

## Format Mappings

Override how JSON Schema formats map to target types:

```yaml
golang:
  format_mappings:
    uuid:
      type: "uuid.UUID"
      import: "github.com/google/uuid"
    date-time:
      type: "time.Time"
      import: "time"

python:
  format_mappings:
    uuid:
      type: "UUID"
      import: "uuid"
    email:
      type: "EmailStr"
      import: "pydantic"
```

## More on Why

While working across many languages that often need to talk to each other over RPC-like transports or perform LLM assisted tool-calls, I often reach for JSON Schema as its a lingua-franca of data structure modelling.

I built this tool primarily to serve the plugin system (which heavily uses JSON-RPC defined in JSON Schema) and the AI Agents features of [Storyden](https://github.com/Southclaws/storyden) which supports plugins written in many languages. So Schemancer was born to easily build SDKs in Go, Python, TypeScript and Java.

This codebase is also somewhat of an experiment of code-generation tooling for LLMs. I've found agentic coding tools really shine when you provide an input and a desired output. As such, **I have absolutely no clue how this code works**, I simply provided the input schemas and desired output in tests, the clanker did the rest. The only input I had was to nudge the agent to use an _intermediate representation_ (`schemancer/ir/ir.go`) much like compilers do, to simplify the architecture. Most of the logic is there then once a schema is in IR format, it's fairly trivial to pipe into a template engine to generate code for each language.

_Using_ generated code is also a brilliant grounding tool for working with agents. I frequently write schemas by hand to build the data model I want then use agentic coding tools to utilise the generated code to achieve the goal I want. I've found this workflow to be very effective in enforcing domain boundaries and invariants.

If you would like a new language or a particular feature, simply provide your example schema and your desired output code, the clanker will do the rest.
