package entity

import (
	"encoding/json"
	"fmt"
)

// ToString is a method that converts a MessageResponse struct to readable JSON string format
func (response *MessageResponse) ToString() string {
	output, err := json.Marshal(response)
	if err != nil {
		return fmt.Sprint(response)
	}

	return string(output)
}

// ToString is a method that converts a ChatMemberResponse struct to readable JSON string format
func (response *ChatMemberResponse) ToString() string {
	output, err := json.Marshal(response)
	if err != nil {
		return fmt.Sprint(response)
	}

	return string(output)
}

// ToString is a method that converts a ChatMembersResponse struct to readable JSON string format
func (response *ChatMembersResponse) ToString() string {
	output, err := json.Marshal(response)
	if err != nil {
		return fmt.Sprint(response)
	}

	return string(output)
}
