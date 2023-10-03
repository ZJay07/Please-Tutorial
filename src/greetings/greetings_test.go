package greetings

import "testing"

func TestGreeting(t *testing.T) {
    if Greeting() == "" {
        panic("Greeting failed to produce a result")
    }
}
