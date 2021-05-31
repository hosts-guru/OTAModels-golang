package otamodel

type EnumString string

type EnumStringBase struct {
	Extension *string     `json:"extension,omitempty"`
	Value     *EnumString `json:"value,omitempty"`
}
