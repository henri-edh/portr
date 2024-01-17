// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type Connection struct {
	ID           string
	Type         string
	Subdomain    interface{}
	Port         interface{}
	Status       string
	TeamMemberID int64
	CreatedAt    time.Time
	StartedAt    interface{}
	ClosedAt     interface{}
	TeamID       interface{}
}

type GlobalSetting struct {
	ID                     int64
	SmtpEnabled            bool
	SmtpHost               interface{}
	SmtpPort               interface{}
	SmtpUsername           interface{}
	SmtpPassword           interface{}
	FromAddress            interface{}
	AddMemberEmailSubject  interface{}
	AddMemberEmailTemplate interface{}
}

type Session struct {
	ID        int64
	UserID    int64
	Token     string
	CreatedAt time.Time
}

type Team struct {
	ID        int64
	Name      string
	Slug      string
	CreatedAt time.Time
}

type TeamMember struct {
	ID            int64
	UserID        int64
	TeamID        int64
	SecretKey     string
	Role          string
	AddedByUserID interface{}
	CreatedAt     time.Time
}

type User struct {
	ID                int64
	Email             string
	FirstName         interface{}
	LastName          interface{}
	IsSuperUser       bool
	GithubAccessToken interface{}
	GithubAvatarUrl   interface{}
	CreatedAt         time.Time
}
