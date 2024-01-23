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

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.MysqlDB.Query(db.GetAllStockQuery())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	simcards := []models.Simcardstock{}
	for rows.Next() {
		var simcard models.Simcardstock
		err := rows.Scan(
			&simcard.Id,
			&simcard.Iccid,
			&simcard.Supplier,
			&simcard.Operator,
			&simcard.Plan,
			&simcard.Apn,
			&simcard.Status,
			&simcard.Obs,
		)
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

func GetAllLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.MysqlDB.Query(db.GetAllLogsQuery())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	timelines := []models.GetLogs{}
	for rows.Next() {
		var timeline models.GetLogs
		err := rows.Scan(
			&timeline.Logid,
			&timeline.SimcardId,
			&timeline.Action,
			&timeline.Timestamp,
			&timeline.Details,
		)
		if err != nil {
			log.Fatal(err)
		}
		timelines = append(timelines, timeline)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timelines)
}
