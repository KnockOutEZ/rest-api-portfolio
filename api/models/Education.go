package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Education struct {
	ID                   uint64    `gorm:"primary_key;auto_increment" json:"id"`
	EducationName        string    `gorm:"size:255;UNIQUE_INDEX:educationindex;" json:"education_name"`
	EducationTitle       string    `gorm:"size:255;" json:"education_title"`
	EducationDescription string    `gorm:"size:255;" json:"education_description"`
	EducationIcon        string    `gorm:"size:255;" json:"education_icon"`
	EducationProgress    string    `gorm:"size:255;" json:"education_progress"`
	EducationLinks       string    `gorm:"size:255;" json:"education_links"`
	User                 User      `json:"-"`
	UserID               uint32    `gorm:"UNIQUE_INDEX:educationindex;not null" json:"user_id"`
	CreatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Education) Prepare() {
	p.ID = 0
	p.EducationName = html.EscapeString(strings.TrimSpace(p.EducationName))
	p.EducationTitle = html.EscapeString(strings.TrimSpace(p.EducationTitle))
	p.EducationDescription = html.EscapeString(strings.TrimSpace(p.EducationDescription))
	p.EducationIcon = html.EscapeString(strings.TrimSpace(p.EducationIcon))
	p.EducationProgress = html.EscapeString(strings.TrimSpace(p.EducationProgress))
	p.EducationLinks = html.EscapeString(strings.TrimSpace(p.EducationLinks))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Education) Validate() error {
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Education) SaveEducation(db *gorm.DB) (*Education, error) {
	var err error
	err = db.Debug().Model(&Education{}).Create(&p).Error
	if err != nil {
		return &Education{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Education{}, err
	// 	}
	// }
	return p, nil
}

func (p *Education) FindAllEducations(db *gorm.DB) (*[]Education, error) {
	var err error
	educations := []Education{}
	err = db.Debug().Model(&Education{}).Find(&educations).Error
	if err != nil {
		return &[]Education{}, err
	}
	// if len(educations) > 0 {
	// 	for i, _ := range educations {
	// 		log.Println(educations[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", educations[i].UserID).Find(&educations[i].User).Error
	// 		if err != nil {
	// 			return &[]Education{}, err
	// 		}
	// 	}
	// }
	return &educations, nil
}

func (p *Education) GoFindAllMyEducations(db *gorm.DB, uid uint64) (*[]Education, error) {
	var err error
	educations := []Education{}
	err = db.Debug().Model(&Education{}).Where("user_id = ?", uid).Limit(100).Find(&educations).Error
	if err != nil {
		return &[]Education{}, err
	}
	// if len(educations) > 0 {
	// 	for i, _ := range educations {
	// 		log.Println(educations[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", educations[i].UserID).Take(&educations[i].User).Error
	// 		if err != nil {
	// 			return &[]Education{}, err
	// 		}
	// 	}
	// }
	return &educations, nil
}

func (p *Education) GoFindEducationByID(db *gorm.DB, pid uint64, uid uint64) (*Education, error) {
	var err error
	err = db.Debug().Model(&Education{}).Where("id = ?", pid).Where("user_id = ?", uid).Take(&p).Error
	if err != nil {
		return &Education{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Education{}, err
	// 	}
	// }
	return p, nil
}

func (p *Education) FindAllMyEducations(db *gorm.DB, uid uint32) (*[]Education, error) {
	var err error
	educations := []Education{}
	err = db.Debug().Model(&Education{}).Where("user_id = ?", uid).Limit(100).Find(&educations).Error
	if err != nil {
		return &[]Education{}, err
	}
	// if len(educations) > 0 {
	// 	for i, _ := range educations {
	// 		log.Println(educations[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", educations[i].UserID).Take(&educations[i].User).Error
	// 		if err != nil {
	// 			return &[]Education{}, err
	// 		}
	// 	}
	// }
	return &educations, nil
}

func (p *Education) FindEducationByID(db *gorm.DB, pid uint64) (*Education, error) {
	var err error
	err = db.Debug().Model(&Education{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Education{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Education{}, err
	// 	}
	// }
	return p, nil
}

func (p *Education) UpdateAEducation(db *gorm.DB) (*Education, error) {

	var err error

	err = db.Debug().Model(&Education{}).Where("id = ?", p.ID).Updates(Education{EducationName: p.EducationName, EducationDescription: p.EducationDescription, EducationTitle: p.EducationTitle, EducationIcon: p.EducationIcon, EducationProgress: p.EducationProgress, EducationLinks: p.EducationLinks, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Education{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Education{}, err
	// 	}
	// }
	return p, nil
}

func (p *Education) DeleteAEducation(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Education{}).Where("id = ? and user_id = ?", pid, uid).Take(&Education{}).Delete(&Education{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Education not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
