package util

import (
	"encoding/json"
	"testing"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
)

func TestRouterDispatchCallsRegisteredHandler(t *testing.T) {
	router := NewRouter()
	client := &Client{}
	payload := json.RawMessage(`{"message":"hello"}`)

	called := false
	router.On("chat.send", func(cl *Client, gotPayload json.RawMessage) {
		called = true

		if cl != client {
			t.Fatal("expected handler to receive the dispatched client")
		}

		if string(gotPayload) != string(payload) {
			t.Fatalf("expected payload %s, got %s", payload, gotPayload)
		}
	})

	router.Dispatch(client, requests.RequestMessage{Action: "chat.send", Payload: payload})

	if !called {
		t.Fatal("expected registered handler to be called")
	}
}

func TestRouterDispatchUnknownActionSendsError(t *testing.T) {
	router := NewRouter()
	client := &Client{send: make(chan responses.ResponseMessage, 1)}

	router.Dispatch(client, requests.RequestMessage{Action: "unknown.action"})

	msg := <-client.send
	if msg.Action != "unknown.action" {
		t.Fatalf("expected action %q, got %q", "unknown.action", msg.Action)
	}

	if msg.Success {
		t.Fatal("expected unknown action response to be unsuccessful")
	}

	if msg.Error != "unknown action" {
		t.Fatalf("expected error %q, got %q", "unknown action", msg.Error)
	}
}
