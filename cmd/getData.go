package data

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/simonvetter/modbus"
)

func Float32frombytesLittle(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func GetData() map[string]float32 {
	var client *modbus.ModbusClient
	var err error
	var regBs []byte
	var i uint16 = 0
	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://localhost:502",
		Timeout: 1 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	err = client.Open()
	if err != nil {
		fmt.Println(err)
	}
	client.SetEncoding(modbus.LITTLE_ENDIAN, modbus.LOW_WORD_FIRST)
	dataMap := make(map[string]float32)
	keys := []string{
		"tic_100",
		"pump_rpm",
		"pump_power",
		"pump_flow_sp",
		"pump_flow",
		"pic_101",
		"pic_100",
		"lic_102",
		"lic_101",
		"fic_100",
	}
	for _, key := range keys {
		regBs, err = client.ReadBytes(i, 4, modbus.INPUT_REGISTER)
		if err != nil {
			break
		}
		// fmt.Printf("LITTLE:\n\tvalue: %v\taddress: %v\n", Float32frombytesLittle(regBs), i)
		dataMap[key] = Float32frombytesLittle(regBs)
		i += 2
	}
	// fmt.Printf("%v\n", dataMap)
	client.Close()
	return dataMap
}
