package document_service

import (
	"document_service/domain/document"
	"document_service/domain/document_access_manager"
	"document_service/domain/user"
	user_documents_manager "document_service/domain/user_documents"
	"fmt"
)

type DocumentService interface {
	CreateDocument(user *user.User, docId uint32, name string) (*document.Document, error)
	ShowDocumentContent(user *user.User, doc document.Document) error
	WriteContent(user *user.User, doc *document.Document, data string) error
	GiveAccess(owner *user.User, other *user.User, doc *document.Document, permission document_access_manager.DocumentAccessType) error
	RemoveAccess(owner *user.User, other *user.User, doc *document.Document, permission document_access_manager.DocumentAccessType) error
}

type documentService struct {
	documentAccessManger *document_access_manager.DocumentAccessManager
	userDocumentsManger  *user_documents_manager.UserDocumentsManger
}

func NewDocumentService(documentAccessManger *document_access_manager.DocumentAccessManager,
	UserDocumentsManger *user_documents_manager.UserDocumentsManger) DocumentService {
	return &documentService{
		documentAccessManger: documentAccessManger,
		userDocumentsManger:  UserDocumentsManger,
	}
}

func (docServ *documentService) CreateDocument(user *user.User, docId uint32, name string) (*document.Document, error) {
	doc := document.NewDocument(docId, name)
	docServ.userDocumentsManger.AddUserDocument(user.GetUserId(), doc)
	docServ.documentAccessManger.GiveAccess(doc.GetDocumentId(), user.GetUserId(), document_access_manager.READ)
	docServ.documentAccessManger.GiveAccess(doc.GetDocumentId(), user.GetUserId(), document_access_manager.WRITE)
	docServ.documentAccessManger.GiveAccess(doc.GetDocumentId(), user.GetUserId(), document_access_manager.DELETE)
	return doc, nil
}

func (docServ *documentService) ShowDocumentContent(user *user.User, doc document.Document) error {
	// check access
	if !docServ.documentAccessManger.CheckAccess(doc.GetDocumentId(), user.GetUserId(), document_access_manager.READ) {
		return fmt.Errorf("no read access")
	}
	fmt.Printf("document content : %s\n", doc.GetContent())
	return nil
}

func (docServ *documentService) WriteContent(user *user.User, doc *document.Document, data string) error {
	// check permission
	if !docServ.documentAccessManger.CheckAccess(user.GetUserId(), doc.GetDocumentId(), document_access_manager.WRITE) {
		fmt.Errorf("no write access")
	}
	doc.Write(data)
	return nil
}

func (docServ *documentService) GiveAccess(owner *user.User, other *user.User, doc *document.Document, permission document_access_manager.DocumentAccessType) error {
	// check if user own the doc
	if !docServ.userDocumentsManger.CheckUserOwnTheDocument(owner.GetUserId(), doc.GetDocumentId()) {
		return fmt.Errorf("user not owner")
	}
	docServ.documentAccessManger.GiveAccess(doc.GetDocumentId(), other.GetUserId(), permission)
	return nil
}

func (docServ *documentService) RemoveAccess(owner *user.User, other *user.User, doc *document.Document, permission document_access_manager.DocumentAccessType) error {
	// check if user own the doc
	if !docServ.userDocumentsManger.CheckUserOwnTheDocument(owner.GetUserId(), doc.GetDocumentId()) {
		return fmt.Errorf("user not owner")
	}
	docServ.documentAccessManger.RevokeAccess(doc.GetDocumentId(), other.GetUserId(), permission)
	return nil
}
