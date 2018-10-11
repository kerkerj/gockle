# gockle

[![Documentation](https://godoc.org/github.com/kerkerj/gockle?status.svg)](https://godoc.org/github.com/kerkerj/gockle)
[![Build](https://travis-ci.org/kerkerj/gockle.svg?branch=master)](https://travis-ci.org/kerkerj/gockle)
[![Report](https://goreportcard.com/badge/github.com/kerkerj/gockle)](https://goreportcard.com/report/github.com/kerkerj/gockle)
[![Test Coverage](https://coveralls.io/repos/github/kerkerj/gockle/badge.svg?branch=master)](https://coveralls.io/github/kerkerj/gockle?branch=master)

forked from https://github.com/willfaught/gockle

*Note: Test coverage is low because there is no Cassandra database for the tests to use. Providing one yields 97.37% test coverage. Some code is uncovered because gocql cannot be mocked. This is one difficulty your code avoids by using gockle.*

## Differences

1. Query Interface (commit [40e207](https://github.com/kerkerj/gockle/commit/40e20799a09a5bf8df8f89b4fc70492a4d3d136b) is from [rtfb@25a7eea](https://github.com/rtfb/gockle/commit/25a7eea56ca2b26ec9e36dc0a89c7283e10179a5) thanks!)

    1. And `Consistency` method in Query Interface

2. Iterator Interface

    1. Add `SliceMap` method

3. change to use [testify/mock](https://github.com/stretchr/testify)
4. remove tests about mock structs (such as BatchMock, QueryMock...), I think there's no need to test mock file.

## TODO

- [ ] Enhance test coverage

