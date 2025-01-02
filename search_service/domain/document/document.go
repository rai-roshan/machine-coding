package document

type Document struct {
	id uint32
	content string
}

func NewDocument(id uint32, content string) *Document {
	return &Document{
		id: id,
		content: content,
	}
}

func (d *Document) GetId() uint32 {
	return d.id
}

func (d *Document) GetContent() string {
	return d.content
}