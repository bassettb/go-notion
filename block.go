package notion

import (
	"encoding/json"
	"time"
)

// Block represents content on the Notion platform.
// See: https://developers.notion.com/reference/block
type Block struct {
	Object         string     `json:"object"`
	ID             string     `json:"id,omitempty"`
	Type           BlockType  `json:"type"`
	CreatedTime    *time.Time `json:"created_time,omitempty"`
	LastEditedTime *time.Time `json:"last_edited_time,omitempty"`
	HasChildren    bool       `json:"has_children,omitempty"`

	Paragraph        *RichTextBlock `json:"paragraph,omitempty"`
	Heading1         *Heading       `json:"heading_1,omitempty"`
	Heading2         *Heading       `json:"heading_2,omitempty"`
	Heading3         *Heading       `json:"heading_3,omitempty"`
	BulletedListItem *RichTextBlock `json:"bulleted_list_item,omitempty"`
	NumberedListItem *RichTextBlock `json:"numbered_list_item,omitempty"`
	ToDo             *ToDo          `json:"to_do,omitempty"`
	Toggle           *RichTextBlock `json:"toggle,omitempty"`
	ChildPage        *ChildPage     `json:"child_page,omitempty"`
	ChildDatabase    *ChildDatabase `json:"child_database,omitempty"`
	Callout          *Callout       `json:"callout,omitempty"`
	Quote            *RichTextBlock `json:"quote,omitempty"`
	Code             *Code          `json:"code,omitempty"`
	Embed            *Embed         `json:"embed,omitempty"`
	Image            *FileBlock     `json:"image,omitempty"`
	Video            *FileBlock     `json:"video,omitempty"`
	File             *FileBlock     `json:"file,omitempty"`
	PDF              *FileBlock     `json:"pdf,omitempty"`
	Bookmark         *Bookmark      `json:"bookmark,omitempty"`
	Equation         *Equation      `json:"equation,omitempty"`
}

type RichTextBlock struct {
	Text     []RichText `json:"text"`
	Children []Block    `json:"children,omitempty"`
}

type Heading struct {
	Text []RichText `json:"text"`
}

type ToDo struct {
	RichTextBlock
	Checked *bool `json:"checked,omitempty"`
}

type ChildPage struct {
	Title string `json:"title"`
}

type ChildDatabase struct {
	Title string `json:"title"`
}

type Callout struct {
	RichTextBlock
	Icon *Icon `json:"icon,omitempty"`
}

type Code struct {
	RichTextBlock
	Language *string `json:"language,omitempty"`
}

type Embed struct {
	URL string `json:"url"`
}

type FileBlock struct {
	Type FileType `json:"type"`

	File     *FileFile     `json:"file,omitempty"`
	External *FileExternal `json:"external,omitempty"`
	Caption  []RichText    `json:"caption,omitempty"`
}

type FileFile struct {
	URL        string   `json:"url"`
	ExpiryTime DateTime `json:"expiry_time"`
}

type FileExternal struct {
	URL string `json:"url"`
}

type FileType string

const (
	FileTypeFile     FileType = "file"
	FileTypeExternal FileType = "external"
)

type Bookmark struct {
	URL     string     `json:"url"`
	Caption []RichText `json:"caption,omitempty"`
}

type BlockType string

const (
	BlockTypeParagraph        BlockType = "paragraph"
	BlockTypeHeading1         BlockType = "heading_1"
	BlockTypeHeading2         BlockType = "heading_2"
	BlockTypeHeading3         BlockType = "heading_3"
	BlockTypeBulletedListItem BlockType = "bulleted_list_item"
	BlockTypeNumberedListItem BlockType = "numbered_list_item"
	BlockTypeToDo             BlockType = "to_do"
	BlockTypeToggle           BlockType = "toggle"
	BlockTypeChildPage        BlockType = "child_page"
	BlockTypeChildDatabase    BlockType = "child_database"
	BlockTypeCallout          BlockType = "callout"
	BlockTypeQuote            BlockType = "quote"
	BlockTypeCode             BlockType = "code"
	BlockTypeEmbed            BlockType = "embed"
	BlockTypeImage            BlockType = "image"
	BlockTypeVideo            BlockType = "video"
	BlockTypeFile             BlockType = "file"
	BlockTypePDF              BlockType = "pdf"
	BlockTypeBookmark         BlockType = "bookmark"
	BlockTypeEquation         BlockType = "equation"
	BlockTypeUnsupported      BlockType = "unsupported"
)

type PaginationQuery struct {
	StartCursor string
	PageSize    int
}

// BlockChildrenResponse contains results (block children) and pagination data returned from a find request.
type BlockChildrenResponse struct {
	Results    []Block `json:"results"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor"`
}

// MarshalJSON implements json.Marshaler.
func (b Block) MarshalJSON() ([]byte, error) {
	type blockAlias Block

	alias := blockAlias(b)
	alias.Object = "block"

	return json.Marshal(alias)
}
