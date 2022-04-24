package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProfessionalExp struct {
	ID                     uint64    `gorm:"primary_key;auto_increment" json:"id"`
	InstitutionName        string    `gorm:"size:255;UNIQUE_INDEX:professionalexpindex;" json:"institution_name"`
	InstitutionTitle       string    `gorm:"size:255;" json:"institution_title"`
	InstitutionDescription string    `gorm:"size:255;" json:"institution_description"`
	InstitutionLink        string    `gorm:"size:255;" json:"institution_link"`
	ResponsibilityLevel    string    `gorm:"size:255;" json:"responsibility_level"`
	Responsibilities       string    `gorm:"size:255;" json:"responsibilities"`
	JobExperienceFrom      string    `gorm:"size:255;" json:"job_experience_from"`
	JobExperienceTo        string    `gorm:"size:255;" json:"job_experience_to"`
	JobMotivation          string    `gorm:"size:255;" json:"Job_motivation"`
	User                   User      `json:"-"`
	UserID                 uint32    `gorm:"UNIQUE_INDEX:professionalexpindex;not null" json:"user_id"`
	CreatedAt              time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt              time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *ProfessionalExp) Prepare() {
	p.ID = 0
	p.InstitutionName = html.EscapeString(strings.TrimSpace(p.InstitutionName))
	p.InstitutionTitle = html.EscapeString(strings.TrimSpace(p.InstitutionTitle))
	p.InstitutionDescription = html.EscapeString(strings.TrimSpace(p.InstitutionDescription))
	p.InstitutionLink = html.EscapeString(strings.TrimSpace(p.InstitutionLink))
	p.ResponsibilityLevel = html.EscapeString(strings.TrimSpace(p.ResponsibilityLevel))
	p.Responsibilities = html.EscapeString(strings.TrimSpace(p.Responsibilities))
	p.JobExperienceFrom = html.EscapeString(strings.TrimSpace(p.JobExperienceFrom))
	p.JobExperienceTo = html.EscapeString(strings.TrimSpace(p.JobExperienceTo))
	p.JobMotivation = html.EscapeString(strings.TrimSpace(p.JobMotivation))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *ProfessionalExp) Validate() error {
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *ProfessionalExp) SaveProfessionalExp(db *gorm.DB) (*ProfessionalExp, error) {
	var err error
	err = db.Debug().Model(&ProfessionalExp{}).Create(&p).Error
	if err != nil {
		return &ProfessionalExp{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &ProfessionalExp{}, err
	// 	}
	// }
	return p, nil
}

func (p *ProfessionalExp) FindAllProfessionalExps(db *gorm.DB) (*[]ProfessionalExp, error) {
	var err error
	professional_exps := []ProfessionalExp{}
	err = db.Debug().Model(&ProfessionalExp{}).Find(&professional_exps).Error
	if err != nil {
		return &[]ProfessionalExp{}, err
	}
	// if len(professional_exps) > 0 {
	// 	for i, _ := range professional_exps {
	// 		log.Println(professional_exps[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", professional_exps[i].UserID).Find(&professional_exps[i].User).Error
	// 		if err != nil {
	// 			return &[]ProfessionalExp{}, err
	// 		}
	// 	}
	// }
	return &professional_exps, nil
}

func (p *ProfessionalExp) GoFindAllMyProfessionalExps(db *gorm.DB, uid uint64) (*[]ProfessionalExp, error) {
	var err error
	professional_exps := []ProfessionalExp{}
	err = db.Debug().Model(&ProfessionalExp{}).Where("user_id = ?", uid).Limit(100).Find(&professional_exps).Error
	if err != nil {
		return &[]ProfessionalExp{}, err
	}
	// if len(professional_exps) > 0 {
	// 	for i, _ := range professional_exps {
	// 		log.Println(professional_exps[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", professional_exps[i].UserID).Take(&professional_exps[i].User).Error
	// 		if err != nil {
	// 			return &[]ProfessionalExp{}, err
	// 		}
	// 	}
	// }
	return &professional_exps, nil
}

func (p *ProfessionalExp) GoFindProfessionalExpByID(db *gorm.DB, pid uint64, uid uint64) (*ProfessionalExp, error) {
	var err error
	err = db.Debug().Model(&ProfessionalExp{}).Where("id = ?", pid).Where("user_id = ?", uid).Take(&p).Error
	if err != nil {
		return &ProfessionalExp{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &ProfessionalExp{}, err
	// 	}
	// }
	return p, nil
}

func (p *ProfessionalExp) FindAllMyProfessionalExps(db *gorm.DB, uid uint32) (*[]ProfessionalExp, error) {
	var err error
	professional_exps := []ProfessionalExp{}
	err = db.Debug().Model(&ProfessionalExp{}).Where("user_id = ?", uid).Limit(100).Find(&professional_exps).Error
	if err != nil {
		return &[]ProfessionalExp{}, err
	}
	// if len(professional_exps) > 0 {
	// 	for i, _ := range professional_exps {
	// 		log.Println(professional_exps[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", professional_exps[i].UserID).Take(&professional_exps[i].User).Error
	// 		if err != nil {
	// 			return &[]ProfessionalExp{}, err
	// 		}
	// 	}
	// }
	return &professional_exps, nil
}

func (p *ProfessionalExp) FindProfessionalExpByID(db *gorm.DB, pid uint64) (*ProfessionalExp, error) {
	var err error
	err = db.Debug().Model(&ProfessionalExp{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProfessionalExp{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &ProfessionalExp{}, err
	// 	}
	// }
	return p, nil
}

func (p *ProfessionalExp) UpdateAProfessionalExp(db *gorm.DB) (*ProfessionalExp, error) {

	var err error

	err = db.Debug().Model(&ProfessionalExp{}).Where("id = ?", p.ID).Updates(ProfessionalExp{InstitutionName: p.InstitutionName, InstitutionDescription: p.InstitutionDescription, InstitutionTitle: p.InstitutionTitle, InstitutionLink: p.InstitutionLink, ResponsibilityLevel: p.ResponsibilityLevel, Responsibilities: p.Responsibilities, JobExperienceFrom: p.JobExperienceFrom, JobExperienceTo: p.JobExperienceTo, JobMotivation: p.JobMotivation, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &ProfessionalExp{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &ProfessionalExp{}, err
	// 	}
	// }
	return p, nil
}

func (p *ProfessionalExp) DeleteAProfessionalExp(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&ProfessionalExp{}).Where("id = ? and user_id = ?", pid, uid).Take(&ProfessionalExp{}).Delete(&ProfessionalExp{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProfessionalExp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
