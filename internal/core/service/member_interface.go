package service

import (
	"github.com/google/uuid"
	"time"
)

type NewMemberRequester struct {
	Username  string `json:"username" binding:"required,min=5,max=24"`
	Firstname string `json:"firstname" binding:"required,min=2,max=64"`
	Lastname  string `json:"lastname" binding:"required,min=2,max=64"`
	Password  string `json:"password" binding:"required,min=6"`
}

type MemberResponse struct {
	ID        uuid.UUID `json:"ID"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}

type MemberService interface {
	NewMember(requester NewMemberRequester) (*MemberResponse, error)
	ReadMembers() ([]*MemberResponse, error)
	ReadMemberById(uuid uuid.UUID) (*MemberResponse, error)
	DeleteMemberById(uuid uuid.UUID) error
}
