/* functions.go Bacafe API shay@inbar.co.il */

package main

import (
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
)

func updateDB(table string, ID int, data map[string]interface{}) error {
    columns := make([]interface{}, 0, len(data))
    values := make([]interface{}, 0, len(data))

    for  key, _ := range data {
       columns = append(columns, key)
       values = append(values, data[key])
    }
    
    i := 0
    of := len(data)
    
    // Will look like this: column = ?, column2 = ?, etc...
    all_text := ""
    
    for i < of {
        
        column := columns[i].(string)
        
        if i == (of - 1) {
            all_text = all_text + column + " = ?"
        } else {
            all_text = all_text + column + " = ?, "
        }
        
        i++
    }
    
    values = append(values, ID)
    
    db, err := sql.Open("mysql", "root:root@/bacafe")
    if err != nil {
        return err
    }
    defer db.Close()
    
    stmtIns, err := db.Prepare("UPDATE " + table + " SET " +  all_text + " WHERE ID = ?")
    if err != nil {
        return err
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
        
    _, err = stmtIns.Exec(values...)
    if err != nil {
        return err
    } else {
        return nil
    }
}

func insertToDB(table string, data map[string]interface{}) (int, error) {

    columns := make([]interface{}, 0, len(data))
    values := make([]interface{}, 0, len(data))

    for  key, _ := range data {
       columns = append(columns, key)
       values = append(values, data[key])
    }
    
    columns_text := ""
    i := 0
    of := len(data)
    
    for i < of {
        column := columns[i].(string)
        
        if i == 0 {
            columns_text = column
        } else {
            columns_text = columns_text + ", " + column
        }
        
        i++
    }
        
    values_text := ""
    
    i = 0
    
    for i < of {
            
        if i == 0 {
            values_text = "?"
        } else {
            values_text = values_text + ", ?"
        }
        
        i++
    }
    
    db, err := sql.Open("mysql", "root:root@/bacafe")
    if err != nil {
        return -1, err
    }
    defer db.Close()
    
    stmtIns, err := db.Prepare("INSERT INTO " + table + " ( " + columns_text + " ) VALUES ( " +  values_text + " )")
    if err != nil {
        return -1, err
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
        
    result, err := stmtIns.Exec(values...)
    if err != nil {
        return -1, err
    } else {
        insertedID, err := result.LastInsertId()
        if err != nil {
            return -1, err
        } else {
            return int(insertedID), nil
        }
    }
}