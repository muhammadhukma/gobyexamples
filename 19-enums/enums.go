package main

import "fmt"

/*
Enumerated types (enums) are a special case of sum type.
An enum is a type that has a fixed number of possible value, each with a distinct name.
Go doesn't have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.
*/

/*
Our enum type ServerState has an underlying int type.
*/
type ServerState int

/*
The possible values for ServerState are defined as constant. The special keyword iota generates succesive constant values automatically;
in this case 0,1,2 and so on.
*/
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

/*
by implementing the fmt.Stringer interface, values of ServerState can be printed out or converted to strings.

This can get cumbersome if there are many possible value.
In such cases the stringer tool can be used in conjunction with go:generate to automate the process.
See this post for a longer explanation.
*/
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	fmt.Println(StateIdle)
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
