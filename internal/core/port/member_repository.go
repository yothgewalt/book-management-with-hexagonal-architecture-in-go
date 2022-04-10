package port

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Member struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Username  string       `gorm:"not null;unique;type:varchar(24)"`
	Firstname string       `gorm:"not null;type:varchar(64)"`
	Lastname  string       `gorm:"not null;type:varchar(64)"`
	Password  string       `gorm:"not null"`
}

type MemberRepository interface {
	CreateMember(member Member) (*Member, error)
	GetAllMember() ([]*Member, error)
	GetMemberById(uuid uuid.UUID) (*Member, error)
	DropMemberById() error
}
