package db

func RecordSimcardsQuery() string {
	return "Insert into simcards (client, iccid, simcon, msisdn, ip, slot, installationdate, activationdate, supplier, operator, plan, apn, status, stock, substituted, obs) Values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetAllSimcardsQuery() string {
	return "select * from simcards"
}

func PostUserQuery() string {
	return "SELECT username, password FROM users WHERE username = ? and password = ?"
}
