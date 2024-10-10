package gis

import (
	"fmt"
	"testing"
)

func TestInsertGisData(t *testing.T) {
	province := "ACEH"
	district := "ACEH SELATAN"
	subDistrict := "BAKONGAN"
	village := "DARUL IKHSAN"
	border := [][]float64{
		{97.4447, 2.9686},
		{97.4384, 2.9737},
		{97.4253, 2.9843},
		{97.3967, 2.9596},
		{97.3988, 2.9576},
		{97.3989, 2.9574},
		{97.4002, 2.9565},
		{97.4058, 2.9514},
		{97.4116, 2.9462},
		{97.4147, 2.9433},
		{97.4177, 2.9404},
		{97.4186, 2.9397},
		{97.4259, 2.9457},
		{97.434, 2.9543},
		{97.4447, 2.9686},
	}

	hasil, err := InsertGisData(province, district, subDistrict, village, border)
	if err != nil {
		t.Errorf("Failed to insert GisData: %v", err)
	} else {
		fmt.Printf("Inserted GisData ID: %v\n", hasil)
	}
}
