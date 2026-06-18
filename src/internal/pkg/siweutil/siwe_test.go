package siweutil_test

import (
	"encoding/json"
	"testing"
	"time"

	siwe "github.com/spruceid/siwe-go"

	"github.com/seedao/seedao-api-server/internal/pkg/siweutil"
)

func TestBuildMessageRoundTrip(t *testing.T) {
	nonce := siwe.GenerateNonce()
	msg, err := siweutil.BuildMessage("app.seedao.top", "0x8c913aec7443fe2018639133398955e0e17fb0c1", nonce, 1, time.Now())
	if err != nil {
		t.Fatal(err)
	}

	if _, err := siwe.ParseMessage(msg); err != nil {
		t.Fatalf("parse built message: %v\nmsg=%q", err, msg)
	}

	payload, err := json.Marshal(map[string]string{"message": msg})
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]interface{}
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	roundtrip := decoded["message"].(string)

	if _, err := siwe.ParseMessage(roundtrip); err != nil {
		t.Fatalf("parse roundtrip message: %v", err)
	}

	_, err = siweutil.Verify(siweutil.VerifyInput{
		Wallet:    "0x8c913aec7443fe2018639133398955e0e17fb0c1",
		Domain:    "app.seedao.top",
		Message:   roundtrip,
		Signature: "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		Nonce:     nonce,
	})
	if err == nil {
		t.Fatal("expected verify error")
	}
	if err != siweutil.ErrSignatureInvalid {
		t.Fatalf("expected signature invalid, got %v", err)
	}
}
