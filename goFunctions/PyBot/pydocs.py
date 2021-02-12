#!/usr/bin/python3
##############
## Script listens to serial port and writes contents into a file
##############
## requires pySerial to be installed
"""
import serial
arrayArduino = []

def Start():
    serial_port = 'COM5';
    baud_rate = 9600; #In arduino, Serial.begin(baud_rate)
    write_to_file_path = "values.txt";

    output_file = open(write_to_file_path, "w+");
    ser = serial.Serial(serial_port, baud_rate)
    inicio = False
    C = 1  # CONTADOR DE SUCESOS DE  LAS MEDICIONES

    while not (inicio):

        if (len(arrayArduino)<8):
            line = ser.readline()
            conversion = str(line.decode('ascii', errors='replace'))
            line = line.decode("utf-8") #ser.readline returns a binary, convert to string
            print(conversion, end="")
            output_file.write(conversion)
            arrayArduino.append(conversion)
            C = C + 1
        else:
            break


######################3


def Arduino():
    Registro = {}
    Registro2 = []
    inicio = False
    arduino = serial.Serial(port='COM4', baudrate=int(9600), timeout=2)
    # arduino2 = serial.Serial(port= 'COM4', baudrate =int(9600), timeout=2)
    C = 1  # CONTADOR DE SUCESOS DE  LAS MEDICIONES

    while not (inicio or C > 20):
        line = arduino.readline()  # EN ESTA PARTE SE LEE TODA LA INFORMACION QUE NOS TRASNMITE LA PLACA ARDUINO UNO
        conversion = str(line.decode('ascii', errors='replace'))
        # conversion = str(line)
        conversion2 = str(line)
        # EntradaDeDatos1 = str(conversion.set())

        objeto = Archivo.Variables_de_precipitacion(conversion)
        # objeto2 = Archivo.Variables_de_precipitacion(conversion2)
        presentacion = objeto.ObtenerDatos()
        # presentacion2 = objeto2.MedirAgua()
        # print("MEDIDOR EN VALORES DE ARDUINO INDICA.....", conversion)
        Registro[C] = conversion
        Registro2.append(conversion)
        # print(Reistro)
        # print(presentacion2)
        # print(conversion2)
        # print(line.decode('ascii', errors='replace'), end='')
        C = C + 1
        # Reistro.clear()
    # print(Registro)

    # print(len(Registro2))
    posicion = 0
    while posicion < len(Registro2):
        if Registro2[posicion] == "":
            Registro2.pop(posicion)
        else:
            posicion = posicion + 1

    Registro2 = [x.replace("\r\n", "") for x in Registro2]
    # print("Asi queda la lista A:  ", Registro2)
    # =======================================================================================
    for i in Registro2:
        playlistbox.insert(0, i)

    # datos = np.arange(0, 100)
    # plt.plot([1, 2, 3, 4,5,6,7,8,9,10])
    plt.ylabel('PAREMETERS ARDUINO')
    # print(A1)
    plt.plot(Registro2)
    plt.show()








"""


