### google-analytics
>Implement Measurement Protocol (Google Analytics 4)

### USAGE
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/ystyle/google-analytics"
)

func main() {
	analytics.SetKeys("goxxxxx-xxxxxxxxxxx", "G-xxxxxxx") // // required
	tt := time.Now().Unix()
	payload := analytics.Payload{
		ClientID: fmt.Sprintf("%d.%d", rand.Int31(), tt), // required
		Events: []analytics.Event{
			{
				Name: "tutorial_begin", // required
				Params: map[string]interface{}{
					
                },
			},
		},
	}
	// analytics.Debug(true)
	analytics.Send(payload)
}
```