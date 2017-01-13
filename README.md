# gorocket
[![Build Status](https://travis-ci.org/detached/gorocket.svg?branch=master)](https://travis-ci.org/detached/gorocket)
[![Coverage Status](https://coveralls.io/repos/github/detached/gorocket/badge.svg?branch=master)](https://coveralls.io/github/detached/gorocket?branch=master)

RocketChat client for golang

This library makes use of [gopackage/ddp](https://github.com/gopackage/ddp) to register users.

# Usage

Fetch source
```
go get github.com/detached/gorocket
```

Use in your code
```
import "github.com/detached/gorocket"
```

###Initialize
```
rocket := Rocket{Protocol: "http", Host: "your-server.com", Port: "3000"}
```

###Register new user
```
rocket.RegisterUser(UserCredentials{Name:"userName", Email:"user@domain.com", Password:"userPassword"})
```

###Login user
```
rocket.Login(UserCredentials{Name:"userName", Email:"user@domain.com", Password:"userPassword"})
```

###Get public channels
```
channels, err := rocket.GetPublicChannels()
```

###Send message
```
rocket.Send(channel, "Text")
```

###Get messages
```
messages, err := rocket.GetMessages(channel, &Page{Count: 20})
```
or without pagination

```
messages, err := rocket.GetMessages(channel, nil)
```

For more information checkout the [godoc](https://godoc.org/github.com/detached/gorocket) and the test files.