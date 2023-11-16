package get

import (
	"encoding/json"
	"kdl_api_rest_internal/db"
	"kdl_api_rest_internal/endpoint/utils"
	"kdl_api_rest_internal/models"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllSimcards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.MysqlDB.Query(db.GetAllSimcardsQuery())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	simcards := []models.Simcards{}
	for rows.Next() {
		simcard, err := utils.ScanSimcardRow(rows)
		if err != nil {
			log.Fatal(err)
		}
		simcards = append(simcards, simcard)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(simcards)
}

// func ScanSimcardRow(row *sql.Rows) (models.Simcards, error) {
// 	var simcard models.Simcards
// 	var err error

// 	// Assuming Simcards model has fields Field1, Field2, ..., FieldN
// 	var Obs sql.NullString
// 	// ...
// 	// var fieldN sql.NullString

// 	err = row.Scan(&Obs)
// 	if err != nil {
// 		return simcard, err
// 	}

// 	simcard.Obs = convertNullString(Obs)
// 	// ...
// 	// simcard.FieldN = convertNullString(fieldN)

// 	return simcard, nil
// }

// func convertNullString(nullString sql.NullString) string {
// 	if nullString.Valid {
// 		return nullString.String
// 	}
// 	return ""
// }
