package main

import (
	"fmt"

	"github.com/olegbannyi/result"
)

type user struct {
	id     int
	name   string
	active bool
}

func main() {
	u := findUser(1).UnwrapOrElse(func() user {
		return user{
			id:     -1,
			name:   "unknown",
			active: false,
		}
	})

	fmt.Printf("findUser(1) => %#v\n", u) //findUser(1) => main.user{id:1, name:"john.smith", active:true}

	u = findUser(0).UnwrapOrElse(func() user {
		return user{
			id:     -1,
			name:   "unknown",
			active: false,
		}
	})

	fmt.Printf("findUser(0) => %#v\n", u) //findUser(0) => main.user{id:-1, name:"unknown", active:false}
}

func findUser(id int) result.Result[user] {
	if id <= 0 {
		return result.Err[user](fmt.Errorf("user not found, id: %d", id))
	}

	return result.Ok(user{
		id:     id,
		name:   "john.smith",
		active: true,
	})
}
