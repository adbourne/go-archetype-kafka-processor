package messages

import "encoding/json"

// SourceMessage is the message picked from the source topic
type SourceMessage struct {
	// Random Seed
	Seed int
}

// Length returns the number of bytes in the SourceMessage JSON representation
// Required by Sarama
func (msg *SourceMessage) Length() int {
	jsn, err := json.Marshal(msg)
	if err == nil {
		return 0
	}

	return len(jsn)
}

// Encode encodes the SourceMessage into a JSON representation
// Required by Sarama
func (msg *SourceMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

// SinkMessage is the message the application places on the output topic
type SinkMessage struct {
	// A random number
	RandomNumber int
}

// Length returns the number of bytes in the SinkMessage JSON representation
// Required by Sarama
func (msg SinkMessage) Length() int {
	jsn, err := json.Marshal(msg)
	if err == nil {
		return 0
	}

	return len(jsn)
}

// Encode encodes the SinkMessage into a JSON representation
// Required by Sarama
func (msg SinkMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}
