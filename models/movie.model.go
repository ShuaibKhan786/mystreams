package models

type Quality string

const (
	HD  Quality = "hd"
	CAM Quality = "cam"
)

type Movie struct {
	ID          *int64   `form:"id,omitempty"`
	Title       *string  `form:"title,omitempty"`
	Description *string  `form:"description,omitempty"`
	Year        *string  `form:"year,omitempty"`
	Quality     *Quality `form:"quality,omitempty"`
	RuntimeMins *uint    `form:"runtime,omitempty"`
	Country     *string  `form:"country,omitempty"`

	Genres []*Genre `json:"genres,omitempty"`
	Roles  []*Role  `json:"roles,omitempty"`
}

type RoleType string

const (
	CastRoleType     RoleType = "cast"
	DirectorRoleType RoleType = "director"
	ProducerRoleType RoleType = "producer"
	WriterRoleType   RoleType = "writer"
	CrewRoleType     RoleType = "crew"
)

type Role struct {
	Person
	RoleType  *RoleType `json:"role_type"`
	Character *string   `json:"character"`
}
