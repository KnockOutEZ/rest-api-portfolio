package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Project struct {
	ID                 uint64    `gorm:"primary_key;auto_increment" json:"id"`
	ProjectName        string    `gorm:"size:255;UNIQUE_INDEX:projectindex;" json:"project_name"`
	ProjectTitle       string    `gorm:"size:255;" json:"project_title"`
	ProjectDescription string    `gorm:"size:255;" json:"project_description"`
	ProjectIcon        string    `gorm:"size:255;" json:"project_icon"`
	ProjectImgs        string    `gorm:"size:255;" json:"project_imgs"`
	ProjectLinks       string    `gorm:"size:255;" json:"project_links"`
	ProjectSkillArea   string    `gorm:"size:255;" json:"project_skill_area"`
	ProjectSkills      string    `gorm:"size:255;" json:"project_skills"`
	ProjectTimeFrom    string    `gorm:"size:255;" json:"project__time_from"`
	ProjectTimeTo      string    `gorm:"size:255;" json:"project_time_to"`
	ProjectClient      string    `gorm:"size:255;" json:"project__client"`
	User               User      `json:"-"`
	UserID             uint32    `gorm:"UNIQUE_INDEX:projectindex;not null" json:"user_id"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Project) Prepare() {
	p.ID = 0
	p.ProjectName = html.EscapeString(strings.TrimSpace(p.ProjectName))
	p.ProjectTitle = html.EscapeString(strings.TrimSpace(p.ProjectTitle))
	p.ProjectDescription = html.EscapeString(strings.TrimSpace(p.ProjectDescription))
	p.ProjectIcon = html.EscapeString(strings.TrimSpace(p.ProjectIcon))
	p.ProjectImgs = html.EscapeString(strings.TrimSpace(p.ProjectImgs))
	p.ProjectLinks = html.EscapeString(strings.TrimSpace(p.ProjectLinks))
	p.ProjectSkillArea = html.EscapeString(strings.TrimSpace(p.ProjectSkillArea))
	p.ProjectSkills = html.EscapeString(strings.TrimSpace(p.ProjectSkills))
	p.ProjectTimeFrom = html.EscapeString(strings.TrimSpace(p.ProjectTimeFrom))
	p.ProjectTimeTo = html.EscapeString(strings.TrimSpace(p.ProjectTimeTo))
	p.ProjectClient = html.EscapeString(strings.TrimSpace(p.ProjectClient))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Project) Validate() error {
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Project) SaveProject(db *gorm.DB) (*Project, error) {
	var err error
	err = db.Debug().Model(&Project{}).Create(&p).Error
	if err != nil {
		return &Project{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Project{}, err
	// 	}
	// }
	return p, nil
}

func (p *Project) FindAllProjects(db *gorm.DB) (*[]Project, error) {
	var err error
	projects := []Project{}
	err = db.Debug().Model(&Project{}).Find(&projects).Error
	if err != nil {
		return &[]Project{}, err
	}
	// if len(projects) > 0 {
	// 	for i, _ := range projects {
	// 		log.Println(projects[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", projects[i].UserID).Find(&projects[i].User).Error
	// 		if err != nil {
	// 			return &[]Project{}, err
	// 		}
	// 	}
	// }
	return &projects, nil
}

func (p *Project) GoFindAllMyProjects(db *gorm.DB, uid uint64) (*[]Project, error) {
	var err error
	projects := []Project{}
	err = db.Debug().Model(&Project{}).Where("user_id = ?", uid).Limit(100).Find(&projects).Error
	if err != nil {
		return &[]Project{}, err
	}
	// if len(projects) > 0 {
	// 	for i, _ := range projects {
	// 		log.Println(projects[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", projects[i].UserID).Take(&projects[i].User).Error
	// 		if err != nil {
	// 			return &[]Project{}, err
	// 		}
	// 	}
	// }
	return &projects, nil
}

func (p *Project) GoFindProjectByID(db *gorm.DB, pid uint64, uid uint64) (*Project, error) {
	var err error
	err = db.Debug().Model(&Project{}).Where("id = ?", pid).Where("user_id = ?", uid).Take(&p).Error
	if err != nil {
		return &Project{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Project{}, err
	// 	}
	// }
	return p, nil
}

func (p *Project) FindAllMyProjects(db *gorm.DB, uid uint32) (*[]Project, error) {
	var err error
	projects := []Project{}
	err = db.Debug().Model(&Project{}).Where("user_id = ?", uid).Limit(100).Find(&projects).Error
	if err != nil {
		return &[]Project{}, err
	}
	// if len(projects) > 0 {
	// 	for i, _ := range projects {
	// 		log.Println(projects[i].UserID)
	// 		err := db.Debug().Model(&User{}).Where("id = ?", projects[i].UserID).Take(&projects[i].User).Error
	// 		if err != nil {
	// 			return &[]Project{}, err
	// 		}
	// 	}
	// }
	return &projects, nil
}

func (p *Project) FindProjectByID(db *gorm.DB, pid uint64) (*Project, error) {
	var err error
	err = db.Debug().Model(&Project{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Project{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Project{}, err
	// 	}
	// }
	return p, nil
}

func (p *Project) UpdateAProject(db *gorm.DB) (*Project, error) {

	var err error
	err = db.Debug().Model(&Project{}).Where("id = ?", p.ID).Updates(Project{ProjectName: p.ProjectName, ProjectDescription: p.ProjectDescription, ProjectTitle: p.ProjectTitle, ProjectIcon: p.ProjectIcon, ProjectImgs: p.ProjectImgs, ProjectLinks: p.ProjectLinks, ProjectSkillArea: p.ProjectSkillArea, ProjectSkills: p.ProjectSkills, ProjectTimeFrom: p.ProjectTimeFrom, ProjectTimeTo: p.ProjectTimeTo, ProjectClient: p.ProjectClient, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Project{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	// 	if err != nil {
	// 		return &Project{}, err
	// 	}
	// }
	return p, nil
}

func (p *Project) DeleteAProject(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Project{}).Where("id = ? and user_id = ?", pid, uid).Take(&Project{}).Delete(&Project{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Project not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
