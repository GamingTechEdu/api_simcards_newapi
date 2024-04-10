package db

import "strconv"

func RecordLogQuery() string {
	return "INSERT INTO simcard_logs (simcard_id, action, timestamp, details) VALUES (?, ?, NOW(), ?);"
}

func UpdateLogQuery() string {
	return "INSERT INTO simcard_logs (simcard_id, action, timestamp, details) VALUES (?, ?, NOW(), ?);"
}

func RecordSimcardsQuery() string {
	return "INSERT into simcards (client, iccid, simcon, msisdn, ip, slot, installationdate, activationdate, supplier, operator, plan, apn, status, stock, substituted, nfsimcon, deliverydate, obs) Values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
}

func updateSimcardQuery() string {
	return ""
}

// DELETE FROM simcards WHERE iccid = ?

// func RecordSimcardsQuery() string {
// 	return "INSERT into simcards (client, iccid, simcon, msisdn, ip, slot, installationdate, activationdate, supplier, operator, plan, apn, status, stock, substituted, nfsimcon, deliverydate, obs) Values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
// }

func RecordStockQuery() string {
	return "INSERT into simcardstock (iccid, supplier, operator, plan, apn, status, obs) Values(?, ?, ?, ?, ?, ?, ?)"
}

func GetAllSimcardsQuery() string {
	return "select * from simcards"
}

func GetAllSimucsQuery() string {
	return "select * from simucs"
}

// func GetListSimucsQuery() string {
// 	return "select * from simucs where numberDoc = 142 && nserlum = 2220003551"
// }

func GetListSimucsQuery(numberDoc, nserlum int) string {
	return "SELECT * FROM simucs WHERE numberDoc = " + strconv.Itoa(numberDoc) + " AND nserlum = " + strconv.Itoa(nserlum)
}

func PostUserQuery() string {
	return "SELECT username, password, hierarchy FROM users WHERE username = ? and password = ?"
}

func DeleteSimcardQuery(ids []string) string {
	query := "DELETE FROM simcards WHERE id IN ("

	for i := range ids {
		query += "?"
		if i < len(ids)-1 {
			query += ", "
		}
	}
	query += ")"
	return query
}

func DeleteSimcardStockQuery(iccid string) string {
	return "DELETE FROM simcardstock WHERE iccid = ?"
}

func DeleteSimcardStockForIncludeQuery(ids []string) string {
	query := "DELETE FROM simcardstock WHERE id IN ("

	for i := range ids {
		query += "?"
		if i < len(ids)-1 {
			query += ", "
		}
	}
	query += ")"
	return query
}
