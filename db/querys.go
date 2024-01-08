package db

func RecordSimcardsQuery() string {
	return "INSERT into simcards (client, iccid, simcon, msisdn, ip, slot, installationdate, activationdate, supplier, operator, plan, apn, status, stock, substituted, nfsimcon, deliverydate, obs) Values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func RecordStockQuery() string {
	return "INSERT into simcardstock (iccid, supplier, operator, plan, apn, status, obs) Values(?, ?, ?, ?, ?, ?, ?)"
}

func GetAllSimcardsQuery() string {
	return "select * from simcards"
}

func GetAllStock() string {
	return "select * from simcardstock"
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
