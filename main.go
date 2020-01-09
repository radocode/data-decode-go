package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("Ingrese cadena a evaluar: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		res, _ := tlvParser([]byte(scanner.Text()))
		for r, i := range res {
			fmt.Println(r, i)
		}
	}
}

func tlvParser(tlv []byte) (map[string]string, error) {

	result := make(map[string]string)

	if tlv == nil {
		return nil, errors.New("Debe ingresar valor a parsear")
	}
	if len(tlv) == 0 {
		return nil, errors.New("Debe ingresar valor a parsear")
	}

	var b bytes.Buffer
	b.Write(tlv)

	for index := 0; len(b.String()) > 0; index++ {
		// obtener largo, 2 primeros caracteres, importante validar si es entero!
		largo := b.Next(2)
		largoInt, _ := strconv.Atoi(string(largo))
		v := reflect.ValueOf(largoInt)
		if v.Kind() != reflect.Int {
			return nil, errors.New("Parametro Largo debe tener valor numerico")
		}
		result["Largo"+string(index)] = string(largo)

		// obtener tipo, primero identificar tipo de dato y luego el largo
		tipoDato := string(b.Next(1))
		if !IsValidTipoData(tipoDato) {
			return nil, errors.New("Parametro Tipo debe ser tipificado A (alfanumerico) o N (numerico)")
		}
		result["Tipo"+string(index)] = getTypeText(tipoDato) + " de largo " + string(b.Next(2))

		// obtener valor, el resto del byte array
		result["Valor"+string(index)] = string(b.Next(largoInt))
	}

	return result, nil
}

func getTypeText(tipo string) string {
	switch tipo {
	case "A":
		return "Alfanumerico"
	case "N":
		return "Numerico"
	default:
		return ""
	}
}

// IsValidTipoData verifica tipo de data
func IsValidTipoData(tipo string) bool {
	switch tipo {
	case
		"A",
		"N":
		return true
	}
	return false
}
