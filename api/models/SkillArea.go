package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SkillArea struct {
	ID                   uint64    `gorm:"primary_key;auto_increment" json:"id"`
	SkillAreaName        string    `gorm:"size:255;UNIQUE_INDEX:skillareaindex;" json:"skill_area_name"`
	SkillAreaTitle       string    `gorm:"size:255;" json:"skill_area_title"`
	SkillAreaDescription string    `gorm:"size:255;" json:"skill_area_description"`
	SkillAreaIcon        string    `gorm:"size:255;" json:"skill_area_icon"`
	SkillAreaProgress    string    `gorm:"size:255;" json:"skill_area_progress"`
	SkillAreaLinks       string    `gorm:"size:255;" json:"skill_area_links"`
	User                 User      `json:"-"`
	UserID               uint32    `gorm:"UNIQUE_INDEX:skillareaindex;not null" json:"user_id"`
	CreatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *SkillArea) Prepare() {
	p.ID = 0
	p.SkillAreaName = html.EscapeString(strings.TrimSpace(p.SkillAreaName))
	p.SkillAreaTitle = html.EscapeString(strings.TrimSpace(p.SkillAreaTitle))
	p.SkillAreaDescription = html.EscapeString(strings.TrimSpace(p.SkillAreaDescription))
	p.SkillAreaIcon = html.EscapeString(strings.TrimSpace(p.SkillAreaIcon))
	p.SkillAreaProgress = html.EscapeString(strings.TrimSpace(p.SkillAreaProgress))
	p.SkillAreaLinks = html.EscapeString(strings.TrimSpace(p.SkillAreaLinks))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *SkillArea) Validate() error {
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *SkillArea) SaveSkillArea(db *gorm.DB) (*SkillArea, error) {
	var err error
	err = db.Debug().Model(&SkillArea{}).Create(&p).Error
	if err != nil {
		return &SkillArea{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &SkillArea{}, err
	// 	}
	// }
	return p, nil
}

func (p *SkillArea) FindAllSkillAreas(db *gorm.DB) (*[]SkillArea, error) {
	var err error
	skill_areas := []SkillArea{}
	err = db.Debug().Model(&SkillArea{}).Find(&skill_areas).Error
	if err != nil {
		return &[]SkillArea{}, err
	}
	// if len(skill_areas) > 0 {
	// 	for i, _ := range skill_areas {
	// 		log.Println(skill_areas[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", skill_areas[i].UserID).Find(&skill_areas[i].User).Error
	// 		if err != nil {
	// 			return &[]SkillArea{}, err
	// 		}
	// 	}
	// }
	return &skill_areas, nil
}

func (p *SkillArea) GoFindAllMySkillAreas(db *gorm.DB, uid uint64) (*[]SkillArea, error) {
	var err error
	skill_areas := []SkillArea{}
	err = db.Debug().Model(&SkillArea{}).Where("user_id = ?", uid).Limit(100).Find(&skill_areas).Error
	if err != nil {
		return &[]SkillArea{}, err
	}
	// if len(skill_areas) > 0 {
	// 	for i, _ := range skill_areas {
	// 		log.Println(skill_areas[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", skill_areas[i].UserID).Take(&skill_areas[i].User).Error
	// 		if err != nil {
	// 			return &[]SkillArea{}, err
	// 		}
	// 	}
	// }
	return &skill_areas, nil
}

func (p *SkillArea) GoFindSkillAreaByID(db *gorm.DB, pid uint64, uid uint64) (*SkillArea, error) {
	var err error
	err = db.Debug().Model(&SkillArea{}).Where("id = ?", pid).Where("user_id = ?", uid).Take(&p).Error
	if err != nil {
		return &SkillArea{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &SkillArea{}, err
	// 	}
	// }
	return p, nil
}

func (p *SkillArea) FindAllMySkillAreas(db *gorm.DB, uid uint32) (*[]SkillArea, error) {
	var err error
	skill_areas := []SkillArea{}
	err = db.Debug().Model(&SkillArea{}).Where("user_id = ?", uid).Limit(100).Find(&skill_areas).Error
	if err != nil {
		return &[]SkillArea{}, err
	}
	// if len(skill_areas) > 0 {
	// 	for i, _ := range skill_areas {
	// 		log.Println(skill_areas[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", skill_areas[i].UserID).Take(&skill_areas[i].User).Error
	// 		if err != nil {
	// 			return &[]SkillArea{}, err
	// 		}
	// 	}
	// }
	return &skill_areas, nil
}

func (p *SkillArea) FindSkillAreaByID(db *gorm.DB, pid uint64) (*SkillArea, error) {
	var err error
	err = db.Debug().Model(&SkillArea{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &SkillArea{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &SkillArea{}, err
	// 	}
	// }
	return p, nil
}

func (p *SkillArea) UpdateASkillArea(db *gorm.DB) (*SkillArea, error) {

	var err error

	err = db.Debug().Model(&SkillArea{}).Where("id = ?", p.ID).Updates(SkillArea{SkillAreaName: p.SkillAreaName, SkillAreaDescription: p.SkillAreaDescription, SkillAreaTitle: p.SkillAreaTitle, SkillAreaIcon: p.SkillAreaIcon, SkillAreaProgress: p.SkillAreaProgress, SkillAreaLinks: p.SkillAreaLinks, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &SkillArea{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &SkillArea{}, err
	// 	}
	// }
	return p, nil
}

func (p *SkillArea) DeleteASkillArea(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&SkillArea{}).Where("id = ? and user_id = ?", pid, uid).Take(&SkillArea{}).Delete(&SkillArea{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("SkillArea not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
