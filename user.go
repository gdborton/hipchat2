package hipchat2

type User struct {
	// TODO - Presence
	// https://www.hipchat.com/docs/apiv2/method/view_user

	// XMPP/Jabber ID of the user.
	XmppJid string `json: "xmpp_jid"`

	// Whether the user has been deleted or not.
	IsDeleted bool `json: "is_deleted"`

	// User's full name.
	Name string `json: "name"`

	// The date in ISO-8601 when the user was last active.
	LastActive string `json: "last_active"`

	// User's title.
	Title string `json: "title"`

	// The date in ISO-8601 when the user was created.
	Created string `json: "created"`

	// User's ID
	Id int `json: "id"`

	// User's @mention name.
	MentionName string `json: "mention_name"`

	// Whether or not this user is an admin of the group.
	IsGroupAdmin bool `json: "is_group_admin"`

	// The desired user timezone.
	Timezone string `json: "timezone"`

	// Whether or not this user is a guest or registered user.
	IsGuest bool `json: "is_guest"`

	// User's email
	Email string `json: "email"`

	// URL to user's photo. 125px on the longest side.
	PhotoUrl string `json: "photo_url"`
}
