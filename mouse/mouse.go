package mouse

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/shoaibahmed997/automata/utils"
)

type MouseMovements struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Timestamp string `json:"timestamp"`
}

func endMouseRecording(recordingMouse *bool) {
	keve := hook.AddEvent("k") // add more resilient shortcuts
	if keve {
		fmt.Println("you press... ", "k")
		*recordingMouse = false
	}
}

func RecordMouse(w http.ResponseWriter, r *http.Request) {
	recordingMouse := true
	fmt.Println("recording mouse... press t to stop recording")
	allMouseMovements := []MouseMovements{}
	// keep recording the mouse movements at every 500 miliseconds
	go endMouseRecording(&recordingMouse)
	for recordingMouse {
		x, y := robotgo.Location()
		currMouseMovement := MouseMovements{X: x, Y: y, Timestamp: time.TimeOnly}
		allMouseMovements = append(allMouseMovements, currMouseMovement)
		robotgo.MilliSleep(300)
	}
	jsonData, err := json.Marshal(allMouseMovements)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	utils.WriteJsonToFile("mouserecords.json", jsonData)

	// // save the mouse movements to a json file.
	fmt.Println("recoding mouse var", recordingMouse)
	fmt.Fprintf(w, "Record Mouse")
}

func MouseTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting test")
	// // ENTER TESTS HERE
	// read from the file
	// perform those actions
	// send the success value to frontend
	filename := "mouseRecords.json"
	jsonData := utils.ReadJsonFromFile(filename)
	// unmarshal json data
	var recordedMouseMovements []MouseMovements
	jsonErr := json.Unmarshal([]byte(jsonData), &recordedMouseMovements)
	if jsonErr != nil {
		fmt.Println("error parsing json data", jsonErr)
	}
	for _, item := range recordedMouseMovements {
		robotgo.MoveSmooth(item.X, item.Y)
	}
	fmt.Fprintf(w, "Successful")

}

func ReturnAllMacros(w http.ResponseWriter, r *http.Request) {
	allmacros := utils.AllMacros()
	wd, _ := os.Getwd()
	temp := template.Must(template.ParseFiles(wd + "/static/templates/allmacros.html"))
	fmt.Println(temp)
	temp.Execute(w, allmacros)
	// fmt.Fprintf(w, "allmacros")

}
