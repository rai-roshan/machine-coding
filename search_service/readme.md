Your organization has started a new tech blog with interesting tech stories and you’re responsible for designing and implementing an in-memory search engine, supporting the search functionality on the blog content.


Requirements:
- It should be possible to create a dataset in the search engine.
- It should be possible to insert documents into a given dataset. Each document is simply a piece of text.
- It should be possible to search through documents for a search pattern in a given dataset.
- It should be possible to order the search results
- Search Pattern: Initially search pattern would be sequence of words, but can be extensible for different search patterns:
    - Sequence of words containing words
    - Sequence of words containing words maintaining the order
    - Sequence of words containing words maintaining the order in sequence(substring)
Ordering Logic: Initially ordering pattern would be latest published article, but can be extensible for different search patterns:
    - Ordering can be recently published articles
    - Ordering can be recently updated articles
    - Ordering can be recently searched articles
    - Ordering can be done based on Author.
