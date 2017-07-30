[![Build Status](https://travis-ci.org/scottjustin5000/dot.env.svg?branch=master)](https://travis-ci.org/scottjustin5000/dot.env)
# dotenv

load .env variables into environment.

### Usage

```go
import (
  "github.com/scottjustin5000/dotenv"
)


func main() {

  err := dotenv.Load()
  if err != nil {
    //error
  }
  ...
}

```
