package dexmodels

// Entry represents a single Entry.
type Entry struct {
	ID                  int    `yaml:"id"`
	CategoryID          int    `yaml:"category_id"`
	LocationID          int    `yaml:"location_id"`
	Status              int    `yaml:"status"` // 0=Inactive, 1=Active
	LastEditBy          string `yaml:"private_last_edit_by"`
	PrivateAdminNotes   string `yaml:"private_admin_notes"`
	PrivateContactEmail string `yaml:"private_contact_email"`
	EntryCommon
}

type EntryCommon struct {
	Salutation  string `yaml:"salutation"`
	Firstname   string `yaml:"firstname"`
	Lastname    string `yaml:"lastname"`
	Credentials string `yaml:"credentials"`
	JobTitle    string `yaml:"job_title"`
	ShortDesc   string `yaml:"short_desc"` // MarkDown (1 line)
	Phone       string `yaml:"phone"`
	Fax         string `yaml:"fax"`
	Address     string `yaml:"address"`
	Email       string `yaml:"email"`
	Email2      string `yaml:"email2"`
	Website     string `yaml:"website"`
	Website2    string `yaml:"website2"`
	Fees        string `yaml:"fees"`        // MarkDown
	Description string `yaml:"description"` // MarkDown
	Company     string `yaml:"company_name"`
}
