package post

import (
	"encoding/json"
	"fmt"
	"io"
	"kdl_api_rest_internal/db"
	"kdl_api_rest_internal/models"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func RecordSimcard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusInternalServerError)
		return
	}

	var dados models.Simcards
	err = json.Unmarshal(body, &dados)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}

	Client := dados.Client
	Iccid := dados.Iccid
	Simcon := dados.Simcon
	Msisdn := dados.Msisdn
	Ip := dados.Ip
	Slot := dados.Slot
	Installationdate := dados.Installationdate
	Activationdate := dados.Activationdate
	Supplier := dados.Supplier
	Operator := dados.Operator
	Plan := dados.Plan
	Apn := dados.Apn
	Status := dados.Status
	Stock := dados.Stock
	Substituted := dados.Substituted
	Obs := dados.Obs

	rows, err := db.MysqlDB.Prepare(db.RecordSimcardsQuery())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	_, err = rows.Exec(
		Client,
		Iccid,
		Simcon,
		Msisdn,
		Ip,
		Slot,
		Installationdate,
		Activationdate,
		Supplier,
		Operator,
		Plan,
		Apn,
		Status,
		Stock,
		Substituted,
		Obs,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Dados gravados com sucesso!")
}
