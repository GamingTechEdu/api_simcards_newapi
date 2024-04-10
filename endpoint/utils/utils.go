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

<<<<<<< HEAD
// func ScanSimucRow(row *sql.Rows) (models.GetListSimuc, error) {
// 	var simcard models.GetListSimuc
// 	err := row.Scan(
// 		&simcard.Id,
// 		&simcard.Nserlum,
// 		&simcard.Doc,
// 		&simcard.NumberDoc,
// 		&simcard.Datedoc,
// 		&simcard.User,
// 		&simcard.Violation,
// 		&simcard.Obs,
// 		&simcard.EvaluatorDate,
// 		&simcard.DefectRelated,
// 		&simcard.ApparentDefect,
// 		&simcard.Defect,
// 		&simcard.Evaluator,
// 		&simcard.Components,
// 		&simcard.Value,
// 		&simcard.Garantee,
// 		&simcard.Nivel,
// 		&simcard.Exist,
// 	)

// 	return simcard, err
// }

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
=======
func DeserializeSimucs(data []byte) ([]models.Simucs, error) {
	var simucs []models.Simucs
	err := json.Unmarshal(data, &simucs)
	if err != nil {
		return nil, err
	}
	return simucs, nil
}

func ScanSimucsRow(row *sql.Rows) (models.Simucs, error) {
	var simuc models.Simucs
	err := row.Scan(
		&simuc.Id,
		&simuc.Nserlum,
		&simuc.Doc,
		&simuc.NumberDoc,
		&simuc.DateDoc,
		&simuc.User,
		&simuc.Violation,
		&simuc.Obs,
		&simuc.EvaluatorDate,
		&simuc.DefectRelated,
		&simuc.ApparentDefect,
		&simuc.Defect,
		&simuc.Evaluator,
		&simuc.Components,
		&simuc.Value,
		&simuc.Garantee,
		&simuc.Nivel,
		&simuc.Exist,
	)

	return simuc, err
>>>>>>> 05e0cc803d9052df065ccf8518e5f9d9d4386821
}
