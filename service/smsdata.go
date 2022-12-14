package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"strings"

	log "github.com/sirupsen/logrus"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func NewSMSData(str []string) *SMSData {
	sms := SMSData{}
	sms.Country = str[0]
	sms.Bandwidth = str[1]
	sms.ResponseTime = str[2]
	sms.Provider = str[3]
	return &sms
}

type StorageSD struct {
	storageSMSData map[int]*SMSData
}

func SmsData() ([]*SMSData, error) {
	log.Info("Получаем данные sms")
	var storageSMS []*SMSData
	providers := []string{"Topol", "Rond", "Kildy"}
	countriesString := helpers.CountryString()
	smsDataCSV := "./simulator/sms.data"
	smsDataString, err := helpers.CsvInString(smsDataCSV)
	if err != nil {
		log.Info(err)
		return storageSMS, err
	}
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 4)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 3)
	splitStrings = helpers.ExaminationCountry(splitStrings, countriesString)
	for _, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewSMSData(s)
		storageSMS = append(storageSMS, l)
	}
	log.Info("Получены данные sms")
	return storageSMS, err
}
