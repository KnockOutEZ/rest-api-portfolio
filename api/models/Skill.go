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
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Skill) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Skill) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("Required Author")
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
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
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
			log.Println(skills[i].AuthorID)
			err := db.Debug().Model(&User{}).Where("id = ?", skills[i].AuthorID).Find(&skills[i].Author).Error
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
	err = db.Debug().Model(&Skill{}).Where("author_id = ?", uid).Limit(100).Find(&skills).Error
	if err != nil {
		return &[]Skill{}, err
	}
	if len(skills) > 0 {
		for i, _ := range skills {
			log.Println(skills[i].AuthorID)
			err := db.Debug().Model(&User{}).Where("id = ?", skills[i].AuthorID).Take(&skills[i].Author).Error
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
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Skill{}, err
		}
	}
	return p, nil
}

func (p *Skill) UpdateASkill(db *gorm.DB) (*Skill, error) {

	var err error

	err = db.Debug().Model(&Skill{}).Where("id = ?", p.ID).Updates(Skill{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Skill{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Skill{}, err
		}
	}
	return p, nil
}

func (p *Skill) DeleteASkill(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Skill{}).Where("id = ? and author_id = ?", pid, uid).Take(&Skill{}).Delete(&Skill{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Skill not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
