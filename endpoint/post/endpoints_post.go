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
	NfSimcon := dados.Nfsimcon
	Deliverydate := dados.Deliverydate
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
		NfSimcon,
		Deliverydate,
		Obs,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Dados gravados com sucesso!")
}

func RecordStock(w http.ResponseWriter, r *http.Request) {
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

	Iccid := dados.Iccid
	Supplier := dados.Supplier
	Operator := dados.Operator
	Plan := dados.Plan
	Apn := dados.Apn
	Status := dados.Status
	Obs := dados.Obs

	rows, err := db.MysqlDB.Prepare(db.RecordStockQuery())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	_, err = rows.Exec(
		Iccid,
		Supplier,
		Operator,
		Plan,
		Apn,
		Status,
		Obs,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Dados gravados com sucesso!")
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var userCreds models.Users
	err := json.NewDecoder(r.Body).Decode(&userCreds)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON da requisição", http.StatusBadRequest)
		return
	}

	row := db.MysqlDB.QueryRow(db.PostUserQuery(), userCreds.Username, userCreds.Password)

	err = row.Scan(&models.SendUser.Username, &models.SendUser.Password, &models.SendUser.Hierarchy)

	if err == nil && models.SendUser.Password == userCreds.Password {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.SendUser)
	} else {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		log.Println("Erro ao verificar credenciais:", err)
	}
}
