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