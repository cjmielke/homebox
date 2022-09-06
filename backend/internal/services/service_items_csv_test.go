package services

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"testing"
)

const CSV_DATA = `
Location,Labels,Name,Description,Serial Number,Mode Number,Manufacturer,Notes,Purchase From,Purchased Price,Purchased At,Sold To,Sold Price,Sold At
Garage,IOT;Home Assistant; Z-Wave,Zooz Universal Relay ZEN17,"Zooz 700 Series Z-Wave Universal Relay ZEN17 for Awnings, Garage Doors, Sprinklers, and More | 2 NO-C-NC Relays (20A, 10A) | Signal Repeater | Hub Required (Compatible with SmartThings and Hubitat)",,ZEN17,Zooz,,Amazon,39.95,10/13/2021,,,
Living Room,IOT;Home Assistant; Z-Wave,Zooz Motion Sensor,"Zooz Z-Wave Plus S2 Motion Sensor ZSE18 with Magnetic Mount, Works with Vera and SmartThings",,ZSE18,Zooz,,Amazon,29.95,10/15/2021,,,
Office,IOT;Home Assistant; Z-Wave,Zooz 110v Power Switch,"Zooz Z-Wave Plus Power Switch ZEN15 for 110V AC Units, Sump Pumps, Humidifiers, and More",,ZEN15,Zooz,,Amazon,39.95,10/13/2021,,,
Downstairs,IOT;Home Assistant; Z-Wave,Ecolink Z-Wave PIR Motion Sensor,"Ecolink Z-Wave PIR Motion Detector Pet Immune, White (PIRZWAVE2.5-ECO)",,PIRZWAVE2.5-ECO,Ecolink,,Amazon,35.58,10/21/2020,,,
Entry,IOT;Home Assistant; Z-Wave,Yale Security Touchscreen Deadbolt,"Yale Security YRD226-ZW2-619 YRD226ZW2619 Touchscreen Deadbolt, Satin Nickel",,YRD226ZW2619,Yale,,Amazon,120.39,10/14/2020,,,
Kitchen,IOT;Home Assistant; Z-Wave,Smart Rocker Light Dimmer,"UltraPro Z-Wave Smart Rocker Light Dimmer with QuickFit and SimpleWire, 3-Way Ready, Compatible with Alexa, Google Assistant, ZWave Hub Required, Repeater/Range Extender, White Paddle Only, 39351",,39351,Honeywell,,Amazon,65.98,09/30/0202,,,
`

func loadcsv() [][]string {
	reader := csv.NewReader(bytes.NewBuffer([]byte(CSV_DATA)))

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func Test_csvRow_getLabels(t *testing.T) {
	type fields struct {
		Labels string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "basic test",
			fields: fields{
				Labels: "IOT;Home Assistant;Z-Wave",
			},
			want: []string{"IOT", "Home Assistant", "Z-Wave"},
		},
		{
			name: "no labels",
			fields: fields{
				Labels: "",
			},
			want: []string{},
		},
		{
			name: "single label",
			fields: fields{
				Labels: "IOT",
			},
			want: []string{"IOT"},
		},
		{
			name: "trailing semicolon",
			fields: fields{
				Labels: "IOT;",
			},
			want: []string{"IOT"},
		},

		{
			name: "whitespace",
			fields: fields{
				Labels: " IOT;		Home Assistant;   Z-Wave ",
			},
			want: []string{"IOT", "Home Assistant", "Z-Wave"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := csvRow{
				Labels: tt.fields.Labels,
			}
			if got := c.getLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csvRow.getLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
