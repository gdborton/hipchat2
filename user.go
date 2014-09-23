package hipchat2

import (
	"fmt"
	"encoding/json"
	"errors"
)

type User struct {
	// TODO - Presence
	// https://www.hipchat.com/docs/apiv2/method/view_user

	// XMPP/Jabber ID of the user.
	XmppJid string `json:"xmpp_jid"`

	// Whether the user has been deleted or not.
	IsDeleted bool `json:"is_deleted"`

	// User's full name.
	Name string `json:"name"`

	// The date in ISO-8601 when the user was last active.
	LastActive string `json:"last_active"`

	// User's title.
	Title string `json:"title"`

	// The date in ISO-8601 when the user was created.
	Created string `json:"created"`

	// User's ID
	Id int `json:"id"`

	// User's @mention name.
	MentionName string `json:"mention_name"`

	// Whether or not this user is an admin of the group.
	IsGroupAdmin bool `json:"is_group_admin"`

	// The desired user timezone.
	Timezone string `json:"timezone"`

	// Whether or not this user is a guest or registered user.
	IsGuest bool `json:"is_guest"`

	// User's email
	Email string `json:"email"`

	// URL to user's photo. 125px on the longest side.
	PhotoUrl string `json:"photo_url"`
}


func GetUsers () ([]User, error) {
	uri := fmt.Sprintf("https://%s/v2/user?auth_token=%s", Host, AuthToken)
	body, err := get(uri)
	if err != nil {
		return nil, err
	}

	usersResponse := &struct{ Items []User }{}

	if err := json.Unmarshal(body, usersResponse); err != nil {
		return nil, err
	}
	return usersResponse.Items, nil
}

func (user *User) Fetch () (error) {
	emailMentionOrName, err := user.emailIdOrMention()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("https://%s/v2/user/%s?auth_token=%s", Host, emailMentionOrName, AuthToken)
	body, err := get(uri)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &user); err != nil {
		return err
	}

	return nil
}

func (user *User) emailIdOrMention () (string, error) {
	switch {
		case user.Id != 0: return fmt.Sprintf("%d", user.Id), nil
		case user.Name != "": return user.Name, nil
		case user.MentionName != "": return user.MentionName, nil
		default: return "", errors.New("A user with a valid email, id, or @mention is required.")
	}
}
