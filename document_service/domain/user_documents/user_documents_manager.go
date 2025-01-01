package user_documents_manager

import (
	"document_service/domain/document"
)

type UserDocumentsManger struct {
	userOwnedDocuments map[uint32][]*document.Document
}

func NewUserDocumentManager() *UserDocumentsManger {
	return &UserDocumentsManger{
		userOwnedDocuments: make(map[uint32][]*document.Document),
	}
}

func (udm *UserDocumentsManger) GetUserDocuments( userId uint32) []*document.Document {
	return udm.userOwnedDocuments[userId]
}

func (udm *UserDocumentsManger) AddUserDocument( userId uint32 , document *document.Document ) error {
	// TODO : manage duplicates
	udm.userOwnedDocuments[userId] = append(udm.userOwnedDocuments[userId], document)
	return nil 
}

func (udm *UserDocumentsManger) RemoveUserDocument(userId uint32, docId uint32) error {
	userDocuments := udm.userOwnedDocuments[userId]
	var newUserDocuments []*document.Document
	for _, doc := range userDocuments {
		if doc.GetDocumentId() != docId {
			newUserDocuments = append(newUserDocuments, doc)
		} 
	}
	udm.userOwnedDocuments[userId] = newUserDocuments
	return nil
}

func (udm *UserDocumentsManger) CheckUserOwnTheDocument( userId uint32, documentId uint32 ) bool {
	documents := udm.userOwnedDocuments[userId]
	for _, doc := range documents {
		if doc.GetDocumentId() == documentId {
			return true
		}
	}
	return false
}