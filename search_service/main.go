package main

import search_service "serach_service/service"

func main() {
	searchService := search_service.NewSearchService()

	dataset1 := searchService.CreateDataSet(1)
	searchService.CreateDocumentInDataset(dataset1, 1, "i am roshan rai")
	searchService.CreateDocumentInDataset(dataset1, 2, "i am rohan rai rai rai")
	searchService.Search("rai")

	dataset2 := searchService.CreateDataSet(2)
	searchService.CreateDocumentInDataset(dataset2, 3, "i am roshan rai rai")
	searchService.CreateDocumentInDataset(dataset2, 4, "i am rohan rai")
	searchService.Search("rai")

}
