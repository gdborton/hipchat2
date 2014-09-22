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
