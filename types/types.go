package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebhookRequestPayload struct {
	Content 		 string					`json:"content,omitempty"`
	Username 		 string					`json:"username,omitempty"`
	AvatarURL 		 string					`json:"avatar_url,omitempty"`
	Tts 			 bool					`json:"tts,omitempty"`
	Embeds 			 []DiscordEmbed			`json:"embeds,omitempty"`
	Allowed_mentions DiscordAllowedMentions	`json:"allowed_mentions,omitempty"`
}

type WebhookResponsePayload struct {
	Id 					string			`json:"id"`
	Type 				int				`json:"type"`
	Content 			string			`json:"content"`
	Channel_id 			string			`json:"channel_id"`
	Author				WebhookAuthor	`json:"author"`
	Attachments 		[]interface{}	`json:"attachments"`
	Embeds 				[]DiscordEmbed	`json:"embeds"`
	Mentions 			[]interface{}	`json:"mentions"`
	Mention_roles 		[]interface{}	`json:"mention_roles"`
	Pinned 				bool			`json:"pinned"`
	Mention_everyone 	bool			`json:"mention_everyone"`
	Tts 				bool			`json:"tts"`
	Timestamp 			string			`json:"timestamp"`
	Edited_timestamp 	string			`json:"edited_timestamp"`
	Flags 				int				`json:"flags"`
	Components 			[]interface{}	`json:"components"`
	Webhook_id 			string			`json:"webhook_id"`
}

type DiscordEmbed struct {
	Title 			string			`json:"title,omitempty"`
	Type 			string			`json:"type,omitempty"`
	Description 	string			`json:"description,omitempty"`
	Url 			string			`json:"url,omitempty"`
	Timestamp 		string			`json:"timestamp,omitempty"`
	Color 			int				`json:"color,omitempty"`
	Footer 			EmbedFooter		`json:"footer,omitempty"`
	Image 			EmbedImage		`json:"image,omitempty"`
	Thumbnail 		EmbedThumbnail	`json:"thumbnail,omitempty"`
	Video 			EmbedVideo		`json:"video,omitempty"`
	Provider 		EmbedProvider	`json:"provider,omitempty"`
	Author 			EmbedAuthor		`json:"author,omitempty"`
	Fields 			[]EmbedField	`json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text 			string	`json:"text"`
	Icon_url 		string	`json:"icon_url,omitempty"`
	Proxy_icon_url 	string	`json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	Url		  string	`json:"url"`
	Proxy_url string	`json:"proxy_url,omitempty"`
	Height	  int		`json:"height,omitempty"`
	Width	  int		`json:"width,omitempty"`
}

type EmbedThumbnail struct {
	Url		  string	`json:"url"`
	Proxy_url string	`json:"proxy_url,omitempty"`
	Height	  int		`json:"height,omitempty"`
	Width	  int		`json:"width,omitempty"`
}

type EmbedVideo struct {
	Url		  string	`json:"url"`
	Proxy_url string	`json:"proxy_url,omitempty"`
	Height	  int		`json:"height,omitempty"`
	Width	  int		`json:"width,omitempty"`
}

type EmbedProvider struct {
	Name	string	`json:"name,omitempty"`
	Url		string	`json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name	string			`json:"name"`
	Url		string			`json:"url,omitempty"`
	Icon_url string			`json:"icon_url,omitempty"`
	Proxy_icon_url string	`json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name 	string 	`json:"name"`
	Value 	string 	`json:"value"`
	Inline 	bool 	`json:"inline,omitempty"`
}

type DiscordAllowedMentions struct {
	Parse []string		`json:"parse,omitempty"`
	Roles []string		`json:"roles,omitempty"`
	Users []string		`json:"users,omitempty"`
	Replied_user bool	`json:"replied_user,omitempty"`
}

type WebhookAuthor struct {
	Bot 		  bool   `json:"bot"`
	Id 			  string `json:"id"`
	Username 	  string `json:"username"`
	Avatar 		  string `json:"avatar"`
	Discriminator string `json:"discriminator"`
}

type StatusPageIncidentsResponse struct {
	Page 		StatusPagePageInformation 	`json:"page"`
	Incidents 	[]StatusPageIncident		`json:"incidents"`
}

type StatusPagePageInformation struct {
	Id 			string `json:"id"`
	Name 		string `json:"name"`
	Url 		string `json:"url"`
	Timezone 	string `json:"time_zone"`
	UpdatedAt 	string `json:"updated_at"`
}

type StatusPageIncident struct {
	Id 				string 						`json:"id"`
	Name 			string 						`json:"name"`
	Status 			string 						`json:"status"`
	CreatedAt 		string 						`json:"created_at,omitempty"`
	UpdatedAt 		string 						`json:"updated_at,omitempty"`
	ResolvedAt 		string 						`json:"resolved_at,omitempty"`
	Impact 			string 						`json:"impact"`
	Shortlink 		string 						`json:"shortlink"`
	StartedAt 		string 						`json:"started_at"`
	PageId 			string 						`json:"page_id"`
	IncidentUpdates []StatusPageIncidentUpdate 	`json:"incident_updates"`
	Components 		[]StatusPageComponent 		`json:"components"`
}

type StatusPageIncidentUpdate struct {
	Id 						string 						`json:"id"`
	Status 					string 						`json:"status"`
	Body 					string 						`json:"body"`
	IncidentId 				string 						`json:"incident_id"`
	CreatedAt 				string 						`json:"created_at"`
	UpdatedAt 				string 						`json:"update_at"`
	DisplayAt 				string 						`json:"display_at"`
	AffectedComponents 		[]StatusPageComponentUpdate `json:"affected_components"`
	DeliverNotifications 	bool 						`json:"deliver_notifications"`
	CustomTweet 			string 						`json:"custom_tweet,omitempty"`
	TweetId 				string 						`json:"tweet_id,omitempty"`
}

type StatusPageComponent struct {
	Id 					string 	`json:"id"`
	Name 				string 	`json:"name"`
	Status 				string 	`json:"status"`
	CreatedAt 			string 	`json:"created_at"`
	UpdatedAt 			string 	`json:"updated_at"`
	Position 			int 	`json:"position"`
	Description 		string 	`json:"description"`
	Showcase 			bool	`json:"showcase"`
	StartDate			string	`json:"start_date,omitempty"`
	GroupId 			string	`json:"group_id,omitempty"`
	PageId 				string	`json:"page_id"`
	Group 				bool	`json:"group,omitempty"`
	OnlyShowIfDegraded 	bool	`json:"only_show_if_degraded"`
}

type StatusPageComponentUpdate struct {
	Code 			string `json:"code"`
	Name 			string `json:"name"`
	OldStatus 		string `json:"old_status"`
	NewStatus 		string `json:"new_status"`
}

type MongoIncident struct {
	Id 				primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	MessageID 		string 			   `json:"message_id" bson:"message_id"`
	IncidentID		string 			   `json:"incident_id" bson:"incident_id"`
	LastUpdate		int64			   `json:"last_update" bson:"last_update"`
	Resolved		bool			   `json:"resolved" bson:"resolved"`
}