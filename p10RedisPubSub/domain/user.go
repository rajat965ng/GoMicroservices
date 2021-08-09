package domain

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (user *User) Marshal() []byte {
	buff, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshalling user: %s\n", err.Error())
	}
	return buff
}

func (user *User) UnMarshal(buff []byte) {
	err := json.Unmarshal(buff, user)
	if err != nil {
		fmt.Printf("Error UnMarshalling user: %s\n", err.Error())
	}
}

func (u *User) String() string {
	return fmt.Sprintf("Use: %s, Age: %d", u.Name, u.Age)
}
