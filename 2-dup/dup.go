/*!
 * \file dup.go
 * \brief Affiche le nombre de fois et le contenu
 *	      des lignes qui apparaîssent plus d'une fois
 *		  dans des fichiers.
 *		  Découverte de "io/ioutil" et de la fonction ReadFile
 * \author lhm
 */

 package main

 import (
 	"fmt"
 	"io/ioutil"
 	"os"
 	"strings"
 )

 func main() {
 	inFiles := make(map[string]map[string]int)	
 	for _,fileName := range os.Args[1:] {		// Parcours des fichiers
 		data,err   := ioutil.ReadFile(fileName)

 		// Gestion erreur chargement fichier
 		if err != nil {
 			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
 			continue
 		}

 		// Transforme data (byte slice) en string splice
 		for _,line := range strings.Split(string(data), "\n") {
 			if inFiles[line] == nil {
 				inFiles[line] = make(map[string]int)
 			}
 			inFiles[line][fileName]++
 		}
 	}

 	for line, filenames := range inFiles {
		fileCount := len(filenames)
		total := 0
		for _, count := range filenames {
			total += count
		}
		if total <= 1 {
			continue
		}

		fmt.Printf("[Found in %d file(s)]\t%s\n", fileCount, line)
		for name, count := range filenames {
			fmt.Printf("\t%d hit(s) in %s\n", count, name)
		}
	}
}