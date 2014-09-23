package hipchat2

import (
	"os"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	user := &User{MentionName: "@TesterMcTesterson"}
	err := user.Fetch()
	assert.Nil(t, err, "Shouldn't find an error.")
	assert.Equal(t, "Tester McTesterson", user.Name)
}

func TestGetUsers(t *testing.T) {
	AuthToken = os.Getenv("TEST_TOKEN")
	users, err := GetUsers()
	assert.Nil(t, err, "Shouldn't find an error")
	assert.Equal(t, true, len(users) > 0, "Expect to find more than zero users.")
}
