# What's Godoc-Lint?

*Godoc-Lint* is a (little) opinionated linter for Go documentation practice, also known as *Go Doc* or *godoc*. While `gofmt` handles formatting and correct usage of godocs, still it does not enforce particular restrictions/stylings to godocs. Here is where Godoc-Lint comes into play by adding a set of extra rules to enhance readability, consistency, and developer experience.

## Rules

Below is a brief description of the linter's rules. Some rules are configurable via the `options` key in the configuration file (See [Configuration](#Configuration) for more details).

### `max-len`

Limits maximum line length for godocs. The default length is 77 characters (not including the `// `, `/*`, or `*/` tokens). The maximum line length can be configured via the `max-len/length` option. The rule skips test files by default. To enable it the `max-len/include-tests` option should be set to `true`.

### `pkg-doc`

Ensures all package godocs start with "Package \[NAME\]". The "Package" word can be configured to any other value via the `pkg-doc/start-with` option. Test files are skipped by default. To enable the rule for them, the `pkg-doc/include-tests` option should be set to `true`.

### `single-pkg-doc`

Technically, every Go file in a package can have a godoc above the `package` statement. This rule enforces only one godoc, if any, for any package. Test files are skipped by default. To enable the rule for them, the `single-pkg-doc/include-tests` option should be set to `true`.

### `require-pkg-doc`

Ensures that every Go package has godoc(s). By default, test files (i.e., `*_test.go`) and therefore test packages (i.e., `*_test`) are ignored. To include them in the check, the `require-pkg-doc/include-tests` should be set to `true`.

### `require-doc`

Ensures all exported and/or (optionally) unexported symbols have godocs. By default, symbols declared in test files, together with any unexported symbols are ignored. To include test files, the `require-doc/include-tests` option should be set to `true`. Unexported symbols can be included in the check if the `require-doc/ignore-unexported` options is set to `false`. Although it is a rare scenario but one may want to ignore exported symbols, for which the `require-doc/ignore-exported` should be set to `true`.

### `start-with-name`

Checks godocs start with the corresponding symbol name. It allows English articles (i.e., *a*, *an*, and *the*) at the beginning of godocs. The `start-with-name/pattern` option can be used to customize the starting pattern. If the `start-with-name/pattern` is set to empty, then all godocs have to start with the symbol names. By default, test files are skipped. To enable the rule for test files, the `start-with-name/include-tests` should be set to `true`.

## Disabling rules

Godoc-Lint supports inline directives to temporarily skip enforcing given set of rules. The directive must be formatted as:

```go
//godoclint:disable [[RULE] ...]
```

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

Rules can be disabled for an entire file. To do this, the `//godoclint:disable` comment should be added at any position at the root-level of the file, in a *non-godoc* comment group. For instance, this will disable all rules for the entire file:

```go
package foo

//godoclint:disable
```

Sometimes, it is not possible/preferred to add the inline `//godoclint:disable` directives to a file (e.g., an auto-generated file, or a legacy file that should not be altered in any way). In such cases, the configuration file is the right place to instruct the linter. All one needs to do is to add the files under the `exclude` key. More about this in the [Configuration](#Configuration) file section.

## Configuration

To have a customized experience, users can define their configuration in a file named `.godoc-lint.yaml` (or `.godoclint.yaml`). The linter looks for this file in the current directory where it is invoked. Alternatively, one can pass the path to the configuration file via the `-config` command line option:

```sh
godoclint -config /path/to/config.yaml ./...
```

Check out [`.godoc-lint.example.yaml`](./.godoc-lint.example.yaml) for an example of the configuration file.

## Contributing

Godoc-Lint loves to see developers contributing to it. So, please feel free to submit a [new issue](https://github.com/godoc-lint/godoc-lint/issues/new) for bug report, feature request, or any kind of discussion.
