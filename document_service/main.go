package main

import (
	// "document_service/domain/document_access_manager"
	"document_service/domain/user"
	user_documents_manager "document_service/domain/user_documents"
	document_service "document_service/service"
	"fmt"
)

func main() {

	// docAccessManger := document_access_manager.NewDocumentAccessManager()
	userDocManger := user_documents_manager.NewUserDocumentManager()

	docService := document_service.NewDocumentService(
		// docAccessManger, 
		userDocManger,
	)

	user1 := user.NewUser(1, "roshan")
	user2 := user.NewUser(2, "rohan")

	doc1, err := docService.CreateDocument(user1, 1, "doc1")
	if err != nil {
		fmt.Println(err)
	}
	doc2, err := docService.CreateDocument(user2, 2, "doc2")
	if err != nil {
		fmt.Println(err)
	}
	err = docService.WriteContent(user1, doc1, "roshan rai content is rthis")
	if err != nil {
		fmt.Println(err)
	}
	err = docService.ShowDocumentContent(user1, *doc1)
	if err != nil {
		fmt.Println(err)
	}

	err = docService.WriteContent(user2, doc2, "doc 222 roshan rai content is rthis")
	if err != nil {
		fmt.Println(err)
	}

	err = docService.ShowDocumentContent(user2, *doc2)
	if err != nil {
		fmt.Println(err)
	}

	err = docService.ShowDocumentContent(user2, *doc1)
	if err != nil {
		fmt.Println(err)
	}
	err = docService.ShowDocumentContent(user1, *doc2)
	if err != nil {
		fmt.Println(err)
	}

}
