# What's `godoc-lint`?

`godoc-lint` is a (little) opinionated linter for Go documentation practice, also known as *Go Doc* or *godoc*. While `gofmt` handles formatting and correct usage of godocs, still it does not enforce particular restrictions/stylings to godocs. Here is where `godoc-lint` comes into play by adding a set of extra rules to enhance readability, consistency, and developer experience.

## rules

- `max-length`: Limits maximum godoc line length. The default length is 77 characters (not including the `// `, `/*`, or `*/` tokens).

For more configuration details see the [Configuration](#Configuration) section.

## Disabling rules

`godoc-lint` supports inline disable hints to temporarily skip enforcing given set of rules. To do so, one needs to add a comment formatted as below in the symbol declaration godoc:

```go
// Foo is a constant.
//
//godoclint:disable [[RULE] ...]
const Foo = 0
```

Any number of rules can be listed (separated with a single whitespace). If no rule is provided, all rules will be disabled. For example, this will disable `max-length` and `name-prefix` rules for the symbol `Foo`:

```go
// Foo is a function.
//
//godoclint:disable max-length name-prefix
func Foo() {}
```

It is also possible to use multiple `//godoclint:disable` directives:

```go
// Foo is a function.
//
//godoclint:disable max-length
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

To have a customized experience, users can define their configuration in a file named `.godoc-lint.yaml` (or `.godoclint.yaml`). `godoc-lint` looks for this file in the current directory where it is invoked. Alternatively, one can pass the path to the configuration file via the `-config` command line option:

```sh
godoc-lint -config /path/to/config.yaml ./...
```

Check out [`.godoc-lint-example.yaml`](./.godoc-lint.example.yaml) for an example of the configuration file.

## Contributing

`godoc-lint` loves to see developers contributing to it. So, please feel free to submit a [new issue](https://github.com/godoc-lint/godoc-lint/issues/new) for bug report, feature request, or any kind of discussion.
