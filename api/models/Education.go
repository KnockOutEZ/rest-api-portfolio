package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Education struct {
	ID                     uint64    `gorm:"primary_key;auto_increment" json:"id"`
	InstitutionName        string    `gorm:"size:255;UNIQUE_INDEX:educationindex;" json:"institution_name"`
	InstitutionTitle       string    `gorm:"size:255;" json:"institution_title"`
	InstitutionDescription string    `gorm:"size:255;" json:"institution_description"`
	InstitutionIcon        string    `gorm:"size:255;" json:"institution_icon"`
	InstitutionProgress    string    `gorm:"size:255;" json:"institution_progress"`
	InstitutionLinks       string    `gorm:"size:255;" json:"institution_links"`
	EducationLevel         string    `gorm:"size:255;" json:"education_level"`
	EducationPeriodFrom    string    `gorm:"size:255;" json:"education_period_from"`
	EducationPeriodTo      string    `gorm:"size:255;" json:"education_period_to"`
	StudyMotivation        string    `gorm:"size:255;" json:"study_motivation"`
	User                   User      `json:"-"`
	UserID                 uint32    `gorm:"UNIQUE_INDEX:educationindex;not null" json:"user_id"`
	CreatedAt              time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt              time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Education) Prepare() {
	p.ID = 0
	p.InstitutionName = html.EscapeString(strings.TrimSpace(p.InstitutionName))
	p.InstitutionTitle = html.EscapeString(strings.TrimSpace(p.InstitutionTitle))
	p.InstitutionDescription = html.EscapeString(strings.TrimSpace(p.InstitutionDescription))
	p.InstitutionIcon = html.EscapeString(strings.TrimSpace(p.InstitutionIcon))
	p.InstitutionProgress = html.EscapeString(strings.TrimSpace(p.InstitutionProgress))
	p.InstitutionLinks = html.EscapeString(strings.TrimSpace(p.InstitutionLinks))
	p.EducationLevel = html.EscapeString(strings.TrimSpace(p.EducationLevel))
	p.EducationPeriodFrom = html.EscapeString(strings.TrimSpace(p.EducationPeriodFrom))
	p.EducationPeriodTo = html.EscapeString(strings.TrimSpace(p.EducationPeriodTo))
	p.StudyMotivation = html.EscapeString(strings.TrimSpace(p.StudyMotivation))
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

	err = db.Debug().Model(&Education{}).Where("id = ?", p.ID).Updates(Education{
		InstitutionName:p.InstitutionName,
		InstitutionTitle:p.InstitutionTitle,
		InstitutionDescription:p.InstitutionDescription,
		InstitutionIcon:p.InstitutionIcon,
		InstitutionProgress:p.InstitutionProgress,
		InstitutionLinks:p.InstitutionLinks,
		EducationLevel:p.EducationLevel,
		EducationPeriodFrom:p.EducationPeriodFrom,
		EducationPeriodTo:p.EducationPeriodTo,
		StudyMotivation:p.StudyMotivation ,
		UpdatedAt: time.Now()}).Error
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
