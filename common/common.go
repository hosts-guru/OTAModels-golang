package common

import "github.com/hosts-guru/OTAModels-golang"

type ImageID struct {
	Type                  string                  `json:"@type"`
	Sort                  int                     `json:"sort"`
	ObjID                 string                  `json:"objID"`
	URL                   string                  `json:"URL,omitempty"` // URL of the multimedia item for a specific format.
	MultimediaDescription []MultimediaDescription `json:"MultimediaDescription"`
}

// Description Assigned Type: ota2-0500:Description
type Description struct {
	Type              string              `json:"@type"`
	DateTimeStamp     *DateTimeStamp      `json:"DateTimeStamp"` // Assigned Type: ota2-0500:DateTimeStamp
	ExtensionPoint    interface{}         `json:"ExtensionPoint"`
	Image             []ImageID           `json:"Image,omitempty"`
	Name              string              `json:"name,omitempty"`   // Assigned Type: ota2-0500:StringTiny
	Number            int64               `json:"number,omitempty"` // The sequence number for the paragraph.
	ParagraphListItem []ParagraphListItem `json:"ParagraphListItem,omitempty"`
	Text              []Text              `json:"Text,omitempty"`
	URL               []string            `json:"URL,omitempty"`
}

type MultimediaDescription struct {
	Type     string `json:"@type"`
	Title    string `json:"title"`
	Language string `json:"language"`
	Caption  string `json:"caption"`
}

// Text Provides text and indicates whether it is formatted or not.
type Text struct {
	Language   string               `json:"language,omitempty"`   // The language in which the text is provided.
	TextFormat *otamodel.EnumString `json:"textFormat,omitempty"` // Indicates the format of text used in the description e.g. unformatted  or html.
	Value      string               `json:"value,omitempty"`
}

type ParagraphListItem struct {
	Language   string               `json:"language,omitempty"`   // The language in which the text is provided.
	ListItem   int64                `json:"listItem,omitempty"`   // The item or sequence number.
	Text2      string               `json:"text2,omitempty"`      // Assigned Type: ota2-0500:Text
	TextFormat *otamodel.EnumString `json:"textFormat,omitempty"` // Indicates the format of text used in the description e.g. unformatted  or html.
}

// Identifier Assigned Type: ota2-0500:Identifier
//
// Identifier provides the ability to create a globally unique identifier.   This could be
// GUID a UUID or any other string that an organization that distinctly identifies the
// specific object.\n\nIdeally, for the identifier to be globally unique it must have a
// system provided identifier and the system must be identified using a global naming
// authority. System identification uses the domain naming system (DNS) to assure they are
// globally unique and should be an URL. The system provided ID will typically be a primary
// or surrogate key in a database.\n\nThe URL, system and company attributes can be omitted
// only when the system context can be implied by a parent or ancestor element.
type Identifier struct {
	Organization string `json:"organization,omitempty"` // Assigned Type: ota2-0500:StringShort
	System       string `json:"system,omitempty"`       // Assigned Type: ota2-0500:StringShort
	URL          string `json:"url,omitempty"`          // Uri of the creating system. The URI should be an URL and be globally unique. Should only; be omitted when the context is clearly implied for the containing document.
	Value        string `json:"value,omitempty"`
}

// DateTimeStamp Assigned Type: ota2-0500:DateTimeStamp
//
// Time stamp of the creation.
type DateTimeStamp struct {
	Create         string `json:"create,omitempty"`         // Time stamp of the creation.
	CreatorID      string `json:"creatorID,omitempty"`      // Assigned Type: ota2-0500:StringShort
	LastModifierID string `json:"lastModifierID,omitempty"` // Assigned Type: ota2-0500:StringShort
	LastModify     string `json:"lastModify,omitempty"`     // Time stamp of last modification.
	Purge          string `json:"purge,omitempty"`          // Date an item will be purged from a database (e.g., from a live database to an archive).
	Value          string `json:"value,omitempty"`
}

type Amenity struct {
	Type           string                   `json:"@type"`
	Code           *Code                    `json:"Code"`                  // Assigned Type: ota2-0500:Code
	Description    string                   `json:"description,omitempty"` // Assigned Type: ota2-0500:StringShort
	ExtensionPoint interface{}              `json:"ExtensionPoint"`
	Name           string                   `json:"name,omitempty"` // Assigned Type: ota2-0500:String
	Proximity      *otamodel.EnumStringBase `json:"Proximity"`      // Assigned Type: ota2-0400:Proximity_Enum
}

// Code Assigned Type: ota2-0500:Code
//
// Any code used to specify an item, for example, type of traveler, service code, room
// amenity, etc.  May be used to reference the OpenTravel RMA Codelist.
type Code struct {
	CodeContext string `json:"codeContext,omitempty"` // Assigned Type: ota2-0500:StringTiny
	Description string `json:"description,omitempty"` // Assigned Type: ota2-0500:String
	ListURI     string `json:"listURI,omitempty"`     // Assigned Type: ota2-0500:String
	Quantity    int64  `json:"quantity,omitempty"`    // Assigned Type: ota2-0500:NonNegativeInteger
	Value       string `json:"value,omitempty"`
}
