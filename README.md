# What's `godoc-lint`?

`godoc-lint` is a (little) opinionated linter for Go documentation practice, also known as *Go Doc* or *godoc*. While `gofmt` handles formatting and correct usage of godocs, still it does not enforce particular restrictions/stylings to godocs. Here is where `godoc-lint` comes into play by adding a set of extra rules to enhance readability, consistency, and developer experience.

## Rules

Below is a brief description of the linter's rules. Some rules are configurable via the `options` key in the configuration file (See [Configuration](#Configuration) for more details).

### `max-len`

Limits maximum line length for godocs. The default length is 77 characters (not including the `// `, `/*`, or `*/` tokens). The maximum line length can be configured via the `max-len` option.

### `pkg-doc`

Ensures all package godocs start with "Package \[NAME\]". The "Package" word can be configured to any other value via the `pkg-doc/start-with` option.

### `single-pkg-doc`

Technically, every Go file in a package can have a godoc above the `package` statement. This rule enforces only one godoc, if any, for any package.

## Disabling rules

`godoc-lint` supports inline disable hints to temporarily skip enforcing given set of rules. To do so, one needs to add a comment formatted as below in the symbol declaration godoc:

```go
// Foo is a constant.
//
//godoclint:disable [[RULE] ...]
const Foo = 0
```

Any number of rules can be listed (separated with a single whitespace). If no rule is provided, all rules will be disabled. For example, this will disable `max-len` and `name-prefix` rules for the symbol `Foo`:

```go
// Foo is a function.
//
//godoclint:disable max-len name-prefix
func Foo() {}
```

It is also possible to use multiple `//godoclint:disable` directives:

```go
// Foo is a function.
//
//godoclint:disable max-len
//godoclint:disable name-prefix
func Foo() {}
```

Rules can be disabled for an entire file. To do this, the `//godoclint:disable` comment should be added at any position at the root-level of the file, in a *non-godoc* comment group. For example, this will disable all rules for the entire file:

```go
package foo

//godoclint:disable
```

Sometimes, it is not possible/preferred to add the inline `//godoclint:disable` directives to a file (e.g., an auto-generated file, or a legacy file that should not be altered in any way). In such cases, the configuration file is the right place to instruct `godoc-lint`. All one needs to do is to add the files under the `disable` key. More about this in the [Configuration](#Configuration) file section.

## Configuration

To have a customized experience, users can define their configuration in a file named `.godoc-lint.yaml` (or `.godoclint.yaml`). The linter looks for this file in the current directory where it is invoked. Alternatively, one can pass the path to the configuration file via the `-config` command line option:

```sh
godoclint -config /path/to/config.yaml ./...
```

Check out [`.godoc-lint.example.yaml`](./.godoc-lint.example.yaml) for an example of the configuration file.

## Contributing

`godoc-lint` loves to see developers contributing to it. So, please feel free to submit a [new issue](https://github.com/godoc-lint/godoc-lint/issues/new) for bug report, feature request, or any kind of discussion.
