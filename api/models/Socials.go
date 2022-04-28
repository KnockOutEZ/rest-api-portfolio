package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Socials struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	SocialsName  string    `gorm:"size:255;UNIQUE_INDEX:socialsindex;" json:"social_name"`
	SocialsTitle string    `gorm:"size:255;" json:"social_title"`
	SocialsIcon  string    `gorm:"size:255;" json:"social_icon"`
	SocialsLinks string    `gorm:"size:255;" json:"social_links"`
	SocialsDescription string    `gorm:"size:255;" json:"social_description"`
	User         User      `json:"-"`
	UserID       uint32    `gorm:"UNIQUE_INDEX:socialsindex;not null" json:"user_id"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Socials) Prepare() {
	p.ID = 0
	p.SocialsName = html.EscapeString(strings.TrimSpace(p.SocialsName))
	p.SocialsTitle = html.EscapeString(strings.TrimSpace(p.SocialsTitle))
	p.SocialsIcon = html.EscapeString(strings.TrimSpace(p.SocialsIcon))
	p.SocialsLinks = html.EscapeString(strings.TrimSpace(p.SocialsLinks))
	p.SocialsDescription = html.EscapeString(strings.TrimSpace(p.SocialsDescription))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Socials) Validate() error {
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Socials) SaveSocials(db *gorm.DB) (*Socials, error) {
	var err error
	err = db.Debug().Model(&Socials{}).Create(&p).Error
	if err != nil {
		return &Socials{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Socials{}, err
	// 	}
	// }
	return p, nil
}

func (p *Socials) FindAllSocialss(db *gorm.DB) (*[]Socials, error) {
	var err error
	socials := []Socials{}
	err = db.Debug().Model(&Socials{}).Find(&socials).Error
	if err != nil {
		return &[]Socials{}, err
	}
	// if len(socials) > 0 {
	// 	for i, _ := range socials {
	// 		log.Println(socials[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", socials[i].UserID).Find(&socials[i].User).Error
	// 		if err != nil {
	// 			return &[]Socials{}, err
	// 		}
	// 	}
	// }
	return &socials, nil
}

func (p *Socials) GoFindAllMySocialss(db *gorm.DB, uid uint64) (*[]Socials, error) {
	var err error
	socials := []Socials{}
	err = db.Debug().Model(&Socials{}).Where("user_id = ?", uid).Limit(100).Find(&socials).Error
	if err != nil {
		return &[]Socials{}, err
	}
	// if len(socials) > 0 {
	// 	for i, _ := range socials {
	// 		log.Println(socials[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", socials[i].UserID).Take(&socials[i].User).Error
	// 		if err != nil {
	// 			return &[]Socials{}, err
	// 		}
	// 	}
	// }
	return &socials, nil
}

func (p *Socials) GoFindSocialsByID(db *gorm.DB, pid uint64, uid uint64) (*Socials, error) {
	var err error
	err = db.Debug().Model(&Socials{}).Where("id = ?", pid).Where("user_id = ?", uid).Take(&p).Error
	if err != nil {
		return &Socials{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Socials{}, err
	// 	}
	// }
	return p, nil
}

func (p *Socials) FindAllMySocialss(db *gorm.DB, uid uint32) (*[]Socials, error) {
	var err error
	socials := []Socials{}
	err = db.Debug().Model(&Socials{}).Where("user_id = ?", uid).Limit(100).Find(&socials).Error
	if err != nil {
		return &[]Socials{}, err
	}
	// if len(socials) > 0 {
	// 	for i, _ := range socials {
	// 		log.Println(socials[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", socials[i].UserID).Take(&socials[i].User).Error
	// 		if err != nil {
	// 			return &[]Socials{}, err
	// 		}
	// 	}
	// }
	return &socials, nil
}

func (p *Socials) FindSocialsByID(db *gorm.DB, pid uint64) (*Socials, error) {
	var err error
	err = db.Debug().Model(&Socials{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Socials{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Socials{}, err
	// 	}
	// }
	return p, nil
}

func (p *Socials) UpdateASocials(db *gorm.DB) (*Socials, error) {

	var err error

	err = db.Debug().Model(&Socials{}).Where("id = ?", p.ID).Updates(Socials{SocialsName: p.SocialsName, SocialsTitle: p.SocialsTitle, SocialsIcon: p.SocialsIcon, SocialsLinks: p.SocialsLinks,SocialsDescription: p.SocialsDescription, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Socials{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Socials{}, err
	// 	}
	// }
	return p, nil
}

func (p *Socials) DeleteASocials(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Socials{}).Where("id = ? and user_id = ?", pid, uid).Take(&Socials{}).Delete(&Socials{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Socials not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
