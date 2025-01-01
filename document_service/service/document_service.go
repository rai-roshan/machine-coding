package document_service

import (
	"document_service/domain/document"
	"document_service/domain/user"
	"document_service/domain/user_document_mapping"
	user_documents_manager "document_service/domain/user_documents"
	"fmt"
)

type DocumentService interface {
	CreateDocument(user *user.User, docId uint32, name string) (*document.Document, error)
	ShowDocumentContent(user *user.User, doc document.Document) error
	WriteContent(user *user.User, doc *document.Document, data string) error
	GiveAccess(owner *user.User, other *user.User, doc *document.Document, permission user_document_mapping.DocumentAccessType) error
	RemoveAccess(owner *user.User, other *user.User, doc *document.Document, permission user_document_mapping.DocumentAccessType) error
}

type documentService struct {
	userDocumentsManger *user_documents_manager.UserDocumentsManger
}

func NewDocumentService(
	UserDocumentsManger *user_documents_manager.UserDocumentsManger) DocumentService {
	return &documentService{
		userDocumentsManger: UserDocumentsManger,
	}
}

func (docServ *documentService) CreateDocument(user *user.User, docId uint32, name string) (*document.Document, error) {
	doc := document.NewDocument(docId, name)

	userDocMapping := user_document_mapping.NewUserDocumentMapping(user, doc, true)
	userDocMapping.AddPermission(user_document_mapping.READ)
	userDocMapping.AddPermission(user_document_mapping.WRITE)
	userDocMapping.AddPermission(user_document_mapping.DELETE)

	docServ.userDocumentsManger.AddUserDocument(user.GetUserId(), userDocMapping)
	return doc, nil
}

func (docServ *documentService) ShowDocumentContent(user *user.User, doc document.Document) error {
	// check access
	if !docServ.userDocumentsManger.CheckPermissionForUser(doc.GetDocumentId(), user.GetUserId(), user_document_mapping.READ) {
		return fmt.Errorf("no read access")
	}
	fmt.Printf("document content : %s\n", doc.GetContent())
	return nil
}

func (docServ *documentService) WriteContent(user *user.User, doc *document.Document, data string) error {
	// check permission
	if !docServ.userDocumentsManger.CheckPermissionForUser(user.GetUserId(), doc.GetDocumentId(), user_document_mapping.WRITE) {
		fmt.Errorf("no write access")
	}
	doc.Write(data)
	return nil
}

func (docServ *documentService) GiveAccess(owner *user.User, other *user.User, doc *document.Document, permission user_document_mapping.DocumentAccessType) error {
	// check if user own the doc
	if !docServ.userDocumentsManger.CheckUserOwnTheDocument(owner.GetUserId(), doc.GetDocumentId()) {
		return fmt.Errorf("user not owner")
	}
	docServ.userDocumentsManger.GiveUserAccess(doc.GetDocumentId(), other.GetUserId(), permission)
	return nil
}

func (docServ *documentService) RemoveAccess(owner *user.User, other *user.User, doc *document.Document, permission user_document_mapping.DocumentAccessType) error {
	// check if user own the doc
	if !docServ.userDocumentsManger.CheckUserOwnTheDocument(owner.GetUserId(), doc.GetDocumentId()) {
		return fmt.Errorf("user not owner")
	}
	docServ.userDocumentsManger.RemoveUserAccess(doc.GetDocumentId(), other.GetUserId(), permission)
	return nil
}
