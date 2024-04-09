package get

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"kdl_api_rest_internal/db"

	"kdl_api_rest_internal/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteSimcards(w http.ResponseWriter, r *http.Request) {
	var req models.IDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	stmt, err := db.MysqlDB.Prepare(db.DeleteSimcardQuery(req.ID))
	if err != nil {
		http.Error(w, "Erro ao preparar a consulta de exclusão", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var args []interface{}
	for _, id := range req.ID {
		args = append(args, id)
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		http.Error(w, "Erro ao executar a consulta de exclusão", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Erro ao obter o número de linhas afetadas", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "O item não existe no banco de dados", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Item(s) excluído(s) com sucesso")
}

func DeleteSimcardsStock(w http.ResponseWriter, r *http.Request) {
	var req models.ICCIDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	stmt, err := db.MysqlDB.Prepare(db.DeleteSimcardStockQuery(req.Iccid))
	if err != nil {
		http.Error(w, "Erro ao preparar a consulta de exclusão", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(req.Iccid)
	if err != nil {
		http.Error(w, "Erro ao executar a consulta de exclusão", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Erro ao obter o número de linhas afetadas", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "O item não existe no banco de dados", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Item excluído com sucesso")
}

func DeleteSimcardsStockForInclude(w http.ResponseWriter, r *http.Request) {
	var req models.IDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	stmt, err := db.MysqlDB.Prepare(db.DeleteSimcardStockForIncludeQuery(req.ID))
	if err != nil {
		http.Error(w, "Erro ao preparar a consulta de exclusão", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var args []interface{}
	for _, id := range req.ID {
		args = append(args, id)
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		http.Error(w, "Erro ao executar a consulta de exclusão", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Erro ao obter o número de linhas afetadas", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "O item não existe no banco de dados", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Item excluído com sucesso")
}
