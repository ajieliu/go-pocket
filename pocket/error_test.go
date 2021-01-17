package pocket

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewServiceError(t *testing.T) {
	testcases := []struct {
		msg  string
		code int
	}{
		{"Unauthorized", 401},
		{"Forbidden", 403},
	}

	for i, tc := range testcases {
		err := NewServiceError(tc.code, tc.msg)
		if err.Message != tc.msg {
			t.Fatalf("[%d] unmatched message: %s != %s", i, tc.msg, err.Message)
		}

		if err.Code != tc.code {
			t.Fatalf("[%d] unmatched code: %d != %d", i, tc.code, err.Code)
		}
	}
}

func TestBadRequestErr(t *testing.T) {

	testcases := []struct {
		msg string
	}{
		{""},
		{"test message"},
	}

	for i, tc := range testcases {
		err := BadRequestErr(tc.msg)
		if err.Code != http.StatusBadRequest {
			t.Fatal("unmatched code")
		}

		if err.Message != tc.msg {
			t.Fatalf("[%d] unmatched message", i)
		}
	}
}

func TestInternalServiceErr(t *testing.T) {

	testcases := []struct {
		msg string
	}{
		{""},
		{"internal service error"},
	}

	for i, tc := range testcases {
		err := InternalServiceErr(tc.msg)
		if err.Code != http.StatusInternalServerError {
			t.Fatal("unmatched code of internal service error")
		}

		if err.Message != tc.msg {
			t.Fatalf("[%d] unmatched message: %s != %s ", i, tc.msg, err.Message)
		}
	}
}

func TestNewResponseFromError(t *testing.T) {
	testcases := []struct {
		err        error
		expectCode int
	}{
		{fmt.Errorf("normal error"), 500},
		{&ServiceError{404, "not found"}, 404},
	}

	for _, tc := range testcases {
		resp := NewResponseFromError(tc.err)
		if resp.Code != tc.expectCode {
			t.Fatalf("unmatched code")
		}
		msg := resp.Data.(*errorResponse).Message
		if msg != tc.err.Error() {
			t.Fatalf("expected %s, got %s", tc.err.Error(), msg)
		}
	}
}
