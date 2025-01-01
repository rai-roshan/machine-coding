package document

type Document struct {
	id uint32
	name string
	content string
}

func (d *Document) GetDocumentId() uint32 {
	return d.id
}

func (d *Document) GetDocumentName() string {
	return d.name
}

func (d *Document) GetContent() string {
	return d.content
}

func (d *Document) Write(data string) error {
	d.content+=" " + data
	return nil
}

func NewDocument(id uint32, name string) *Document {
	return &Document{
		id: id,
		name: name,
	}
}