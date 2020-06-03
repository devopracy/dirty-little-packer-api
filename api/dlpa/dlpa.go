package dlpa

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// This will redirect to the manifest bucket with a lookup for the job.
	http.Redirect(w, r, "", http.StatusFound)
}

func BuildHandler(w http.ResponseWriter, r *http.Request) {
	// Dummy code for now.
	cmd := exec.Command("packer", "--help")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Print(err)
	}
	fmt.Fprintf(w, "%v", &out)
}
