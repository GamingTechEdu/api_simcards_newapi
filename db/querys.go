package db

func RecordSimcardsQuery() string {
	return "Insert into simcards (client, iccid, simcon, msisdn, ip, slot, installationdate, activationdate, supplier, operator, plan, apn, status, stock, substituted, obs) Values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetAllSimcardsQuery() string {
	return "select * from simcards"
}

func GetUser() {
	return "SELECT EXISTS(SELECT 1 FROM users WHERE username=? AND password=?)" + user.Username, user.Password + ""
}
