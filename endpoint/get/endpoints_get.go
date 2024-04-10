package get

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"kdl_api_rest_internal/db"
	"kdl_api_rest_internal/endpoint/utils"
	"kdl_api_rest_internal/models"
	"log"
	"net/http"
	"strconv"

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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(timelines)
}

func GetListIccids(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.MysqlDB.Query(db.GetListIccidsQuery())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	listIccids := []models.GetListIccid{}
	for rows.Next() {
		var listIccid models.GetListIccid
		err := rows.Scan(
			&listIccid.Iccid,
		)
		if err != nil {
			log.Fatal(err)
		}
		listIccids = append(listIccids, listIccid)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(listIccids)
}

// func GetListSimucs(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	var requestData map[string]interface{}
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Extrair numberDoc e nserlum do requestData como strings
// 	numberDocStr, ok := requestData["numberDoc"].(string)
// 	if !ok {
// 		http.Error(w, "Invalid numberDoc", http.StatusBadRequest)
// 		return
// 	}
// 	nserlumStr, ok := requestData["nserlum"].(string)
// 	if !ok {
// 		http.Error(w, "Invalid nserlum", http.StatusBadRequest)
// 		return
// 	}

// 	// Converter strings para inteiros
// 	numberDoc, err := strconv.Atoi(numberDocStr)
// 	if err != nil {
// 		http.Error(w, "Invalid numberDoc format", http.StatusBadRequest)
// 		return
// 	}
// 	nserlum, err := strconv.Atoi(nserlumStr)
// 	if err != nil {
// 		http.Error(w, "Invalid nserlum format", http.StatusBadRequest)
// 		return
// 	}

// 	query := fmt.Sprintf("SELECT * FROM simucs WHERE numberDoc = %d AND nserlum = %d", numberDoc, nserlum)

// 	rows, err := db.MysqlDB.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	listSimucs := []models.GetListSimuc{}
// 	for rows.Next() {
// 		listSimuc, err := utils.ScanSimucRow(rows)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		listSimucs = append(listSimucs, listSimuc)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Verificar se a lista está vazia e retornar BadRequest se for o caso
// 	if len(listSimucs) == 0 {
// 		http.Error(w, "No items found with the given parameters", http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	json.NewEncoder(w).Encode(listSimucs)
// }

func GetListSimucs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var requestData map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Extrair numberDoc e nserlum do requestData como strings
	numberDocStr, ok := requestData["numberDoc"].(string)
	if !ok {
		http.Error(w, "Invalid numberDoc", http.StatusBadRequest)
		return
	}
	nserlumStr, ok := requestData["nserlum"].(string)
	if !ok {
		http.Error(w, "Invalid nserlum", http.StatusBadRequest)
		return
	}

	// Converter strings para inteiros
	numberDoc, err := strconv.Atoi(numberDocStr)
	if err != nil {
		http.Error(w, "Invalid numberDoc format", http.StatusBadRequest)
		return
	}
	nserlum, err := strconv.Atoi(nserlumStr)
	if err != nil {
		http.Error(w, "Invalid nserlum format", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("SELECT * FROM simucs WHERE numberDoc = %d AND nserlum = %d LIMIT 1", numberDoc, nserlum)

	row := db.MysqlDB.QueryRow(query)
	listSimuc := models.GetListSimuc{}
	err = utils.ScanSimucRow(row, &listSimuc)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No item found with the given parameters", http.StatusNotFound)
			return
		}
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(listSimuc)
}
