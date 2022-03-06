package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gocarina/gocsv"
	"github.com/kamil5b/GoNeuralNetwork"
)

func LoadAUSProcessedTargetFloat64(namafile string) ([][]float64, []int) {
	/*[Date  float64 Location  float64 MinTemp  float64 MaxTemp  float64 Rainfall  float64 Evaporation
	  Sunshine  float64 WindGustDir  float64 WindGustSpeed  float64 WindDir9am  float64 WindDir3pm
	  WindSpeed9am  float64 WindSpeed3pm  float64 Humidity9am  float64 Humidity3pm
	  Pressure9am  float64 Pressure3pm  float64 Cloud9am  float64 Cloud3pm  float64 Temp9am
	  Temp3pm  float64 RainToday  float64 RainTomorrow  float64]*/
	type AUS struct {
		Date          float64
		Location      float64
		MinTemp       float64
		MaxTemp       float64
		Rainfall      float64
		Evaporation   float64
		Sunshine      float64
		WindGustDir   float64
		WindGustSpeed float64
		WindDir9am    float64
		WindDir3pm    float64
		WindSpeed9am  float64
		WindSpeed3pm  float64
		Humidity9am   float64
		Humidity3pm   float64
		Pressure9am   float64
		Pressure3pm   float64
		Cloud9am      float64
		Cloud3pm      float64
		Temp9am       float64
		Temp3pm       float64
		RainToday     float64
		RainTomorrow  float64
	}
	AUSFile, err := os.OpenFile(namafile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer AUSFile.Close()

	aus := []*AUS{}
	if err := gocsv.UnmarshalFile(AUSFile, &aus); err != nil { // Load clients from file
		panic(err)
	}

	features := make([][]float64, len(aus))
	targets := make([]int, len(aus))
	for i, s := range aus {
		val := reflect.ValueOf(*s)
		for j := 0; j < val.NumField()-1; j++ {
			features[i] = append(features[i], val.Field(j).Float())
		}
		targets[i] = int(s.RainTomorrow)
		/*
			features[i] = make([]float64, 10)
			//fmt.Println(*s)

			features[i][0] = s.Jenis_Kelamin
			features[i][1] = s.Umur
			features[i][2] = s.SIM
			features[i][3] = s.Kode_Daerah
			features[i][4] = s.Sudah_Asuransi
			features[i][5] = s.Umur_Kendaraan
			features[i][6] = s.Kendaraan_Rusak
			features[i][7] = s.Premi
			features[i][8] = s.Kanal_Penjualan
			features[i][9] = s.Lama_Berlangganan*/
	}
	return features, targets
}

func LoadAUSProcessed(namafile string) ([][]float64, []int) {
	/*[Date  float64 Location  float64 MinTemp  float64 MaxTemp  float64 Rainfall  float64 Evaporation
	  Sunshine  float64 WindGustDir  float64 WindGustSpeed  float64 WindDir9am  float64 WindDir3pm
	  WindSpeed9am  float64 WindSpeed3pm  float64 Humidity9am  float64 Humidity3pm
	  Pressure9am  float64 Pressure3pm  float64 Cloud9am  float64 Cloud3pm  float64 Temp9am
	  Temp3pm  float64 RainToday  float64 RainTomorrow  float64]*/
	type AUS struct {
		Date          float64
		Location      float64
		MinTemp       float64
		MaxTemp       float64
		Rainfall      float64
		Evaporation   float64
		Sunshine      float64
		WindGustDir   float64
		WindGustSpeed float64
		WindDir9am    float64
		WindDir3pm    float64
		WindSpeed9am  float64
		WindSpeed3pm  float64
		Humidity9am   float64
		Humidity3pm   float64
		Pressure9am   float64
		Pressure3pm   float64
		Cloud9am      float64
		Cloud3pm      float64
		Temp9am       float64
		Temp3pm       float64
		RainToday     float64
		RainTomorrow  float64
	}
	AUSFile, err := os.OpenFile(namafile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer AUSFile.Close()

	aus := []*AUS{}
	if err := gocsv.UnmarshalFile(AUSFile, &aus); err != nil { // Load clients from file
		panic(err)
	}

	features := make([][]float64, len(aus))
	targets := make([]int, len(aus))
	for i, s := range aus {
		val := reflect.ValueOf(*s)
		for j := 0; j < val.NumField()-1; j++ {
			features[i] = append(features[i], val.Field(j).Float())
		}
		targets[i] = int(s.RainTomorrow)

		/*
			features[i] = make([]float64, 10)
			//fmt.Println(*s)

			features[i][0] = s.Jenis_Kelamin
			features[i][1] = s.Umur
			features[i][2] = s.SIM
			features[i][3] = s.Kode_Daerah
			features[i][4] = s.Sudah_Asuransi
			features[i][5] = s.Umur_Kendaraan
			features[i][6] = s.Kendaraan_Rusak
			features[i][7] = s.Premi
			features[i][8] = s.Kanal_Penjualan
			features[i][9] = s.Lama_Berlangganan*/
	}
	return features, targets
}

func LoadKendaraanProcessed(namafile string) ([][]float64, []int) {

	type Kendaraan struct {
		/*
			Jenis_Kelamin,
			Umur,
			SIM,
			Kode_Daerah,
			Sudah_Asuransi,
			Umur_Kendaraan,
			Kendaraan_Rusak,
			Premi,
			Kanal_Penjualan,
			Lama_Berlangganan,
			Tertarik
		*/
		Jenis_Kelamin     float64
		Umur              float64
		SIM               float64
		Kode_Daerah       float64
		Sudah_Asuransi    float64
		Umur_Kendaraan    float64
		Kendaraan_Rusak   float64
		Premi             float64
		Kanal_Penjualan   float64
		Lama_Berlangganan float64
		Tertarik          float64
	}
	kendaraanFile, err := os.OpenFile(namafile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer kendaraanFile.Close()

	kendaraan := []*Kendaraan{}
	if err := gocsv.UnmarshalFile(kendaraanFile, &kendaraan); err != nil { // Load clients from file
		panic(err)
	}

	features := make([][]float64, len(kendaraan))
	targets := make([]int, len(kendaraan))
	for i, s := range kendaraan {
		features[i] = make([]float64, 10)
		targets[i] = int(s.Tertarik)
		val := reflect.ValueOf(*s)
		for j := 0; j < val.NumField()-1; j++ {
			features[i] = append(features[i], val.Field(j).Float())
		}
	}
	return features, targets
}
func LoadKendaraanUnprocessed(namafile string) ([][]float64, []int) {

	type Kendaraan struct {
		/*
			Jenis_Kelamin,
			Umur,
			SIM,
			Kode_Daerah,
			Sudah_Asuransi,
			Umur_Kendaraan,
			Kendaraan_Rusak,
			Premi,
			Kanal_Penjualan,
			Lama_Berlangganan,
			Tertarik
		*/
		Jenis_Kelamin     string
		Umur              int
		SIM               int
		Kode_Daerah       int
		Sudah_Asuransi    int
		Umur_Kendaraan    string
		Kendaraan_Rusak   string
		Premi             int
		Kanal_Penjualan   int
		Lama_Berlangganan int
		Tertarik          int
	}
	kendaraanFile, err := os.OpenFile(namafile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer kendaraanFile.Close()

	kendaraan := []*Kendaraan{}
	if err := gocsv.UnmarshalFile(kendaraanFile, &kendaraan); err != nil { // Load clients from file
		panic(err)
	}

	features := make([][]float64, len(kendaraan))
	targets := make([]int, len(kendaraan))
	for i, s := range kendaraan {
		features[i] = make([]float64, 10)
		targets[i] = s.Tertarik
		if s.Jenis_Kelamin == "Pria" {
			features[i][0] = 0
		}
		if s.Jenis_Kelamin == "Wanita" {
			features[i][0] = 0
		}
		features[i][1] = float64(s.Umur)
		features[i][2] = float64(s.SIM)
		features[i][3] = float64(s.Kode_Daerah)
		features[i][4] = float64(s.Sudah_Asuransi)

		if s.Umur_Kendaraan == "< 1 Tahun" {
			features[i][5] = 0
		}
		if s.Umur_Kendaraan == "1-2 Tahun" {
			features[i][5] = 1
		}
		if s.Umur_Kendaraan == "> 2 Tahun" {
			features[i][5] = 2
		}
		if s.Kendaraan_Rusak == "Pernah" {
			features[i][6] = 1
		}
		if s.Kendaraan_Rusak == "Tidak" {
			features[i][6] = 0
		}

		features[i][7] = float64(s.Premi)
		features[i][8] = float64(s.Kanal_Penjualan)
		features[i][9] = float64(s.Lama_Berlangganan)
	}
	return features, targets
}

func main() {
	/*
		dataset := [][]float64{{2.7810836, 2.550537003},
			{1.465489372, 2.362125076},
			{3.396561688, 4.400293529},
			{1.38807019, 1.850220317},
			{3.06407232, 3.005305973},
			{7.627531214, 2.759262235},
			{5.332441248, 2.088626775},
			{6.922596716, 1.77106367},
			{8.675418651, -0.242068655},
			{7.673756466, 3.508563011},
		}
		target := []int{
			0, 0, 0, 0, 0, 1, 1, 1, 1, 1,
		}  */
	//*
	//n_outputs, _ := GoNeuralNetwork.CalculateUnique(target)
	dataset, target := LoadAUSProcessed("train_AUS.csv")
	datatest, expected := LoadAUSProcessed("test_AUS.csv")
	//fmt.Println(dataset, target)*/

	n_input := len(dataset[0])

	network := GoNeuralNetwork.InitializeNetwork(n_input)

	network.AddSoftMaxLayer(12)
	network.AddTanhLayer(10)
	network.AddTanhLayer(6)
	//network.AddSoftMaxLayer(12)
	//network.AddSoftMaxLayer(8)
	network.AddSoftMaxLayer(4)
	//network.AddRelULayer(2)
	network.AddDenseLayer(2)
	//network.AddLeakyRelULayer(2)

	//network.AddLayer(GoNeuralNetwork.CreateLayer(4, 2, GoNeuralNetwork.Softmax, GoNeuralNetwork.DerivativeSoftMax))
	//network.Show()

	//network.Show()
	network.Training(dataset, target, 0.5, 20)
	//network.Show()
	predict := network.Predict(datatest)
	fmt.Println(predict)
	fmt.Println(GoNeuralNetwork.AccuracyScore(predict, expected))
	fmt.Println(GoNeuralNetwork.ConfusionMatrix(predict, expected))
	//fmt.Println(predict)
	//fmt.Println(target)
	/*
		network.TrainNetwork(&network, dataset, target, 0.5, 20, n_outputs)

		predict := classifications.Predict(&network, dataset)

		fmt.Println(classifications.AccuracyScoreFloat64(target, predict))*/

}
