package model

import (
	"time"
)

type Profile struct {
	ID       string `gorm:"primaryKey"`
	UserID   string `gorm:"unique"`
	Name     string
	ImageURL string `gorm:"type:text"`
	Email    string `gorm:"type:text"`

	Servers  []Server  `gorm:"foreignKey:ProfileID"`
	Members  []Member  `gorm:"foreignKey:ProfileID"`
	Channels []Channel `gorm:"foreignKey:ProfileID"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Server struct {
	ID         string `gorm:"primaryKey"`
	Name       string
	ImageURL   string `gorm:"type:text"`
	InviteCode string `gorm:"unique"`

	ProfileID string
	Profile   Profile `gorm:"foreignKey:ProfileID"`

	Members  []Member  `gorm:"foreignKey:ServerID"`
	Channels []Channel `gorm:"foreignKey:ServerID"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Member struct {
	ID   string     `gorm:"primaryKey"`
	Role MemberRole `gorm:"default:GUEST"`

	ProfileID string
	Profile   Profile `gorm:"foreignKey:ProfileID"`

	ServerID string
	Server   Server `gorm:"foreignKey:ServerID"`

	Messages               []Message       `gorm:"foreignKey:MemberID"`
	DirectMessagesSent     []DirectMessage `gorm:"foreignKey:MemberID"`
	DirectMessagesReceived []DirectMessage `gorm:"foreignKey:MemberID"`

	ConversationsInitiated []Conversation `gorm:"foreignKey:MemberOneID"`
	ConversationsReceived  []Conversation `gorm:"foreignKey:MemberTwoID"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Channel struct {
	ID   string `gorm:"primaryKey"`
	Name string

	ProfileID string
	Profile   Profile `gorm:"foreignKey:ProfileID"`

	ServerID string
	Server   Server `gorm:"foreignKey:ServerID"`

	Type ChannelType `gorm:"default:TEXT"`

	Messages []Message `gorm:"foreignKey:ChannelID"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Message struct {
	ID      string `gorm:"primaryKey"`
	Content string `gorm:"type:text"`
	FileURL string `gorm:"type:text"`

	MemberID string
	Member   Member `gorm:"foreignKey:MemberID"`

	ChannelID string
	Channel   Channel `gorm:"foreignKey:ChannelID"`

	Deleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Conversation struct {
	ID          string `gorm:"primaryKey"`
	MemberOneID string
	MemberOne   Member `gorm:"foreignKey:MemberOneID"`

	MemberTwoID string
	MemberTwo   Member `gorm:"foreignKey:MemberTwoID"`

	DirectMessages []DirectMessage `gorm:"foreignKey:ConversationID"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type DirectMessage struct {
	ID      string `gorm:"primaryKey"`
	Content string `gorm:"type:text"`
	FileURL string `gorm:"type:text"`

	MemberID string
	Member   Member `gorm:"foreignKey:MemberID"`

	ConversationID string
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`

	Deleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type ServerImages struct {
	ID        uint `gorm:"primaryKey"`
	ImageURL  string
	ServerID  uint
	ProfileID string
}

type MessageFile struct {
	ID        uint `gorm:"primaryKey"`
	FilePath  string
	ServerID  uint
	ProfileID string
}

type MemberRole string

const (
	ADMIN     MemberRole = "ADMIN"
	MODERATOR MemberRole = "MODERATOR"
	GUEST     MemberRole = "GUEST"
)

type ChannelType string

const (
	TEXT  ChannelType = "TEXT"
	AUDIO ChannelType = "AUDIO"
	VIDEO ChannelType = "VIDEO"
)
