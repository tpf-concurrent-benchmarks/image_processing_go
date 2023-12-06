package utils

import (
	"os"
)

func GetNodeID() string {
	if os.Getenv("LOCAL") == "" {
		return os.Getenv("NODE_ID")
	}
	return "resolution_worker"
}