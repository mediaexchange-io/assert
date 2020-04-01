[![Build Status](https://travis-ci.org/mediaexchange/assert.svg)](https://travis-ci.org/mediaexchange/assert)
[![GoDoc](https://godoc.org/github.com/mediaexchange/assert/github?status.svg)](https://godoc.org/github.com/mediaexchange/assert)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go version](https://img.shields.io/badge/go-~%3E1.12-green.svg)](https://golang.org/doc/devel/release.html#go1.12)
[![Go version](https://img.shields.io/badge/go-~%3E1.13-green.svg)](https://golang.org/doc/devel/release.html#go1.13)
[![Go version](https://img.shields.io/badge/go-~%3E1.14-green.svg)](https://golang.org/doc/devel/release.html#go1.14)

# assert

Lightweight assertion library based on the fluent interface from
[assertj](http://joel-costigliola.github.io/assertj/)

## Features

The matchers included in our `assert` library are fully compatible with, and
depend on the standard Go [testing package](https://golang.org/pkg/testing/).
These just add a little syntactic sugar on top of the familiar test patterns.

To use the example from the testing documentation, here is how one would
normally write a test in Go:

```go
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

With the matchers included in our `assert` package, one would write:

```go
import "github.com/MediaExchange/assert"

func TestAbs(t *testing.T) {
    got := Abs(-1)
    assert.With(t).
        That(got).
        IsEqualTo(1)
}
```

This is much more readable and ultimately leads to more maintainable code.

## Usage

The matchers currently included in the `assert` package are:

1. IsEmpty/IsNotEmpty

    ```go
    func TestIsEmpty(t *testing.T) {
        s := ""
        assert.With(t).
            That(s).
            IsEmpty()
    }
    ```
    
    ```go
    func TestIsNotEmpty(t *testing.T) {
        s := "foobar"
        assert.With(t).
            That(s).
            IsNotEmpty()
    }
    ```

1. IsEqualTo

    ```go
    func TestEquals(t *testing.T) {
        got := Abs(-1)
        assert.With(t).
            That(got).
            IsEqualTo(1)
    }
    ```

1. IsGreaterThan

    ```go
    func TestIsEmpty(t *testing.T) {
        x := 1
        assert.With(t).
            That(x).
            IsGreaterThan(0)
    }
    ```

1. IsNil/IsNotNil

    ```go
    func TestIsNil(t *testing.T) {
        var s *string
        assert.With(t).
            That(s).
            IsNil()
    }
    ```

    ```go
    func TestIsNotNil(t *testing.T) {
        var s string
        assert.With(t).
            That(s).
            IsNotNil()
    }
    ```

1. IsOk

    ```go
    func TestIsOk(t *testing.T) {
        f, err := io.Open("filename.ext")
        assert.With(t).
            That(err).
            IsOk()
    }
    ```

1. ThatPanics

    ```go
    func TestThatPanics(t *testing.T) {
        f := func() {
            panic("error")
        }
        assert.With(t).
            ThatPanics(f)
    }
    ```

## Contributing

 1.  Fork it
 2.  Create a feature branch (`git checkout -b new-feature`)
 3.  Commit changes (`git commit -am "Added new feature xyz"`)
 4.  Push the branch (`git push origin new-feature`)
 5.  Create a new pull request.

## Maintainers

* [Media Exchange](http://github.com/MediaExchange/)

## License

Copyright 2019 MediaExchange.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
