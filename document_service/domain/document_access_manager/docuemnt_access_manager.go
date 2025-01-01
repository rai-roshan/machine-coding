package document_access_manager

import "fmt"

type DocumentAccessType string

const (
	READ   DocumentAccessType = "READ"
	WRITE  DocumentAccessType = "WRITE"
	DELETE DocumentAccessType = "DELETE"
)

type userIdToAccessList map[uint32][]DocumentAccessType
type documentIdToUserAccess map[uint32]userIdToAccessList

type DocumentAccessManager struct {
	accessTable documentIdToUserAccess
}

func NewDocumentAccessManager() *DocumentAccessManager {
	return &DocumentAccessManager{
		accessTable: make(documentIdToUserAccess),
	}
}

func (dacm *DocumentAccessManager) CheckAccess(docId, userId uint32, accessType DocumentAccessType) bool {
	documentUserAccessData, ok := dacm.accessTable[docId]
	if !ok {
		return false
	}
	accessList, ok := documentUserAccessData[userId]
	if !ok {
		return false
	}

	for _, access := range accessList {
		if access == accessType {
			return true
		}
	}
	return false
}

func (dacm *DocumentAccessManager) GiveAccess(docId, userId uint32, accessType DocumentAccessType) error {
	if dacm.CheckAccess(docId, userId, accessType) {
		return nil
	}
	userAccessTable, ok := dacm.accessTable[docId]
	if ok {
		userAccessTable[userId] = append(userAccessTable[userId], accessType)
	} else {
		userAccessTable = make( userIdToAccessList )
		userAccessTable[userId] = []DocumentAccessType{accessType}
	}
	dacm.accessTable[docId] = userAccessTable
	return nil
}

func (dacm *DocumentAccessManager) RevokeAccess(docId, userId uint32, accessType DocumentAccessType) error {
	if dacm.CheckAccess(docId, userId, accessType) {
		accessList := dacm.accessTable[docId][userId]
		newAccessList := []DocumentAccessType{}
		for _, access := range accessList {
			if access != accessType {
				newAccessList = append(newAccessList, access)
			}
		}
		dacm.accessTable[docId][userId] = newAccessList
	}
	return fmt.Errorf("user already dont have the access")
}