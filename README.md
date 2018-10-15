# numletter
Programa en go que dado un número no negativo, el nombre de moneda y céntimos, desplegara el número en letras al español y un tipo de número flotante con metodo convertidor al español e ingles ("es" y "en" ) respectivamente, ambos con un limite del número de 999,999,999

Por ejemplo:

Num2Letter(987654.32, "Pesos", "Centavos", false)

debe devolver:

NOVECIENTOS OCHENTA Y SIETE MIL SEISCIENTOS CINCUENTA Y CUATRO PESOS CON TREINTA Y DOS CENTAVOS 

El ultimo argumento booleano sirve para cuando no hay decimales pero se quiere incluir que son cero céntimos 
