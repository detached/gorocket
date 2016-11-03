# gorocket
[![Build Status](https://travis-ci.org/detached/gorocket.svg?branch=master)](https://travis-ci.org/detached/gorocket)

RocketChat client for golang

This library makes use of [github.com/gopackage/ddp](https://github.com/gopackage/ddp) to register users.

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

###Get public rooms
```
rooms, err := rocket.GetPublicRooms()
```

###Send message
```
rocket.Send(room, "Text")
```

###Get messages
```
messages, err := rocket.GetMessages(room, &Page{Skip: 10, Limit: 20})
```
