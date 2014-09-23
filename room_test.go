package hipchat2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	testRoomName = "Test Room1"
)

func TestCreateRoom(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	user := &User{MentionName: "@TesterMcTesterson"}
	userFetchError := user.Fetch()
	assert.Nil(t, userFetchError, "Shouldn't find an error")
	room := &Room{Name: testRoomName, Topic: "topic", IsGuestAccessible: false, Privacy: "public", Owner: user}
	err := room.Save()
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, true, room.Id > 0, "roomId should be a valid id greater than zero.")
}

func TestGetRoom(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	room := &Room{Name: testRoomName}
	err := room.Fetch()
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, false, room.IsArchived, "Expect the room to not be archived.")
	assert.Equal(t, "public", room.Privacy, "Expect the room to be public.")
}

func TestGetRooms(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	rooms, err := GetRooms()
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, true, len(rooms) > 0, "Expect to find at least one room.")
}

func TestSendRoomNotification(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	message := &Message{
		Message: "This is a test message.",
		Notify: false,
		Color: "green",
		MessageFormat: "text"}

	room := &Room{
		Name: "test"}
	err := room.SendNotification(message)
	assert.Nil(t, err, "Shouldn't find an error.")
}

func TestInviteUser(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	user := &User{MentionName: "@TesterMcTesterson"}
	user.Fetch()
	room := &Room{Name: testRoomName}
	err := room.InviteUser(user, "Any reason that I want.")
	assert.Nil(t, err, "Shouldn't find an error.")
}

func TestUpdate(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	room := &Room{Name: testRoomName}
	err := room.Fetch()
	assert.Nil(t, err, "Shouldn't find an error.")
	topicString := "My new topic"
	room.Topic = topicString
	room.Save()
	room2 := &Room{Name: testRoomName}
	err = room2.Fetch()
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, topicString, room2.Topic, "Expect to find the topic string that was set.")
}

func TestSetTopic(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	room := &Room{Name: testRoomName}
	err := room.Fetch()
	assert.Nil(t, err, "Shouldn't find an error.")
	topicString := "My new topic2"
	room.SetTopic(topicString)
	room2 := &Room{Name: testRoomName}
	err = room2.Fetch()
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, topicString, room2.Topic, "Expect to find the topic string that was set.")
}

func TestDelete(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	room := &Room{Name: testRoomName}
	err := room.Delete()
	assert.Nil(t, err, "Shouldn't find an error.")
}

/*
func TestGetUsers(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	users, err := client.GetUsers()
	assert.Nil(t, err, "Shouldn't find an error")
	assert.Equal(t, true, len(users) > 0, "Expect to find more than zero users.")
}
*/
