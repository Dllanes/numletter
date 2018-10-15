package numletter

import(
	"fmt"
	"strings"
	"strconv"
)

var Unidades = []string{
	"",
	"UN ",
	"DOS ",
	"TRES ",
	"CUATRO ",
	"CINCO ",
	"SEIS ",
	"SIETE ",
	"OCHO ",
	"NUEVE ",
	"DIEZ ",
	"ONCE ",
	"DOCE ",
	"TRECE ",
	"CATORCE ",
	"QUINCE ",
	"DIECISEIS ",
	"DIECISIETE ",
	"DIECIOCHO ",
	"DIECINUEVE ",
	"VEINTE ",
}

var Decenas = []string{
	"VENTI",
    "TREINTA ",
    "CUARENTA ",
    "CINCUENTA ",
    "SESENTA ",
    "SETENTA ",
    "OCHENTA ",
    "NOVENTA ",
}

var Centenas = []string{
	"CIENTO ",
    "DOSCIENTOS ",
    "TRESCIENTOS ",
    "CUATROCIENTOS ",
    "QUINIENTOS ",
    "SEISCIENTOS ",
    "SETECIENTOS ",
    "OCHOCIENTOS ",
    "NOVECIENTOS ",
}

var Units = []string{
	"ONE ",
	"TWO ",
	"THREE ",
	"FOUR ",
	"FIVE ",
	"SIX ",
	"SEVEN ",
	"EIGHT ",
	"NINE ",
	"TEN ",
	"ELEVEN ",
	"TWELVE ",
	"THIRTEEN ",
	"FOURTEEN ",
	"FIFTEEN ",
	"SIXTEEN ",
	"SEVENTEEN ",
	"EIGHTEEN ",
	"NINETEEN ",
}

var Teens = []string{
	"TWENTY",
    "THIRTY ",
    "FORTY ",
    "FIFTY ",
    "SIXTY ",
    "SEVENTY ",
    "EIGHTY ",
    "NINETY ",
}

type Feria float64

func (num Feria) Trans(lenguaje string) string {
	switch lenguaje {
	case "es":
		return FloatLetterEs(strconv.FormatFloat(float64(num), 'f', 2, 64) )
	case "en":
		return FloatLetterEn(strconv.FormatFloat(float64(num), 'f', 2, 64) )
	default:
		return "No tenemos conversion a ese lenguaje"	
	}
}

func FloatLetterEn(numero string) string {
	
	convertido := ""
	decimales := ""

	div_decimales := strings.Split(numero,".")

	
	numero = div_decimales[0]
	decNumberStr := div_decimales[1]
	if decNumberStr != "00"{
		decNumberStrFill := strings.Repeat("0", 7) + decNumberStr
		decCientos := decNumberStrFill[6 : ]
		decimales = ConvertGroup(decCientos, "en")
	}

	var valor_convertido string

	numeroFill := strings.Repeat("0", 9 - len(numero)) + numero

	millions := numeroFill[ : 3]
	thousands := numeroFill[3 : 6]
	hundreads := numeroFill[6 : ]
	
	var s bool

	if millions != "000" {
		convertido += fmt.Sprintf("%sMILLION ", ConvertGroup(millions, "en") )
		s = true
	}

	if thousands != "000" {
		convertido += fmt.Sprintf("%sTHOUSAND ", ConvertGroup(thousands, "en") )
		s = true
	}

	if f, _ := strconv.Atoi(hundreads); f > 0 {
		
		convertido += ConvertGroup(hundreads, "en")
		s = true
	}
	if s {
		valor_convertido = convertido
	}else{
		valor_convertido = "ZERO "
	}

	if decimales != "" {
		valor_convertido += "POINT " + decimales
	}

	return valor_convertido
}



func FloatLetterEs(numero string ) string {

	convertido := ""
	decimales := ""

	div_decimales := strings.Split(numero,".")

	
	numero = div_decimales[0]
	decNumberStr := div_decimales[1]
	if decNumberStr != "00"{
		decNumberStrFill := strings.Repeat("0", 7) + decNumberStr
		decCientos := decNumberStrFill[6 : ]
		decimales = ConvertGroup(decCientos, "es")
	}

	var valor_convertido string

	if l, _ := strconv.Atoi(numero); l%1000000 != 0 {
		numeroFill := strings.Repeat("0", 9 - len(numero)) + numero

		millones := numeroFill[ : 3]
		miles := numeroFill[3 : 6]
		cientos := numeroFill[6 : ]

		if millones != "000" {
			if millones == "001" {
				convertido += "UN MILLON "
			}else {
				convertido += fmt.Sprintf("%sMILLONES ", ConvertGroup(millones, "es") )
			}
		}

		if miles != "000" {

			convertido += fmt.Sprintf("%sMIL ", ConvertGroup(miles, "es") )

		}

		if cientos != "000" {

			convertido += ConvertGroup(cientos, "es")

		}

		valor_convertido = convertido
	}else{
		if l == 0 {
			valor_convertido = "CERO "
		}else{
			numeroFill := strings.Repeat("0", 9 - len(numero)) + numero
			millones := numeroFill[ : 3]

			if f, _ := strconv.Atoi(millones); f > 0 {
				if millones == "001" {
					convertido += "UN MILLON DE "
				}else {
					convertido += fmt.Sprintf("%sMILLONES ", ConvertGroup(millones, "es") )
				}
			}
			valor_convertido = convertido
		}
	}

	if decimales != "" {
		valor_convertido += "CON " + decimales
	}

	return valor_convertido

}

func ConvertGroup( n, lenguaje string ) string {
	final := ""
	if lenguaje == "es"{  
		if n[0] != '0' {
			if n == "100" {
				return "CIEN "
			}
			final += Centenas[n[0]-49]
		}

		k, _ := strconv.Atoi(n[1 : ])

		if k <= 20 {
			final += Unidades[k]
		}else{
			if k > 30 && n[2] != '0' {
				final += Decenas[n[1] - 50] + "Y " + Unidades[n[2] - 48]
			}else{
				final += Decenas[n[1] - 50] + Unidades[n[2] - 48]
			}
		}
	}else{
		if n[0] != '0' {
			final = Units[n[0]-49] + "HUNDREAD "
		}

		k, _ := strconv.Atoi(n[1 : ])

		if k < 20 {
			final += Units[k-1]
		}else{
			final += Teens[n[1] - 50]
			if k%10 != 0 {
				final = final[ : len(final)-1] + "-" + Units[n[2] - 49]
			}
		}
	}

	return final

}
