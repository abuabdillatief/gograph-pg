***DEFINITIONS:***
- **Logger:** 
     _a way or function to save/see logs when querying to DB_
- **Dataloader:** 
    when we are querying to database, where the query has a child objects,
    to decrease the amount of request to database, we collect all the related id
    of each child objects, and then we send the request.

    example:
        if we send query:
            query GetBooks {
                Books{
                    id
                    title
                    author {
                        id
                        name
                    }
                }
            }
        if there is 10 books in our database, the number of queries sent will be 11.
            - 1 to get all books
            - 10 to get each author of corresponding book
        by using dataloader, we can decrease the amount of request to only 2
            - 1 to get all books
            - 1 to get all author of corresponding book
        and this is done by collecting all corresponded id of each book beofre sending the query.

    to do this in golang, we use github.com/vektah/dataloaden package.
    ***AND TO USE THIS, WE HAVE TO RUN THE COMMAND INSIDE OUR MODEL FOLDER***
    command:
        go run github.com/vektah/dataloaden UserLoader string '[]*github.com/abuabdillatief/gograph-tutorial/graph/model.User'
        


