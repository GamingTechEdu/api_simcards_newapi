package utils

import (
	"database/sql"
	"encoding/json"
	"kdl_api_rest_internal/models"

	_ "github.com/go-sql-driver/mysql"
)

func DeserializeSimcards(data []byte) ([]models.Simcards, error) {
	var simcards []models.Simcards
	err := json.Unmarshal(data, &simcards)
	if err != nil {
		return nil, err
	}
	return simcards, nil
}

func ScanSimcardRow(row *sql.Rows) (models.Simcards, error) {
	var simcard models.Simcards
	err := row.Scan(
		&simcard.Id,
		&simcard.Client,
		&simcard.Iccid,
		&simcard.Simcon,
		&simcard.Msisdn,
		&simcard.Ip,
		&simcard.Slot,
		&simcard.Installationdate,
		&simcard.Activationdate,
		&simcard.Supplier,
		&simcard.Operator,
		&simcard.Plan,
		&simcard.Apn,
		&simcard.Status,
		&simcard.Stock,
		&simcard.Substituted,
		&simcard.Nfsimcon,
		&simcard.Deliverydate,
		&simcard.Obs,
	)

	return simcard, err
}

func ScanSimucRow(row *sql.Row, simcard *models.GetListSimuc) error {
	err := row.Scan(
		&simcard.Id,
		&simcard.Nserlum,
		&simcard.Doc,
		&simcard.NumberDoc,
		&simcard.Datedoc,
		&simcard.User,
		&simcard.Violation,
		&simcard.Obs,
		&simcard.EvaluatorDate,
		&simcard.DefectRelated,
		&simcard.ApparentDefect,
		&simcard.Defect,
		&simcard.Evaluator,
		&simcard.Components,
		&simcard.Value,
		&simcard.Garantee,
		&simcard.Nivel,
		&simcard.Exist,
	)

	return err
}

func ScanListSimucRow(row *sql.Rows) (models.GetListSimuc, error) {
	var simcard models.GetListSimuc
	err := row.Scan(
		&simcard.Id,
		&simcard.Nserlum,
		&simcard.Doc,
		&simcard.NumberDoc,
		&simcard.Datedoc,
		&simcard.User,
		&simcard.Violation,
		&simcard.Obs,
		&simcard.EvaluatorDate,
		&simcard.DefectRelated,
		&simcard.ApparentDefect,
		&simcard.Defect,
		&simcard.Evaluator,
		&simcard.Components,
		&simcard.Value,
		&simcard.Garantee,
		&simcard.Nivel,
		&simcard.Exist,
	)

	return simcard, err
}
