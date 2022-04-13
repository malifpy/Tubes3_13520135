package dbhandler;

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
)

type Album struct {
    ID     int64  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func Insert(db *sql.DB, alb Album){

    sqlQuery := `INSERT INTO albums (title, artist, price) VALUES($1, $2, $3) RETURNING id;`
    id := 0
    err := db.QueryRow(sqlQuery, alb.Title, alb.Artist, alb.Price).Scan(&id)
    if err != nil {
        panic(err)
    }
}

func ViewAll(db *sql.DB) ([]Album, error){
    var albums []Album
    rows, err := db.Query("SELECT * FROM albums;")
    if err != nil {
        return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
    }
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
    }
    defer rows.Close()
    return albums, nil
}
