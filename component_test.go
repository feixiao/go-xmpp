package xmpp

import "testing"

func TestHandshake(t *testing.T) {
	c := Component{
		Host:   "test.localhost",
		Secret: "mypass",
	}

	streamID := "1263952298440005243"
	expected := "c77e2ef0109fbbc5161e83b51629cd1353495332"

	result := c.Handshake(streamID)
	if result != expected {
		t.Errorf("incorrect handshake calculation '%s' != '%s'", result, expected)
	}
}
