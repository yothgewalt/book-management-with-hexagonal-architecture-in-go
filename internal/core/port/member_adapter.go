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
			return nil, errors.New("this username could not be created because it is already in use")
		}
	}

	return nil, nil
}

func (m memberRepository) GetAllMember() ([]*Member, error) {
	var members []*Member
	tx := m.database.Table("members").Select("id", "created_at", "username", "firstname", "lastname").Find(&members).Error
	if tx != nil {
		return nil, tx
	}

	return members, nil
}

func (m memberRepository) GetMemberById(uuid uuid.UUID) (*Member, error) {
	var member *Member
	tx := m.database.Table("members").Select("id", "created_at", "username", "firstname", "lastname").Where("id = ?", uuid).First(&member).Error
	if tx != nil {
		return nil, tx
	}

	return member, nil
}

func (m memberRepository) GetMemberByNameWithPassword(name string) (*Member, error) {
	var member *Member
	tx := m.database.Table("members").Select("id", "created_at", "username", "firstname", "lastname", "password").Where("username = ?", name).Take(&member).Error
	if tx != nil {
		return nil, tx
	}

	return member, nil
}

func (m memberRepository) DropMemberById(uuid uuid.UUID) error {
	var member *Member
	tx := m.database.Table("members").Select("id").Where("id = ?", uuid).Delete(&member).Error
	if tx != nil {
		return tx
	}

	return nil
}

func (m memberRepository) LoginMember(member Member) (*Member, error) {
	if err := m.database.Table("member").Select("username").Where("username = ?", member.Username).Take(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("this username could not be found, please register the member before login")
		} else {
			return nil, err
		}
	}

	return &member, nil
}
