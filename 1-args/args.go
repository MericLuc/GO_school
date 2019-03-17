/*!
 * \file args.go
 * \brief Echo command line arguments.
 *		  Découverte package "os".
 * \author lhm
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args1(os.Args)
	args2(os.Args)
	args3(os.Args)
	args4(os.Args)
}

/*!
 * Méthode 1 - Compacte, utilisation du formattage
 * des strings prédéfini par fmt.Println
 */
 func args1(args []string) {
	fmt.Println(args[1:])
}

/*!
 * Méthode 2 - Pas de boucle, simple concaténation
 */
func args2(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

/*!
 * Méthode 3 - Boucle sur tous les éléments de os.Args
 */
func args3(args []string) {
	var l_out, l_sep string
	for i:= 1; i < len(args); i++ {
		l_out += l_sep + args[i]
		l_sep = " "
	}
	fmt.Println(l_out)
}

/*!
 * Méthode 4 - Boucle mais itération sur les indices
 * implicite avec "range" (découverte blank identifier _)
 */
 func args4(args []string) {
 	var l_out, l_sep string
 	for _ , l_args := range args[1:] {
 		l_out += l_sep + l_args
 		l_sep = " "
 	}
 	fmt.Println(l_out)
 }