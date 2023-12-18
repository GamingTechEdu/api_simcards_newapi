package db

func RecordSimcardsQuery() string {
	return "INSERT into simcards (client, iccid, simcon, msisdn, ip, slot, installationdate, activationdate, supplier, operator, plan, apn, status, stock, substituted, obs) Values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetAllSimcardsQuery() string {
	return "select * from simcards"
}

func PostUserQuery() string {
	return "SELECT username, password, hierarchy FROM users WHERE username = ? and password = ?"
}

func DeleteSimcardQuery() string {
	return "DELETE FROM simcards WHERE id = ?"
}
