/*scriptino per spostare i file dalla cartella da cui li scarico a quella del sito*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("ls", "/home/marco/Scaricati/Immagini da email per sito/")
	imgs, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	imgsArray := strings.Fields(string(imgs))

	for _, img := range imgsArray {
		origin := "/home/marco/Scaricati/Immagini da email per sito/" + img

		endpoint := "/home/marco/Documenti/marcozeccola/uploads/" + img

		os.Rename(origin, endpoint)
	}
}
