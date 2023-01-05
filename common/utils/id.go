package utils

import "github.com/lithammer/shortuuid/v4"

func GenSuitId() string {
	return "suit_" + shortuuid.New()
}

func GenTestId() string {
	return "test_" + shortuuid.New()
}

func GenRequestResponseId() string {
	return "re_" + shortuuid.New()
}
