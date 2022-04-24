package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Hobby struct {
	ID               uint64    `gorm:"primary_key;auto_increment" json:"id"`
	HobbyName        string    `gorm:"size:255;UNIQUE_INDEX:hobbyindex;" json:"hobby_name"`
	HobbyTitle       string    `gorm:"size:255;" json:"hobby_title"`
	HobbyDescription string    `gorm:"size:255;" json:"hobby_description"`
	HobbyIcon        string    `gorm:"size:255;" json:"hobby_icon"`
	HobbyLinks       string    `gorm:"size:255;" json:"hobby_links"`
	User             User      `json:"-"`
	UserID           uint32    `gorm:"UNIQUE_INDEX:hobbyindex;not null" json:"user_id"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Hobby) Prepare() {
	p.ID = 0
	p.HobbyName = html.EscapeString(strings.TrimSpace(p.HobbyName))
	p.HobbyTitle = html.EscapeString(strings.TrimSpace(p.HobbyTitle))
	p.HobbyDescription = html.EscapeString(strings.TrimSpace(p.HobbyDescription))
	p.HobbyIcon = html.EscapeString(strings.TrimSpace(p.HobbyIcon))
	p.HobbyLinks = html.EscapeString(strings.TrimSpace(p.HobbyLinks))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Hobby) Validate() error {
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Hobby) SaveHobby(db *gorm.DB) (*Hobby, error) {
	var err error
	err = db.Debug().Model(&Hobby{}).Create(&p).Error
	if err != nil {
		return &Hobby{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Hobby{}, err
	// 	}
	// }
	return p, nil
}

func (p *Hobby) FindAllHobbys(db *gorm.DB) (*[]Hobby, error) {
	var err error
	hobbys := []Hobby{}
	err = db.Debug().Model(&Hobby{}).Find(&hobbys).Error
	if err != nil {
		return &[]Hobby{}, err
	}
	// if len(hobbys) > 0 {
	// 	for i, _ := range hobbys {
	// 		log.Println(hobbys[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", hobbys[i].UserID).Find(&hobbys[i].User).Error
	// 		if err != nil {
	// 			return &[]Hobby{}, err
	// 		}
	// 	}
	// }
	return &hobbys, nil
}

func (p *Hobby) GoFindAllMyHobbys(db *gorm.DB, uid uint64) (*[]Hobby, error) {
	var err error
	hobbys := []Hobby{}
	err = db.Debug().Model(&Hobby{}).Where("user_id = ?", uid).Limit(100).Find(&hobbys).Error
	if err != nil {
		return &[]Hobby{}, err
	}
	// if len(hobbys) > 0 {
	// 	for i, _ := range hobbys {
	// 		log.Println(hobbys[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", hobbys[i].UserID).Take(&hobbys[i].User).Error
	// 		if err != nil {
	// 			return &[]Hobby{}, err
	// 		}
	// 	}
	// }
	return &hobbys, nil
}

func (p *Hobby) GoFindHobbyByID(db *gorm.DB, pid uint64, uid uint64) (*Hobby, error) {
	var err error
	err = db.Debug().Model(&Hobby{}).Where("id = ?", pid).Where("user_id = ?", uid).Take(&p).Error
	if err != nil {
		return &Hobby{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Hobby{}, err
	// 	}
	// }
	return p, nil
}

func (p *Hobby) FindAllMyHobbys(db *gorm.DB, uid uint32) (*[]Hobby, error) {
	var err error
	hobbys := []Hobby{}
	err = db.Debug().Model(&Hobby{}).Where("user_id = ?", uid).Limit(100).Find(&hobbys).Error
	if err != nil {
		return &[]Hobby{}, err
	}
	// if len(hobbys) > 0 {
	// 	for i, _ := range hobbys {
	// 		log.Println(hobbys[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", hobbys[i].UserID).Take(&hobbys[i].User).Error
	// 		if err != nil {
	// 			return &[]Hobby{}, err
	// 		}
	// 	}
	// }
	return &hobbys, nil
}

func (p *Hobby) FindHobbyByID(db *gorm.DB, pid uint64) (*Hobby, error) {
	var err error
	err = db.Debug().Model(&Hobby{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Hobby{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Hobby{}, err
	// 	}
	// }
	return p, nil
}

func (p *Hobby) UpdateAHobby(db *gorm.DB) (*Hobby, error) {

	var err error

	err = db.Debug().Model(&Hobby{}).Where("id = ?", p.ID).Updates(Hobby{HobbyName: p.HobbyName, HobbyDescription: p.HobbyDescription, HobbyTitle: p.HobbyTitle, HobbyIcon: p.HobbyIcon, HobbyLinks: p.HobbyLinks, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Hobby{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Hobby{}, err
	// 	}
	// }
	return p, nil
}

func (p *Hobby) DeleteAHobby(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Hobby{}).Where("id = ? and user_id = ?", pid, uid).Take(&Hobby{}).Delete(&Hobby{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Hobby not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
