package hipchat2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"fmt"
)

// These are mostly integration tests, but should be generic enough to be run on any account with a little setup.

func TestGetRooms(t *testing.T) {
	client := NewClient(os.Getenv("TEST_TOKEN"))
	rooms, err := client.GetRooms()
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, true, len(rooms) > 0, "Expect to find at least one room.")
	assert.Equal(t, "test", rooms[0].Name, "Expect the first room to be the test room.")
}

func TestGetRoom(t *testing.T) {
	client := NewClient(os.Getenv("TEST_TOKEN"))
	room, err := client.GetRoom("test")
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, false, room.IsArchived, "Expect the room to not be archived.")
	assert.Equal(t, "public", room.Privacy, "Expect the room to be public.")
}

func TestCreateRoom(t *testing.T) {
	client := NewClient(os.Getenv("TEST_TOKEN"))
	user, err := client.GetUser("@TesterMcTesterson")
	assert.Nil(t, err, "Shouldn't find an error")
	roomId, err := client.CreateRoom("topic", false, "Test Room", fmt.Sprintf("%d", user.Id), "public")
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, true, roomId > 0, "roomId should be a valid id greater than zero.")
}

func TestGetUsers(t *testing.T) {
	client := NewClient(os.Getenv("TEST_TOKEN"))
	users, err := client.GetUsers()
	assert.Nil(t, err, "Shouldn't find an error")
	assert.Equal(t, true, len(users) > 0, "Expect to find more than zero users.")
}

func TestGetUser(t *testing.T) {
	client := NewClient(os.Getenv("TEST_TOKEN"))
	user, err := client.GetUser("@TesterMcTesterson")
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, "Tester McTesterson", user.Name)
}

func TestDefaultHost(t *testing.T) {
	client := NewClient("a")
	assert.Equal(t, "api.hipchat.com", client.Host)
}

func TestSetHost(t *testing.T) {
	client := NewClient("a")
	client.Host = "example.hipchat.com"
	assert.Equal(t, "example.hipchat.com", client.Host)
}

func TestAuthToken(t *testing.T) {
	client := NewClient("a")
	assert.Equal(t, "a", client.AuthToken)
}

