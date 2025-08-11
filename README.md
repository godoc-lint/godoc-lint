# What's Godoc-Lint?

[![Go Reference](https://pkg.go.dev/badge/github.com/godoc-lint/godoc-lint.svg)](https://pkg.go.dev/github.com/godoc-lint/godoc-lint)
[![CI](https://github.com/godoc-lint/godoc-lint/actions/workflows/ci.yaml/badge.svg)](https://github.com/godoc-lint/godoc-lint/actions/workflows/ci.yaml)

*Godoc-Lint* is a *little* opinionated linter for Go documentation practice, also known as *Go Doc* or *godoc*. Godocs are well explained in this official Golang document, titled [*Go Doc Comments*][godoc-ref].

[godoc-ref]: https://go.dev/doc/comment

While `gofmt` handles formatting and correct usage of godocs, still it does not enforce particular restrictions/stylings to godocs. Here is where Godoc-Lint comes into play by adding a set of extra rules to enhance readability, consistency, and developer experience.

> [!IMPORTANT]
> Godoc-Lint is still under development (`v0.x.x`). So, the Go API (for using the linter as a dependency) is not fully stable until `v1.x.x` has been released. However, the CLI experience is stable enough at this stage.

## Installation

Godoc-Lint binaries are available in the repository's [Releases][releases] page. This is the recommended method to get the stable versions.

However, users can also install Godoc-Lint from source:

```sh
go install github.com/godoc-lint/godoc-lint/cmd/godoclint
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

| Option        | Description                                                          |
| ------------- | -------------------------------------------------------------------- |
| `-enable`     | Comma-separated list of rules to enable (multiple usage allowed)     |
| `-disable`    | Comma-separated list of rules to disable (multiple usage allowed)    |
| `-include`\* | Regexp pattern of relative paths to include (multiple usage allowed) |
| `-exclude`\* | Regexp pattern of relative paths to exclude (multiple usage allowed) |

> [!WARNING]
> **(\*)** The path patterns supplied via `-include` or `-exclude` options should assume Unix-like paths (i.e. separated by forward slashes, `/`). This is to ensure a consistent behavior across different platforms.

## Rules

Below is a brief description of the linter's rules. Some rules are configurable via the `options` key in the configuration file (See [Configuration](#Configuration) for more details).

### `pkg-doc`

Ensures all package godocs start with "Package \<NAME\>":

```go
// This is an example package.  // (Bad)
package foo

// Package foo is an example.   // (Good)
package foo
```

The "Package" word can be configured to any other value via the `pkg-doc/start-with` option. Test files are skipped by default. To enable the rule for them, the `pkg-doc/include-tests` option should be set to `true`.

> [!NOTE]
> As of [*Go Doc Comments*][godoc-cmd-ref], command packages (i.e., packages named `main`) are exceptions to this rule. So, Godoc-Lint ignores them and their test packages (i.e., `main_test`) by default.

[godoc-cmd-ref]: https://go.dev/doc/comment#cmd

### `single-pkg-doc`

Technically, every Go file in a package can have a godoc above the `package` statement. This rule enforces only one godoc, if any, for any package. Test files are skipped by default. To enable the rule for them, the `single-pkg-doc/include-tests` option should be set to `true`.

### `require-pkg-doc`

Ensures that every Go package has godoc(s). By default, test files (i.e., `*_test.go`) and therefore test packages (i.e., `*_test`) are ignored. To include them in the check, the `require-pkg-doc/include-tests` should be set to `true`.

### `start-with-name`

Checks godocs start with the corresponding symbol name:

```go
// This is a constant.  // (Bad)
const foo = 0

// foo is a constant.   // (Good)
const foo = 0
```

It allows English articles (i.e., *a*, *an*, and *the*) at the beginning of godocs. The `start-with-name/pattern` option can be used to customize the starting pattern. If the `start-with-name/pattern` is set to empty, then all godocs have to start with the symbol names. By default, test files are skipped. To enable the rule for test files, the `start-with-name/include-tests` should be set to `true`.

### `require-doc`

Ensures all exported and/or (optionally) unexported symbols have godocs. By default, symbols declared in test files, together with any unexported symbols are ignored. To include test files, the `require-doc/include-tests` option should be set to `true`. Unexported symbols can be included in the check if the `require-doc/ignore-unexported` options is set to `false`. Although it is a rare scenario but one may want to ignore exported symbols, for which the `require-doc/ignore-exported` should be set to `true`.

### `max-len`

Limits maximum line length for godocs. The default length is 77 characters (not including the `// `, `/*`, or `*/` tokens):

```go
// This is a super loooooooooooooooooooooooooooooooooooooooooong godoc.  // (Bad)
const foo = 0

// This is not a super long godoc.   // (Good)
const foo = 0
```

The pre-formatted sections (e.g., codes), or link definitions are ignored.

The maximum line length can be configured via the `max-len/length` option. The rule skips test files by default. To enable it the `max-len/include-tests` option should be set to `true`.

> [!TIP]
> A long hyperlink in the godoc text can break this rule. In such cases, it is best to define the link at the end of the godoc and use the reference in the text:
>
> ```go
> // Check this [link].
> //
> // [link]: https://foo.com/super/loooooooooooooooooooooooooooooooooooooooong/link
> const foo = 0
> ```

### `no-unused-link`

Checks for unused links in the godoc text:

```go
// This is a godoc with an unused link.  // (Bad)
//
// [docs]: https://foo.com/docs
const foo = 0

// Check [docs] here.                    // (Good)
//
// [docs]: https://foo.com/docs
const foo = 0
```

The rule skips test files by default. To include them, the `no-unused-link/include-tests` option should be set to `true`.

## Disabling rules

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
