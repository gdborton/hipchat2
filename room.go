package hipchat2

import (
	"fmt"
	"encoding/json"
	"errors"
)

type Room struct {
	// TODO - owner object, links, participants.

	// The ID of the room
	Id int `json:"id,omitempty"`

	// The name of the room
	Name string `json:"name,omitempty"`

	// Whether or not this room is archived.
	IsArchived bool `json:"is_archived,omitempty"`

	// XMPP/Jabber ID of the room.
	XmppJid string `json:"xmpp_jid,omitempty"`

	// Time the room was created in UTC.
	Created string `json:"created,omitempty"`

	// URL for guest access, if enabled.
	GuestAccessUrl string `json:"guest_access_url,omitempty"`

	// Whether or not guests can access this room.
	IsGuestAccessible bool `json:"is_guest_accessible,omitempty"`

	// Last time the room was active in UTC.
	LastActive string `json:"last_active,omitempty"`

	// Privacy Setting, valid values are ["private", "public"]
	Privacy string `json:"privacy,omitempty"`

	// Current topic of the room.
	Topic string `json:"topic,omitempty"`

	Owner *User `json:"owner,omitempty"`

	ownerUserId string `json:"owner_user_id,omitempty"`  // Dirty extra field is needed for creating rooms.

}

// TODO -- invite
// TODO -- share file
// TODO -- create
// TODO -- get all
// TODO -- view recent history
// TODO -- send notification
// TODO -- update room
// TODO -- get room
// TODO -- delete room
// TODO -- create webhook
// TODO -- get all webhooks
// TODO -- get room statistics
// TODO -- reply to message
// TODO -- get all members
// TODO -- set topic
// TODO -- share link with room
// TODO -- add member
// TODO -- remove member
// TODO -- delete webhook
// TODO -- get webhook
// TODO -- view room history
// TODO -- get room message
// TODO -- get all participants

func GetRooms () ([]Room, error) {
	uri := fmt.Sprintf("https://%s/v2/room?auth_token=%s", Host, AuthToken)

	body, err := get(uri)
	if err != nil {
		return nil, err
	}

	roomsResp := &struct{ Items []Room }{}

	if err := json.Unmarshal(body, roomsResp); err != nil {
		return nil, err
	}

	return roomsResp.Items, nil
}

func (room *Room) Fetch () (error) {
	nameOrId, err := room.nameOrId()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("https://%s/v2/room/%s?auth_token=%s", Host, nameOrId, AuthToken)

	body, err := get(uri)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &room); err != nil {
		return err
	}

	return nil
}

func (room *Room) Save () (error) {
	if room.Id == 0 {
		return room.saveNewRoom()
	}else {
		return room.saveExistingRoom()
	}
}

func (room *Room) saveExistingRoom () (error) {
	// TODO - finish me.
	return nil
}

func (room *Room) saveNewRoom () (error) {
	uri := fmt.Sprintf("https://%s/v2/room?auth_token=%s", Host, AuthToken)

	ownerId, err := room.Owner.emailIdOrMention()
	if err != nil {
		return err
	}
	room.ownerUserId = ownerId
	payload, marshalError := json.Marshal(room)
	if marshalError != nil {
		return marshalError
	}
	body, err := post(uri, payload)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, room); err != nil {
		return err
	}
	return nil
}

func (room *Room) InviteUser (user *User, reason string) (error) {
	roomNameOrId, err := room.nameOrId()
	if err != nil {
		return err
	}

	userIdOrEmail, err := user.emailIdOrMention()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("https://%s/v2/room/%s/invite/%s?auth_token=%s", Host, roomNameOrId, userIdOrEmail, AuthToken)
	payload, err := json.Marshal(&struct{Reason string `json:"reason"`}{reason})
	if err != nil {
		return err
	}
	body, err := post(uri, payload)
	_ = body
	if err != nil {
		return err
	}
	return nil
}

func (room *Room) Delete () (error) {
	nameOrId, err := room.nameOrId()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("https://%s/v2/room/%s?auth_token=%s", Host, nameOrId, AuthToken)
	return delete(uri)
}

func (room *Room) SendNotification (message *Message) (error) {
	nameOrId, err := room.nameOrId()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("https://%s/v2/room/%s/notification?auth_token=%s", Host, nameOrId, AuthToken)

	payload, marshalError := json.Marshal(message)
	if marshalError != nil {
		return marshalError
	}
	body, err := post(uri, payload)
	_ = body
	if err != nil {
		return err
	}
	return nil
}

func (room *Room) nameOrId () (string, error) {
	switch {
		case room.Id != 0: return fmt.Sprintf("%s", room.Id), nil
		case room.Name != "": return room.Name, nil
		default: return "", errors.New("A room with a valid name or id is required.")
	}
}
