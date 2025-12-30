# What's Godoc-Lint?

[![Go Reference](https://pkg.go.dev/badge/github.com/godoc-lint/godoc-lint.svg)](https://pkg.go.dev/github.com/godoc-lint/godoc-lint)
[![Release](https://img.shields.io/github/v/release/godoc-lint/godoc-lint)](https://github.com/godoc-lint/godoc-lint/releases)
[![License](https://img.shields.io/github/license/godoc-lint/godoc-lint)](/LICENSE)
[![CI](https://github.com/godoc-lint/godoc-lint/actions/workflows/ci.yaml/badge.svg)](https://github.com/godoc-lint/godoc-lint/actions/workflows/ci.yaml)

*Godoc-Lint* is a fast, *little* opinionated linter for Go documentation practice, also known as *Go Doc* or *godoc* (See [*Go Doc Comments*][godoc-ref]), ready to be used **out of the box** without further configuration. It is highly recommended to be used when developing **reusable Go modules, like SDKs, API clients, or special-purpose libraries,** which need consistent/standard developer experience in IDEs, as well as on [pkg.go.dev](https://pkg.go.dev) docs.

[godoc-ref]: https://go.dev/doc/comment

## Using via Golangci-lint

Godoc-Lint is now available as part of [Golangci-lint][golangci-lint] suite of linters (since `v2.5.0`). To enable the linter, it should be added to the `.golangci.yml` file:

```yaml
version: "2"
linters:
  enable:
    - godoclint
```

When used via Golangci-lint, the linter's configuration will be different from what is covered in this document. Users should consult Golangci-lint [docs][golangci-lint-config] for more details.

> [!TIP]
> In most cases, it is best to exclude Go test files (i.e., `*_test.go`) from the analysis. This can be done by adding the following lines to `.golangci.yml`:
>
> ```yaml
> linters:
>   exclusions:
>     rules:
>       - path: _test\.go$
>         linters:
>           - godoclint
> ```
>
> More about exclusion methods is available [here][golangci-lint-fp].


[golangci-lint]: https://golangci-lint.run
[golangci-lint-config]: https://golangci-lint.run/docs/linters/configuration/#godoclint
[golangci-lint-fp]: https://golangci-lint.run/docs/linters/false-positives/

## Installation

Godoc-Lint binaries are available in the repository's [Releases][releases] page.

Users can also install Godoc-Lint binary from source code by using any of these commands:

```sh
# Latest version
go install github.com/godoc-lint/godoc-lint/cmd/godoclint@latest

# Specific version
go install github.com/godoc-lint/godoc-lint/cmd/godoclint@v0.10.0
```

Additionally, the linter can be run from source code via the following command:

```sh
# Latest version
go run github.com/godoc-lint/godoc-lint/cmd/godoclint@latest ./...

# Specific version
go run github.com/godoc-lint/godoc-lint/cmd/godoclint@v0.10.0 ./...
```

[releases]: https://github.com/godoc-lint/godoc-lint/releases

## Usage

Users can simply run the `godoclint` CLI at the root directory of their Go source code:

```sh
godoclint ./...
```

This will run the linter with its default configuration on all Go packages. It is also possible to narrow the search path:

```sh
godoclint ./internal/foo/bar # Exact package, no sub-packages
godoclint ./internal/...     # All sub-packages
```

Godoc-Lint looks for `.godoc-lint.yaml` file in the working directory for configuration (Check the [Configuration](#Configuration) section for more details). If not found, the linter will use sensible defaults.

Although it is best to set the configuration parameters in a file, there are a number of CLI options to modify linter parameters:

| Option       | Description                                                               |
| ------------ | ------------------------------------------------------------------------- |
| `-default`   | Default set of rules to enable, one of `basic` (default), `all` or `none` |
| `-enable`    | Comma-separated list of rules to *also* enable (multiple usage allowed)   |
| `-disable`   | Comma-separated list of rules to disable (multiple usage allowed)         |
| `-include`\* | Regexp pattern of relative paths to include (multiple usage allowed)      |
| `-exclude`\* | Regexp pattern of relative paths to exclude (multiple usage allowed)      |

> [!WARNING]
> **(\*)** The path patterns supplied via `-include` or `-exclude` options should assume Unix-like paths (i.e., separated by forward slashes, `/`). This is to ensure a consistent behavior across different platforms.

## Rules

The linter provides a number of rules that can be categorized as in this table:

| Category          | Rules                                                                                  | Notes                                                              |
| ----------------- |----------------------------------------------------------------------------------------| ------------------------------------------------------------------ |
| Basic *(default)* | `pkg-doc` </br> `single-pkg-doc` </br> `start-with-name` </br> `deprecated`            | Recommended by [*Go Doc Comments*][godoc-ref], and **low-effort**  |
| Strict            | `require-doc` </br> `require-pkg-doc`                                                  | Recommended by [*Go Doc Comments*][godoc-ref], and **high-effort** |
| Extra             | `max-len` </br> `no-unused-link` </br> `require-stdlib-doclink`                        | Extra but compatible with [*Go Doc Comments*][godoc-ref]           |

**Rules under the *Basic* category are enabled by default** and do not need further configuration, unless, of course, one wants to tune their parameters. The rest has to be explicitly enabled via configuration.

Below is a brief description of the linter's rules. Some rules are configurable via the `options` key in the configuration file (See [Configuration](#Configuration) for more details).

### `pkg-doc`

> Since `v0.1.0`, Golangci-lint `v2.5.0`.

Ensures all package godocs start with "Package \<NAME\>":

```go
// This is an example package.  // (Bad)
package foo

// Package foo is an example.   // (Good)
package foo
```

Test files are skipped by default. To enable the rule for them, the `pkg-doc/include-tests` option should be set to `true`.

> [!NOTE]
> As of [*Go Doc Comments*][godoc-cmd-ref], command packages (i.e., packages named `main`) are exceptions to this rule. So, Godoc-Lint ignores them and their test packages (i.e., `main_test`) by default.

[godoc-cmd-ref]: https://go.dev/doc/comment#cmd

### `single-pkg-doc`

> Since `v0.1.0`, Golangci-lint `v2.5.0`.

Technically, every Go file in a package can have a godoc above the `package` statement. This rule enforces only one godoc, if any, for any package. Test files are skipped by default. To enable the rule for them, the `single-pkg-doc/include-tests` option should be set to `true`.

### `require-pkg-doc`

> Since `v0.1.0`, Golangci-lint `v2.5.0`.

Ensures that every Go package has godoc(s). By default, test files (i.e., `*_test.go`) and therefore test packages (i.e., `*_test`) are ignored. To include them in the check, the `require-pkg-doc/include-tests` should be set to `true`.

### `start-with-name`

> Since `v0.1.0`, Golangci-lint `v2.5.0`.

Checks godocs start with the corresponding symbol name:

```go
// This is a constant.  // (Bad)
const Foo = 0

// Foo is a constant.   // (Good)
const Foo = 0
```

It allows English articles (i.e., *a*, *an*, and *the*) at the beginning of godocs.

By default, unexported symbols are skipped. To include them the `start-with-name/include-unexported` option should be set to `true`. Test files are also skipped. To enable the rule for test files, the `start-with-name/include-tests` option should be set to `true`.

### `require-doc`

> Since `v0.1.0`, Golangci-lint `v2.5.0`.

Ensures all exported and/or (optionally) unexported symbols have godocs. By default, symbols declared in test files, together with any unexported symbols are ignored. To include test files, the `require-doc/include-tests` option should be set to `true`. Unexported symbols can be included in the check if the `require-doc/ignore-unexported` options is set to `false`. Although it is a rare scenario but one may want to ignore exported symbols, for which the `require-doc/ignore-exported` should be set to `true`.

### `deprecated`

> Since `v0.9.0`, Golangci-lint `v2.5.0`.

Checks if deprecation notes are formatted correctly. This rule only applies to exported symbols.

```go
// Foo is a symbol.
//
// DEPRECATED: do not use  // (Bad)
const Foo = 0

// Foo is a symbol.
//
// Deprecated: do not use  // (Good)
const Foo = 0
```

### `max-len`

> Since `v0.1.0`, Golangci-lint `v2.5.0`.

Limits maximum line length for godocs. The default length is 77 characters (not including the `// `, `/*`, or `*/` tokens):

```go
// Foo has a super loooooooooooooooooooooooooooooooooooooooooooooooooooong godoc.  // (Bad)
const Foo = 0

// Foo has a reasonably long godoc.  // (Good)
const Foo = 0
```

The pre-formatted sections (e.g., codes), or link definitions are ignored.

The maximum line length can be configured via the `max-len/length` option. The rule skips test files by default. To enable it the `max-len/include-tests` option should be set to `true`.

Specific long lines (for example, ones matching known patterns) can be excluded from this rule by listing regexp patterns under the `max-len/ignore-patterns` option; any rendered godoc line matching at least one of these patterns is not checked for length. Note that, when using Golangci-lint, pattern-based exclusions are available via [`source` text matching](https://golangci-lint.run/docs/linters/false-positives/#exclude-issue-by-text).

> [!TIP]
> A long hyperlink in the godoc text can break this rule. In such cases, it is best to define the link at the end of the godoc and use the reference in the text:
>
> ```go
> // Foo is a const. Check this [link].
> //
> // [link]: https://foo.com/super/loooooooooooooooooooooooooooooooooooooooong/link
> const Foo = 0
> ```

### `no-unused-link`

> Since `v0.2.0`, Golangci-lint `v2.5.0`.

Checks for unused links in the godoc text:

```go
// Foo godoc has an unused link.     // (Bad)
//
// [link]: https://foo.com/docs
const Foo = 0

// Foo godoc uses a defined [link].  // (Good)
//
// [link]: https://foo.com/docs
const Foo = 0
```

The rule skips test files by default. To include them, the `no-unused-link/include-tests` option should be set to `true`.

### `require-stdlib-doclink`

> Since `v0.11.0`.

Suggests turning plain-text mentions of standard-library identifiers into [*doc links*](https://go.dev/doc/comment#doclinks), when possible. For example, the text `encoding/json.Encoder` in a godoc can be turned into a doc link like `[encoding/json.Encoder]` so that it links to the corresponding stdlib symbol on [`pkg.go.dev`](https://pkg.go.dev/encoding/json#Encoder). To avoid false positives, the linter only detects potential doc links of longer forms like `pkg.name` or `pkg.recv.name`.

```go
// Println is the same as fmt.Println.    // (Bad)
func Println(a ...any) (n int, err error) {}

// Println is the same as [fmt.Println].  // (Good)
func Println(a ...any) (n int, err error) {}
```

The rule skips test files by default. To include them, the `require-stdlib-doclink/include-tests` option should be set to `true`.

## Disabling rules

> [!TIP]
> Users who run the linter via Golangci-lint can also use the `//nolint:godoclint` directive to disable the linter. The `//nolint` directive usage is explained in the Golangci-lint's official [docs][golangci-nolint].

[golangci-nolint]: https://golangci-lint.run/docs/linters/false-positives/#nolint-directive

Godoc-Lint supports inline directives to temporarily skip enforcing given set of rules. The directive must be formatted as:

```go
//godoclint:disable [[RULE] ...]
```

> [!NOTE]
> There must be no whitespace between `//` and `godoclint:disable`.

For example, this will temporarily disable the `start-with-name` rule for the `Foo` symbol's godoc:

```go
// This is a constant.
//
//godoclint:disable start-with-name
const Foo = 0
```

Any number of rules can be listed, separated with whitespaces. If no rule is provided, all rules will be disabled. For example, this will disable `start-with-name` and `max-len` rules for the `Foo` symbol's godoc:

```go
// This is a function.
//
//godoclint:disable start-with-name max-len
func Foo() {}
```

It is also possible to use multiple `//godoclint:disable` directives:

```go
// This is a function.
//
//godoclint:disable start-with-name
//godoclint:disable max-len
func Foo() {}
```

There are cases where one would want to disable all linter rules for a specific declaration. This can be done by just omitting the rule names in the directive:

```go
// This is a function.
//
//godoclint:disable
func Foo() {}
```

Rules can be disabled for an entire file. To do this, the `//godoclint:disable` comment should be added at any position at the top level, in a *non-godoc* comment group. For instance, this will disable all rules for the entire file:

```go
package foo

//godoclint:disable
```

Sometimes, it is not possible/preferred to add the inline `//godoclint:disable` directives to a file (e.g., an auto-generated file, or a legacy file that should not be altered). In such cases, the configuration file is the right place to instruct the linter. All one needs to do is to add the files under the `exclude` key. More about this in the [Configuration](#Configuration) file section.

## Configuration

To have a customized experience, users can define their configuration in a file named `.godoc-lint.yaml` (or `.godoclint.yaml`). The linter looks for this file in the working directory where it is invoked. However, users can specify a different file name via the `-config` option:

```sh
godoclint -config the-config-file.yaml ./...
```

Godoc-Lint comes with a sensible default configuration that will be used when there is no configuration file. Check out [`.godoc-lint.default.yaml`](./.godoc-lint.default.yaml) for more details.

### Overriding configuration

> [!WARNING]
> When used via Golangci-lint, the linter's configuration will be different from what is covered here. Users should consult Golangci-lint [docs][golangci-lint-config] for more details.

In addition to the root directory, Godoc-Lint allows configuration files in sub-directories. When a package is being processed, the configuration file in the package's directory, if any, will be used by the linter. If there is no such file, the linter looks for it in the directory's parents, up until the root where the linter was invoked.

For example, in the file tree below package `foo` gets processed with the configuration expressed in `foo/.godoc-lint.yaml`.

```text
├─ foo
│  ├─ .godoc-lint.yaml
|  ├─ foo.go
├─ .godoc-lint.yaml
├─ main.go
```

## Contributing

Godoc-Lint loves to see developers contributing to it. So, please feel free to submit a [new issue](https://github.com/godoc-lint/godoc-lint/issues/new) for bug report, feature request, or any kind of discussion.

## Links

- [Go Doc Comments](https://go.dev/doc/comment)
- [Golangci-lint configuration](https://golangci-lint.run/docs/linters/configuration/#godoclint)
