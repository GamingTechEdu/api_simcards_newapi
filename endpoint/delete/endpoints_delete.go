package get

import (
	"encoding/json"
	"fmt"
	"kdl_api_rest_internal/db"
	"kdl_api_rest_internal/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteSimcards(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar o corpo da requisição em uma struct IDRequest
	var req models.IDRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	// Executar a query de delete no banco de dados
	stmt, err := db.MysqlDB.Prepare(db.DeleteSimcardQuery())
	if err != nil {
		http.Error(w, "Erro ao preparar a consulta de exclusão", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(req.ID)
	if err != nil {
		http.Error(w, "Erro ao executar a consulta de exclusão", http.StatusInternalServerError)
		return
	}

	// Verificar se alguma linha foi afetada pela exclusão
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Erro ao obter o número de linhas afetadas", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "O item não existe no banco de dados", http.StatusNotFound)
		return
	}

	// Responder com uma mensagem de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Item excluído com sucesso")
}
