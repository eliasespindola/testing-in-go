package user

import (
	"testing"
	"time"
)

func TestGetOne(t *testing.T) {
	expected := User{
		ID:       42,
		Username: "mrobot",
	}
	users = []User{expected}

	got, err := getOne(expected.ID)

	if err != nil {
		t.Fatal(err)
	}

	if got != expected {
		t.Errorf("did not get expected user. Got %+v, expected %+v", got, expected)
	}
}

func TestSlowOne(t *testing.T) {
	t.Parallel()
	t.Skip("skipped")
	time.Sleep(1 * time.Second)
}

func TestSlowTwo(t *testing.T) {
	t.Parallel()
	t.Skip("skipped")
	time.Sleep(1 * time.Second)
}
