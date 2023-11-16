package test

import (
	"encoding/json"
	"fmt"
	"kdl_api_rest_internal/models"
	"testing"
)

func simcard() *models.Simcards {
	simcard := models.Simcards{
		Client:           "ClientA",
		Iccid:            "9876543210123456789",
		Simcon:           "2222001",
		Msisdn:           "1234567890",
		Ip:               "192.168.1.1",
		Slot:             "1",
		Installationdate: "2023-01-01",
		Activationdate:   "2023-01-02",
		Supplier:         "SupplierA",
		Operator:         "OperatorA",
		Plan:             "PlanA",
		Apn:              "apn.example.com",
		Obs:              "Observation",
	}

	return &simcard
}

func TestSimcardsJSONSerialization(t *testing.T) {
	// Atribuir o resultado da função a uma variável
	simcardInstance := simcard()

	// Serializar para JSON
	jsonData, err := json.Marshal(simcardInstance)
	if err != nil {
		t.Errorf("Erro ao serializar Simcards para JSON: %v", err)
	}

	// Desserializar JSON de volta para Simcards
	var decodedSimcard models.Simcards
	err = json.Unmarshal(jsonData, &decodedSimcard)
	if err != nil {
		t.Errorf("Erro ao desserializar JSON para Simcards: %v", err)
	}

	// Verificar se os valores são iguais
	if *simcardInstance != decodedSimcard {
		t.Errorf("Os dados desserializados não correspondem aos dados originais.\nEsperado: %+v\nObtido:   %+v", *simcardInstance, decodedSimcard)
	}

	fmt.Println("TestSimcardsJSONSerialization() passed!")
}

func TestSimcardsEmptyFields(t *testing.T) {
	// Criar uma instância de Simcards com campos vazios, exceto para Obs
	simcard := &models.Simcards{
		Client:           "",
		Iccid:            "",
		Simcon:           "",
		Msisdn:           "",
		Ip:               "",
		Slot:             "",
		Installationdate: "",
		Activationdate:   "",
		Supplier:         "",
		Operator:         "",
		Plan:             "",
		Apn:              "",
		Obs:              "Observation",
	}

	// Validar campos antes da serialização
	if err := validateSimcard(simcard); err != nil {
		t.Errorf("Erro ao validar Simcards com campos vazios: %v", err)
		return
	}

	// Serializar para JSON
	jsonData, err := json.Marshal(simcard)
	if err != nil {
		t.Errorf("Erro ao serializar Simcards com campos vazios para JSON: %v", err)
		return
	}

	// Desserializar JSON de volta para Simcards
	var decodedSimcard models.Simcards
	err = json.Unmarshal(jsonData, &decodedSimcard)
	if err != nil {
		t.Errorf("Erro ao desserializar JSON para Simcards com campos vazios: %v", err)
		return
	}

	// Verificar se os valores são iguais
	if *simcard != decodedSimcard {
		t.Errorf("Os dados desserializados não correspondem aos dados originais com campos vazios.\nEsperado: %+v\nObtido:   %+v", *simcard, decodedSimcard)
	}
}

func validateSimcard(simcard *models.Simcards) error {
	if simcard.Client == "" {
		return ErrEmptyField("Client")
	}
	if simcard.Iccid == "" {
		return ErrEmptyField("Iccid")
	}
	if simcard.Simcon == "" {
		return ErrEmptyField("Simcon")
	}
	if simcard.Msisdn == "" {
		return ErrEmptyField("Msisdn")
	}
	if simcard.Ip == "" {
		return ErrEmptyField("Ip")
	}
	if simcard.Slot == "" {
		return ErrEmptyField("Slot")
	}
	if simcard.Installationdate == "" {
		return ErrEmptyField("Installationdate")
	}
	if simcard.Activationdate == "" {
		return ErrEmptyField("Activationdate")
	}
	if simcard.Supplier == "" {
		return ErrEmptyField("Supplier")
	}
	if simcard.Operator == "" {
		return ErrEmptyField("Operator")
	}
	if simcard.Plan == "" {
		return ErrEmptyField("Plan")
	}
	if simcard.Apn == "" {
		return ErrEmptyField("Apn")
	}

	return nil
}

type ErrEmptyField string

func (e ErrEmptyField) Error() string {
	return "campo vazio: " + string(e)
}
