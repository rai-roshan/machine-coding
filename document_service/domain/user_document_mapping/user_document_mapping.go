package user_document_mapping

import (
	"document_service/domain/document"
	"document_service/domain/user"
)

type DocumentAccessType string

const (
	READ   DocumentAccessType = "READ"
	WRITE  DocumentAccessType = "WRITE"
	DELETE DocumentAccessType = "DELETE"
)

type UserDocumentMapping struct {
	user        *user.User
	document    *document.Document
	isOwner     bool
	permissions []DocumentAccessType
}

func NewUserDocumentMapping(user *user.User, document *document.Document, isOwner bool) *UserDocumentMapping {
	return &UserDocumentMapping{
		user:     user,
		document: document,
		isOwner:  isOwner,
	}
}

func (udm *UserDocumentMapping) IsOwner() bool {
	return udm.isOwner
}

func (udm *UserDocumentMapping) AddPermission(permission DocumentAccessType) error {
	udm.permissions = append(udm.permissions, permission)
	return nil
}

func (udm *UserDocumentMapping) RemovePermission(permission DocumentAccessType) error {
	currentPermissions := udm.permissions
	var newSetOfPermissions []DocumentAccessType
	for _, perm := range currentPermissions {
		if perm != permission {
			newSetOfPermissions = append(newSetOfPermissions, perm)
		}
	}
	udm.permissions = newSetOfPermissions
	return nil
}

func (udm *UserDocumentMapping) CheckPermission(permission DocumentAccessType) bool {
	for _, perm := range udm.permissions {
		if perm == permission {
			return true
		}
	}
	return false
}

func (udm *UserDocumentMapping) GetUserId() uint32 {
	return udm.user.GetUserId()
}

func (udm *UserDocumentMapping) GetDocumentId() uint32 {
	return udm.document.GetDocumentId()
}
