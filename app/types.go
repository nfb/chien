package main

type Interaction struct {
	id             snowflake
	application_id snowflake
	interaction    interactionType
	data           interactionData
	//guild
	//guild_id
	//channel
	//channel_id
	//member
	//user
	//token
	//version
	//message
	//app_permissions
	//locale
	//guild_locale
	//entitlements
	//authorizing_integration_owners
	//context
}

type snowflake uint64
type resolvedData struct {
	Users    map[snowflake]string
	Members  map[snowflake]string
	Roles    map[snowflake]string
	Channels map[snowflake]string
	//Members     map[snowflake]string
	Attachments map[snowflake]string
}

type interactionType int
type interactionData struct {
	Id   snowflake
	Name string
	Type interactionType
}
