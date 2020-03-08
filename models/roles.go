package models

//Scopes struct for scopes
type Scopes struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Edit   bool `json:"edit"`
	Delete bool `json:"delete"`
}

//Roles struct for roles user
type Roles struct {
	RoleType string `json:"roleType"`
	Scope    Scopes `json:"scope"`
}
