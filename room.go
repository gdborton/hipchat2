package hipchat2

type Room struct {
	// TODO - owner object, links, participants.

	// The ID of the room
	Id int `json:"id"`

	// The name of the room
	Name string `json: "name"`

	// Whether or not this room is archived.
	IsArchived bool `json: "is_archived"`

	// XMPP/Jabber ID of the room.
	XmppJid string `json: "xmpp_jid"`

	// Time the room was created in UTC.
	Created string `json: "created"`

	// URL for guest access, if enabled.
	GuestAccessUrl string `json: "guest_access_url"`

	// Whether or not guests can access this room.
	IsGuestAccessible bool `json: "is_guest_accessible"`

	// Last time the room was active in UTC.
	LastActive string `json: "last_active"`

	// Privacy Setting, valid values are ["private", "public"]
	Privacy string `json: "privacy"`

	// Current topic of the room.
	Topic string `json: topic`

}

