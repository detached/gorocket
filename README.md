# gorocket
[![Build Status](https://travis-ci.org/detached/gorocket.svg?branch=master)](https://travis-ci.org/detached/gorocket)
[![Coverage Status](https://coveralls.io/repos/github/detached/gorocket/badge.svg?branch=master)](https://coveralls.io/github/detached/gorocket?branch=master)

RocketChat client for golang. Compatible to the rest API of version 0.48.2.

The tests are failing because the library is not fully compatible to the newest version of RocketChat.
I will not update the lib because I am not using RocketChat any more.

RocketChat provides a rest and a realtime interface. This library provides clients for both.

```
go get github.com/detached/gorocket/rest
go get github.com/detached/gorocket/realtime
```

For more information checkout the [rest-godoc](https://godoc.org/github.com/detached/gorocket/rest) and [realtime-godoc](https://godoc.org/github.com/detached/gorocket/realtime), the test files or the examples.
