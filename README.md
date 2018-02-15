# Experiments with Bazel and Go

This repo doesn't provide any meaningful feature.
It's just a playground to experiment with some interactions between Bazel, rules_go, gazelle...

## Play with binary stamping

Conditional binary stamping is provided via the `release` config.

Without stamping:

```
$ bazel run testapp
...
time="2018-02-15T11:50:09+01:00" level=info msg="Exposed value is DEFAULT" test=bazel
```

With stamping:
```
$ bazel run --config=release testapp
...
time="2018-02-15T11:50:21+01:00" level=info msg="Exposed value is 100" test=bazel
```
