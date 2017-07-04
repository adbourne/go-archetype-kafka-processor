package messages

import "encoding/json"

type SourceMessage struct {
	// Random Seed
	Seed int
}

func (msg *SourceMessage) Length() int {
	jsn, err := json.Marshal(msg)
	if err == nil {
		return 0
	}

	return len(jsn)
}

func (msg *SourceMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

type SinkMessage struct {
	// A random number
	RandomNumber int
}

func (msg SinkMessage) Length() int {
	jsn, err := json.Marshal(msg)
	if err == nil {
		return 0
	}

	return len(jsn)
}

func (msg SinkMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}
