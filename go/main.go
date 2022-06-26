package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var strLen = 25   // Length of string to generate
var strQty = 1    // Number of strings to generate
var noSym = 0     // 1 = Don't include symbols
var alphaOnly = 0 // 1 = Only include alpha characters
var numOnly = 0   // 1 = Only include numeric characters
var noForm = 0    // 1 = Don't display the form, just the resulting string

func makeString() string {
	//
	returnString := ""
	//
	for len(returnString) < strLen {
		b := make([]byte, strLen*3)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("error:", err)
			return ""
		}

		okChar := 1
		for i := 0; i < strLen; i++ {
			chrOrd := b[i]
			if (chrOrd > 32) && (chrOrd < 128) {
				okChar = 1

				if (noSym == 0) && (alphaOnly == 0) && (numOnly == 0) {
					if chrOrd == 32 {
						okChar = 0
					} // space
					//if (chrOrd == 96) 	{ okChar = 0 }	// `
					if chrOrd == 34 {
						okChar = 0
					} // "
					//if (chrOrd == 39) 	{ okChar = 0 }	// '
					if chrOrd == 9 {
						okChar = 0
					} // tab
					if chrOrd == 10 {
						okChar = 0
					} // cr
					//if (chrOrd == 124) 	{ okChar = 0 }	// |
					//if (chrOrd == 37) 	{ okChar = 0 }	// %
					//if (chrOrd == 40) 	{ okChar = 0 }	// (
					//if (chrOrd == 41) 	{ okChar = 0 }	// )
					//if (chrOrd == 47) 	{ okChar = 0 }	// /
					//if (chrOrd == 58) 	{ okChar = 0 }	// :
					//if (chrOrd == 60) 	{ okChar = 0 }	// <
					//if (chrOrd == 62) 	{ okChar = 0 }	// >
				} else if alphaOnly == 1 {
					okChar = 0
					if (chrOrd >= 65) && (chrOrd <= 90) {
						okChar = 1
					}
					if (chrOrd >= 97) && (chrOrd <= 122) {
						okChar = 1
					}
				} else if numOnly == 1 {
					okChar = 0
					if (chrOrd >= 48) && (chrOrd <= 57) {
						okChar = 1
					}
				} else {
					okChar = 0
					if (chrOrd >= 48) && (chrOrd <= 57) {
						okChar = 1
					}
					if (chrOrd >= 65) && (chrOrd <= 90) {
						okChar = 1
					}
					if (chrOrd >= 97) && (chrOrd <= 122) {
						okChar = 1
					}
				}
				//
				if okChar == 1 {
					returnString = returnString + string(chrOrd)
				}
			}
		}
	}

	return returnString[:strLen]
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//
	/* Reset variables each run to ensure remnants don't hang around */
	strLen = 25
	strQty = 1
	noSym = 0
	alphaOnly = 0
	numOnly = 0
	noForm = 0
	//
	normModeChecked := "checked"
	noSymChecked := ""     // Will contain 'checked' if set to 1, for the HTML form
	alphaOnlyChecked := "" // Will contain 'checked' if set to 1, for the HTML form
	numOnlyChecked := ""   // Will contain 'checked' if set to 1, for the HTML form
	//
	pwString := ""
	//
	if len(r.URL.Query().Get("mode")) > 0 {
		mode := r.URL.Query().Get("mode")
		if mode == "alphaonly" {
			alphaOnlyChecked = "checked"
			normModeChecked = ""
			alphaOnly = 1
		} else if mode == "numonly" {
			numOnlyChecked = "checked"
			normModeChecked = ""
			numOnly = 1
		} else if mode == "nosym" {
			noSym = 1
			noSymChecked = "checked"
			normModeChecked = ""
		}
	}
	//
	if len(r.URL.Query().Get("alphaonly")) > 0 {
		alphaOnly, _ = strconv.Atoi(r.URL.Query().Get("alphaonly"))
		if alphaOnly == 1 {
			alphaOnlyChecked = "checked"
		}
	}
	if len(r.URL.Query().Get("numonly")) > 0 {
		numOnly, _ = strconv.Atoi(r.URL.Query().Get("numonly"))
		if numOnly == 1 {
			numOnlyChecked = "checked"
		}
	}
	if len(r.URL.Query().Get("nosym")) > 0 {
		noSym, _ = strconv.Atoi(r.URL.Query().Get("nosym"))
		if noSym == 1 {
			noSymChecked = "checked"
		}
	}
	if len(r.URL.Query().Get("qty")) > 0 {
		strQty, _ = strconv.Atoi(r.URL.Query().Get("qty"))
	}
	if len(r.URL.Query().Get("len")) > 0 {
		strLen, _ = strconv.Atoi(r.URL.Query().Get("len"))
	}
	if len(r.URL.Query().Get("noform")) > 0 {
		noForm, _ = strconv.Atoi(r.URL.Query().Get("noform"))
	}

	for i := 1; i <= strQty; i++ {
		pwString = pwString + makeString() + "\n"
	}

	if noForm == 0 {
		formString := ""
		//
		formString = formString + "<html><body>\n"
		formString = formString + "<form method='GET' action='/' name='MAIN'>\n"
		formString = formString + "<table align='center' border='0'>\n"
		formString = formString + "<tr><td align='center'><b>Params:</b></td><td><table border='0'>\n"
		formString = formString + "<tr><td>Quantity: <td><td><input type='text' name='qty' size='5' value='" + strconv.Itoa(strQty) + "'></td></tr>\n"
		formString = formString + "<tr><td>Length: <td><td><input type='text' name='len' size='5' value='" + strconv.Itoa(strLen) + "'></td></tr>\n"
		formString = formString + "<tr><td>All Characters: </td><td><input type='radio' name='mode' value='norm' " + normModeChecked + "></td></tr>\n"
		formString = formString + "<tr><td>No Symbols: </td><td><input type='radio' name='mode' value='nosym' " + noSymChecked + "></td></tr>\n"
		formString = formString + "<tr><td>Alpha Only: </td><td><input type='radio' name='mode' value='alphaonly' " + alphaOnlyChecked + "></td></tr>\n"
		formString = formString + "<tr><td>Num Only: </td><td><input type='radio' name='mode' value='numonly' " + numOnlyChecked + "></td></tr>\n"
		formString = formString + "<tr><td colspan='2'><center><input type='submit' value='Generate New'></td></tr>\n"
		formString = formString + "</table></td></tr>\n"
		formString = formString + "<tr><td align='center'><b>Password List:</b></td><td><textarea cols=" + strconv.Itoa(strLen+10) + " rows=" + strconv.Itoa(strQty+5) + ">" + pwString + "</textarea></td></tr>\n"
		formString = formString + "</table>\n"
		formString = formString + "</form>\n"
		formString = formString + "</body></html>\n"
		w.Write([]byte(formString))
	} else {
		w.Write([]byte(pwString))
	}
}

func main() {
	// Declare a new router
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler).Methods("GET")
	// http.Handle("/", r)

	http.ListenAndServe(":80", r)
}
