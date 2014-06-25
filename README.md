MapSQL
======

Library written in go for inserting or updating MySQL Tables with `map[string]interface{}` Will save you the time to write insert function again and again for every table!

Usage
======

Example of inserting:

    import "github.com/shaymargolis/mapsql"
    func insertCafe(cafe_to map[string]interface{}) (int, error) {
        return mapsql.insertToDB("cafe", cafe_to) 
    }
    //..
    // Insert:
    cafe_to_insert := map[string]interface{} {
        "name": "The Club",
        "chain": "Aroma",
    }
    insertedID, err := insertCafe(cafe_to_insert)
    if err != nil {
        // Error
        fmt.Prinln(err)
    } else {
        // All was fine. the function will insert to table `cafe` cafe with name "The club" and chain "Aroma".
        fmt.Println("We? Did it! inserted id: ", insertedID)
    }
    
Exmaple of updating:

    func updateCafe(ID int, cafe_to map[string]interface{}) error {
        return updateDB("cafe", ID, cafe_to)
    }
    //..
    // Update:
    cafe_to_update := map[string]interface{} {
        "name": "The *Hilarious* Club",
        "chain": "OtherHilariousChain",
        "is_kosher": 1, // Other Column that we wish to change
    }
    err := insertCafe(insertedID, cafe_to_insert)
    if err != nil {
        // Error
        fmt.Prinln(err)
    } else {
        // All was fine. the function will change table `cafe` item with ID insertedID to "name":"The *Hilarious* Club", "chain":"OtherHilariousChain" and "is_kosher":1 (true)
        fmt.Println("Yay!")
    }
    
