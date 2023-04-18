package cmdstr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type CommandRequest struct {
	Command string `json:"command"`
}

func main() {
	http.HandleFunc("/run", handleRunCommand)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRunCommand(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request CommandRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := exec.Command("bash", "-c", request.Command).CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", output)
}
