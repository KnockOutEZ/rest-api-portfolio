package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Skill struct {
	ID               uint64    `gorm:"primary_key;auto_increment" json:"id"`
	SkillName        string    `gorm:"size:255;UNIQUE_INDEX:compositeindex;" json:"skill_name"`
	SkillTitle       string    `gorm:"size:255;" json:"skill_title"`
	SkillDescription string    `gorm:"size:255;" json:"skill_description"`
	SkillIcon        string    `gorm:"size:255;" json:"skill_icon"`
	SkillProgress    string    `gorm:"size:255;" json:"skill_progress"`
	SkillLinks       string    `gorm:"size:255;" json:"skill_links"`
	User             User      `json:"user"`
	UserID           uint32    `gorm:"UNIQUE_INDEX:compositeindex;not null" json:"user_id"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Skill) Prepare() {
	p.ID = 0
	p.SkillName = html.EscapeString(strings.TrimSpace(p.SkillName))
	p.SkillTitle = html.EscapeString(strings.TrimSpace(p.SkillTitle))
	p.SkillDescription = html.EscapeString(strings.TrimSpace(p.SkillDescription))
	p.SkillIcon = html.EscapeString(strings.TrimSpace(p.SkillIcon))
	p.SkillProgress = html.EscapeString(strings.TrimSpace(p.SkillProgress))
	p.SkillLinks = html.EscapeString(strings.TrimSpace(p.SkillLinks))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Skill) Validate() error {
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Skill) SaveSkill(db *gorm.DB) (*Skill, error) {
	var err error
	err = db.Debug().Model(&Skill{}).Create(&p).Error
	if err != nil {
		return &Skill{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Skill{}, err
		}
	}
	return p, nil
}

func (p *Skill) FindAllSkills(db *gorm.DB) (*[]Skill, error) {
	var err error
	skills := []Skill{}
	err = db.Debug().Model(&Skill{}).Find(&skills).Error
	if err != nil {
		return &[]Skill{}, err
	}
	if len(skills) > 0 {
		for i, _ := range skills {
			log.Println(skills[i].UserID)
			err := db.Debug().Model(&User{}).Where("id = ?", skills[i].UserID).Find(&skills[i].User).Error
			if err != nil {
				return &[]Skill{}, err
			}
		}
	}
	return &skills, nil
}

func (p *Skill) FindAllMySkills(db *gorm.DB, uid uint32) (*[]Skill, error) {
	var err error
	skills := []Skill{}
	err = db.Debug().Model(&Skill{}).Where("user_id = ?", uid).Limit(100).Find(&skills).Error
	if err != nil {
		return &[]Skill{}, err
	}
	if len(skills) > 0 {
		for i, _ := range skills {
			log.Println(skills[i].UserID)
			err := db.Debug().Model(&User{}).Where("id = ?", skills[i].UserID).Take(&skills[i].User).Error
			if err != nil {
				return &[]Skill{}, err
			}
		}
	}
	return &skills, nil
}

func (p *Skill) FindSkillByID(db *gorm.DB, pid uint64) (*Skill, error) {
	var err error
	err = db.Debug().Model(&Skill{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Skill{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Skill{}, err
		}
	}
	return p, nil
}

func (p *Skill) UpdateASkill(db *gorm.DB) (*Skill, error) {

	var err error

	err = db.Debug().Model(&Skill{}).Where("id = ?", p.ID).Updates(Skill{SkillName: p.SkillName, SkillDescription: p.SkillDescription, SkillTitle: p.SkillTitle, SkillIcon: p.SkillIcon, SkillProgress: p.SkillProgress,SkillLinks: p.SkillLinks, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Skill{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Skill{}, err
		}
	}
	return p, nil
}

func (p *Skill) DeleteASkill(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Skill{}).Where("id = ? and user_id = ?", pid, uid).Take(&Skill{}).Delete(&Skill{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Skill not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
