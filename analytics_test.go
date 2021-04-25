package analytics

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	SetKeys(os.Getenv("SECRET"), os.Getenv("MEASUREMENT"))
	tt := time.Now().Unix()
	p := Payload{
		ClientID: fmt.Sprintf("%d.%d", rand.Int31(), tt),
		UserID:   "cedbd94e-fafa-4195-a99b-49ea0943a0a2",
		Events: []Event{
			{
				Name:   "tutorial_begin",
				Params: map[string]interface{}{},
			},
		},
	}
	//Debug(true)
	Send(p)
}
