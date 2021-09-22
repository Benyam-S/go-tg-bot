package entity

import (
	"encoding/json"
	"fmt"
)

// ToString is a method that converts a BotResponse struct to readable JSON string format
func (response *BotResponse) ToString() string {
	output, err := json.Marshal(response)
	if err != nil {
		return fmt.Sprint(response)
	}

	return string(output)
}
