package user_documents_manager

import (
	// "document_service/domain/document"
	// "document_service/domain/document_access_manager"
	"document_service/domain/user_document_mapping"
)

type UserDocumentsManger struct {
	documentIdToUsers map[uint32]map[uint32]*user_document_mapping.UserDocumentMapping
	userIdToDocuments map[uint32]map[uint32]*user_document_mapping.UserDocumentMapping
}

func NewUserDocumentManager() *UserDocumentsManger {
	return &UserDocumentsManger{
		documentIdToUsers: make(map[uint32]map[uint32]*user_document_mapping.UserDocumentMapping),
		userIdToDocuments: make(map[uint32]map[uint32]*user_document_mapping.UserDocumentMapping),
	}
}

func (udm *UserDocumentsManger) AddUserDocument(userId uint32, userDocMapp *user_document_mapping.UserDocumentMapping) error {

	_, ok := udm.documentIdToUsers[userDocMapp.GetDocumentId()]
	if !ok {
		udm.documentIdToUsers[userDocMapp.GetDocumentId()] = make(map[uint32]*user_document_mapping.UserDocumentMapping)
		udm.documentIdToUsers[userDocMapp.GetDocumentId()][userDocMapp.GetUserId()] = userDocMapp
	} else {
		udm.documentIdToUsers[userDocMapp.GetDocumentId()][userDocMapp.GetUserId()] = userDocMapp
	}

	_, ok = udm.userIdToDocuments[userDocMapp.GetUserId()]
	if !ok {
		udm.userIdToDocuments[userDocMapp.GetUserId()] = make(map[uint32]*user_document_mapping.UserDocumentMapping)
		udm.userIdToDocuments[userDocMapp.GetUserId()][userDocMapp.GetDocumentId()] = userDocMapp
	} else {
		udm.userIdToDocuments[userDocMapp.GetUserId()][userDocMapp.GetDocumentId()] = userDocMapp
	}

	return nil
}

func (udm *UserDocumentsManger) RemoveUserDocument(userId uint32, docId uint32) error {

	delete(udm.documentIdToUsers, docId)
	delete(udm.userIdToDocuments, userId)

	return nil
}

func (udm *UserDocumentsManger) CheckUserOwnTheDocument(userId uint32, documentId uint32) bool {
	docIdToUserDocMapp, ok := udm.userIdToDocuments[userId]
	if !ok {
		return false
	}
	userDocMapp, ok := docIdToUserDocMapp[documentId]
	if !ok || userDocMapp == nil {
		return false
	}

	return userDocMapp.IsOwner()
}

func (udm *UserDocumentsManger) CheckPermissionForUser(userId uint32, docId uint32, permission user_document_mapping.DocumentAccessType) bool {
	docIdToUserDocMapp, ok := udm.userIdToDocuments[userId]
	if !ok {
		return false
	}
	if docIdToUserDocMapp[docId] == nil {
		return false
	}
	return docIdToUserDocMapp[docId].CheckPermission(permission)
}

func (udm *UserDocumentsManger) GiveUserAccess(userId, docId uint32, permission user_document_mapping.DocumentAccessType) {
	docIdToUserDocMapp, ok := udm.userIdToDocuments[userId]
	if !ok {
		return
	}
	docIdToUserDocMapp[docId].AddPermission(permission)
}

func (udm *UserDocumentsManger) RemoveUserAccess(userId uint32, docId uint32, permission user_document_mapping.DocumentAccessType) {
	docIdToUserDocMapp, ok := udm.userIdToDocuments[userId]
	if !ok {
		return
	}
	docIdToUserDocMapp[docId].RemovePermission(permission)
}
