package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(db *gorm.DB) {
	DB = db
}

type Accounts struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	MetamaskAddress string     `gorm:"unique;not null;size:42;index" json:"metamask_address"`
	AccountType     string     `gorm:"not null" json:"account_type"`
	Verified        bool       `gorm:"default:false" json:"verified"`
	Credentials     int        `gorm:"default:0" json:"credentials"`
	LastLoginAt     *time.Time `json:"last_login_at"`
	IsActive        bool       `gorm:"default:true" json:"is_active"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	OwnerID   uint   `gorm:"not null"` 
    OwnerType string `gorm:"not null"` 
}

type Users struct {
	ID             		uint   `gorm:"primaryKey" json:"id"`
	MetamaskAddress 	string `gorm:"unique;size:42" json:"metamask_address"`
	Email               string `gorm:"not null;unique" json:"email"`
	FirstName           string `gorm:"not null" json:"first_name"`
	LastName            string `gorm:"not null" json:"last_name"`
	StudentEmail        string `gorm:"size:100" json:"student_id"`
	IsVerified 			bool   `gorm:"default:false" json:"is_verified"`
	
    Account Accounts `gorm:"polymorphic:Owner;"`

	PendingRequests []PendingRequest `gorm:"foreignKey:RequesterID"`

}

type Organization struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	MetamaskAddress string `gorm:"unique;size:42" json:"metamask_address"`
	AcadEmail       string `gorm:"not null;unique" json:"acad_email"`
	OrgName         string `gorm:"not null" json:"org_name"`
	OrgType         string `gorm:"size:100" json:"org_type"`
	OrgUrl          string `gorm:"size:255" json:"org_url"`
	OrgDesc         string `gorm:"type:text" json:"org_desc"`
	Country         string `gorm:"size:100" json:"country"`
	State           string `gorm:"size:100" json:"state"`
	City            string `gorm:"size:100" json:"city"`
	Address         string `gorm:"type:text" json:"address"`
	PostalCode      string `gorm:"size:20" json:"postal_code"`
	IsVerified 		bool   `gorm:"default:false" json:"is_verified"`

	TotalStudents     int `gorm:"default:0" json:"total_students"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    Account Accounts `gorm:"polymorphic:Owner;"`
    Credentials  []Credential `gorm:"foreignKey:OrganizationID"`

	PendingRequests []PendingRequest `gorm:"foreignKey:OrganizationID"`
}

type PendingRequest struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	RequesterID    uint      `gorm:"not null" json:"requester_id"`
	OrganizationID uint      `gorm:"not null" json:"organization_id"`
	IsApproved     bool      `gorm:"default:false" json:"is_approved"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Requester    Users        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"requester"`
	Organization Organization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"organization"`
}


type Credential struct {
	ID 				 string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`	
	DegreeID        *uint     `gorm:"column:degree_id;default:NULL" json:"degree_id"`
	StudentWallet    string    `gorm:"not null;size:42" json:"student_wallet"`
	UniversityWallet string    `gorm:"not null;size:42" json:"university_wallet"`
	DegreeName      string    `gorm:"not null;size:255" json:"degree_name"`
	Description     string    `gorm:"type:text" json:"description"`
	Type            string    `gorm:"not null;size:100" json:"type"`
	Major           string    `gorm:"size:255" json:"major"`
	IssuedDate      time.Time `gorm:"not null" json:"issued_date"`
	GraduationDate  string    `gorm:"size:50" json:"graduation_date"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	IPFSLink 		string 	  `gorm:"not null" json:"ipfs_link"`
	DeanSig 		string 	  `gorm:"not null" json:"dean_sig"`

	UserID         uint         `json:"user_id"`
	User           Users        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`

	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"organization"`
}

type Transaction struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TxHash        string    `gorm:"not null;size:66" json:"tx_hash"`
	BlockNumber   string    `gorm:"not null" json:"block_number"`
	From          string    `gorm:"not null;size:42" json:"from"`
	To            string    `gorm:"not null;size:42" json:"to"`
	ValueETH      string   `gorm:"not null" json:"value_eth"`
	Gas           string    `gorm:"not null" json:"gas"`
	GasPrice      string    `gorm:"not null" json:"gas_price"`
	Timestamp     time.Time `gorm:"not null" json:"timestamp"`
}
