## Go Leaderboard

A redis backed leaderboard handler written in go.

## Installation

```
$ go get github.com/tehsis/leaderboard
```

### Example

```go
package main;

import (
  "fmt"

	"github.com/tehsis/leaderboard"
	redis "gopkg.in/redis.v5"
) 

func main() {
  client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

  scores := leaderboard.NewRedisLeaderBoard("Space Invaders", client)

  scores.Set("Tehsis", 100)
  scores.Set("0xBunny", 10)
  scores.Set("Plaurino", 50)  

  top2 := scores.GetTop(2)

  fmt.Println("THE BEST TWO PLAYERS!")
  for index, score := range top2 {
    fmt.Printf("%v - %v : %v\n", index+1, score.Username, score.Points)
  }

  currentUser := "0xBunny"
  currentUserScore, currentUserPosition := scores.Get(currentUser)

  fmt.Printf("User %v is in position %v with %v points\n", currentUser, currentUserScore, currentUserPosition)
}
```

The previous snippet will produce the following output:

![snippet output](http://i.imgur.com/dhvEWlR.png)
