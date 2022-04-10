package port

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type memberRepository struct {
	database *gorm.DB
}

func NewMemberRepository(database *gorm.DB) MemberRepository {
	return &memberRepository{database: database}
}

func (m memberRepository) CreateMember(member Member) (*Member, error) {
	if err := m.database.Table("members").Select("username").Where("username = ?", member.Username).Take(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx := m.database.Table("members").Model(&Member{}).Create(&member).Error
			if tx != nil {
				return nil, tx
			}

			return &member, nil
		} else {
			return nil, err
		}
	}

	return nil, nil
}

func (m memberRepository) GetAllMember() ([]*Member, error) {
	var members []*Member
	tx := m.database.Table("members").Select("id", "username", "firstname", "lastname").Find(&members).Error
	if tx != nil {
		return nil, tx
	}

	return members, nil
}

func (m memberRepository) GetMemberById(uuid uuid.UUID) (*Member, error) {
	var member *Member
	tx := m.database.Table("members").Select("id").Where("id = ?", uuid).First(&member).Error
	if tx != nil {
		return nil, tx
	}

	return member, nil
}

func (m memberRepository) DropMemberById() error {
	return nil
}
