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
	IsArchived bool `json:"is_archived"`

	// XMPP/Jabber ID of the room.
	XmppJid string `json:"xmpp_jid,omitempty"`

	// Time the room was created in UTC.
	Created string `json:"created,omitempty"`

	// URL for guest access, if enabled.
	GuestAccessUrl string `json:"guest_access_url,omitempty"`

	// Whether or not guests can access this room.
	IsGuestAccessible bool `json:"is_guest_accessible"`

	// Last time the room was active in UTC.
	LastActive string `json:"last_active,omitempty"`

	// Privacy Setting, valid values are ["private", "public"]
	Privacy string `json:"privacy,omitempty"`

	// Current topic of the room.
	Topic string `json:"topic,omitempty"`

	Owner *User `json:"owner,omitempty"`

	ownerUserId string `json:"owner_user_id,omitempty"`  // Dirty extra field is needed for creating rooms.

}

// TODO -- share file
// TODO -- view recent history
// TODO -- create webhook
// TODO -- get all webhooks
// TODO -- get room statistics
// TODO -- reply to message
// TODO -- get all members
// TODO -- share link with room
// TODO -- add member
// TODO -- remove member
// TODO -- delete webhook
// TODO -- get webhook
// TODO -- view room history
// TODO -- get room message
// TODO -- get all participants


// Get all rooms - https://www.hipchat.com/docs/apiv2/method/get_all_rooms
// TODO -- query params: start-index, max-results, include-archived
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

// Get Room - https://www.hipchat.com/docs/apiv2/method/get_room
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

// See save existing/save new.
func (room *Room) Save () (error) {
	if room.Id == 0 {  // If you don't have an Id, assume the room doesn't yet exist.
		return room.saveNewRoom()
	}else {
		return room.saveExistingRoom()
	}
}

// Update Existing Room - https://www.hipchat.com/docs/apiv2/method/update_room
func (room *Room) saveExistingRoom () (error) {
	nameOrId, err := room.nameOrId()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("https://%s/v2/room/%s?auth_token=%s", Host, nameOrId, AuthToken)
	payload, marshalError := json.Marshal(room)
	if marshalError != nil {
		return marshalError
	}
	body, err := put(uri, payload)
	_ = body
	if err != nil {
		return err
	}
	return nil
}

// Create Room - https://www.hipchat.com/docs/apiv2/method/create_room
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

// Invite a user to the room - https://www.hipchat.com/docs/apiv2/method/invite_user
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

// Delete Room - https://www.hipchat.com/docs/apiv2/method/delete_room
func (room *Room) Delete () (error) {
	nameOrId, err := room.nameOrId()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("https://%s/v2/room/%s?auth_token=%s", Host, nameOrId, AuthToken)
	return delete(uri)
}

// Send Notification to the room - https://www.hipchat.com/docs/apiv2/method/send_room_notification
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

// Set Topic for the room - https://www.hipchat.com/docs/apiv2/method/set_topic
func (room *Room) SetTopic (topic string) (error) {
	nameOrId, err := room.nameOrId()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("https://%s/v2/room/%s/topic?auth_token=%s", Host, nameOrId, AuthToken)
	room.Topic = topic
	payload, err := json.Marshal(room)
	if err != nil {
		return err
	}
	body, err := put(uri, payload)
	_ = body
	if err != nil {
		return err
	}
	return nil
}

// Checks the room for a name or id field to use with the api, returns an error if neither are found.
func (room *Room) nameOrId () (string, error) {
	switch {
		case room.Id != 0: return fmt.Sprintf("%d", room.Id), nil
		case room.Name != "": return room.Name, nil
		default: return "", errors.New("A room with a valid name or id is required.")
	}
}
