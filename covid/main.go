/*https://github.com/pcm-dpc/COVID-19*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Nazione struct {
	Data                      string `json:"data"`
	Stato                     string `json:"stato"`
	RicoveratiConSintomi      int    `json:"ricoverati_con_sintomi"`
	TerapiaIntensiva          int    `json:"terapia_intensiva"`
	TotaleOspedalizzati       int    `json:"totale_ospedalizzati"`
	IsolamentoDomiciliare     int    `json:"isolamento_domiciliare"`
	TotalePositivi            int    `json:"totale_positivi"`
	VariazioneTotalePositivi  int    `json:"variazione_totale_positivi"`
	NuoviPositivi             int    `json:"nuovi_positivi"`
	DimessiGuariti            int    `json:"dimessi_guariti"`
	Deceduti                  int    `json:"deceduti"`
	CasiDaSospettoDiagnostico string `json:"casi_da_sospetto_diagnostico"`
	CasiDaScreening           string `json:"casi_da_screening"`
	TotaleCasi                int    `json:"totale_casi"`
	Tamponi                   int    `json:"tamponi"`
	CasiTestati               int    `json:"casi_testati"`
	Note                      string `json:"note"`
	IngressiTerapiaIntensiva  int    `json:"ingressi_terapia_intensiva"`
	NoteTest                  string `json:"note_test"`
	NoteCasi                  string `json:"note_casi"`
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	var italia []Nazione

	response, err := http.Get("https://raw.githubusercontent.com/pcm-dpc/COVID-19/master/dati-json/dpc-covid19-ita-andamento-nazionale-latest.json")
	if err != nil {
		panic(fmt.Sprintf("\nErrore nel caricamento dei dati:\n%v", err))
	}

	defer response.Body.Close()
	responseByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(fmt.Sprintf("\nErrore nella conversione dei dati:\n%v", err))
	}
	json.Unmarshal(responseByte, &italia)

	for _, stato := range italia {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Totale positivi: ", stato.TotalePositivi)
		fmt.Println("Totale deceduti: ", stato.Deceduti)
		fmt.Println("Totale in ospedale: ", stato.TotaleOspedalizzati)
		fmt.Println("Totale in terapia intensiva: ", stato.TerapiaIntensiva)
		fmt.Println("-----------------------------------------------------")
		fmt.Println("OGGI:")
		fmt.Println("Nuovi positivi: ", stato.NuoviPositivi)
		fmt.Println("Variazione positivi: ", stato.VariazioneTotalePositivi)
		fmt.Println("Nuovi in terapia intensiva: ", stato.IngressiTerapiaIntensiva)
	}
}
