# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

schemancer (JSON Schema Code Generator) generates type-safe code from JSON Schema definitions. It supports multiple target languages: Go, TypeScript, TypeScript-Zod, Java, and Python.

## Commands

```bash
# Run all tests
go test ./...

# Run tests for a specific language
go test ./tests/golang/...
go test ./tests/typescript/...
go test ./tests/java/...
go test ./tests/python/...

# Run a single test
go test ./tests/golang/simple_object/... -v

# Generate code (single language to output directory)
go run . schema.yaml golang ./output
go run . schema.yaml typescript ./output
go run . schema.yaml typescript-zod ./output
go run . schema.yaml java ./output
go run . schema.yaml python ./output

# Generate code (all configured languages from schemancer.yaml)
go run . schema.yaml

# Output to stdout
go run . schema.yaml golang -
```

## Architecture

The codebase follows a pipeline: **Schema → IR → Generated Code**

### Pipeline Flow

1. **Loader** (`schemancer/loader/`) - Reads YAML/JSON schema files
2. **Deref** (`schemancer/deref/`) - Resolves external `$ref` references
3. **Schema to IR** (`schemancer/schema_to_ir.go`) - Converts JSON Schema to language-agnostic IR
4. **Detect** (`schemancer/detect/`) - Identifies patterns like discriminated unions
5. **Generators** (`schemancer/generators/<lang>/`) - Emit target language code from IR

### Intermediate Representation (IR)

The IR (`schemancer/ir/ir.go`) is the core abstraction that decouples schema parsing from code generation:

- `IRType` - Represents a type with kind (struct, alias, enum, discriminated_union)
- `IRField` - A field within a struct
- `IRTypeRef` - Reference to a type (builtin, named, array, map)
- `IRFormat` - JSON Schema formats (uuid, date-time, email, etc.)
- `IRDiscriminatedUnion` - Union type with discriminator field and variants

### Generator Structure

Each generator in `schemancer/generators/<lang>/generator.go`:

- Implements `generators.Generator` interface
- Has `DefaultFormatMappings` for type conversions
- Uses Go `text/template` for code generation
- Supports language-specific options via the Option pattern

### Configuration

`cli/config/config.go` defines the `schemancer.yaml` structure with per-language configs:

- `output` - Output directory path (enables multi-language generation when set)
- `package` - Package/module name (Go, Java)
- `format_mappings` - Custom type mappings for JSON Schema formats
- Language-specific options (e.g., `optional_style` for Go, `branded_primitives` for TypeScript)

**Note:** The `output` field specifies a directory, not a file. Generated files will be created inside this directory with appropriate names determined by the generator.

## Test Structure

Tests are in `tests/<language>/<feature>/`:

- `schema.yaml` - Input JSON Schema
- `expected.py` / `expected.ts` / `expected.java` / `expected_test.go` - Expected output
- `generate_test.go` - Test that generates and compares output
- `output.py` / `output.ts` / `output.java` / `output.go` - Actual generated output (for comparison)

The `tests/multi/` directory tests multi-language generation from a single schema using a `schemancer.yaml` config file.

## Adding a New Generator

1. Create `schemancer/generators/<lang>/generator.go` implementing `generators.Generator`
2. Add language constant to `schemancer/generators/generators.go`
3. Register in `schemancer/generate.go` Generators map
4. Add CLI support in `main.go` (case in `getGeneratorOptions`)
5. Add config struct in `cli/config/config.go`
6. Create test suite in `tests/<lang>/`
7. Update the multi-test in `tests/multi/` this test generates all languages from a single schema and config file.
