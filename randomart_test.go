package randomart

import (
	"encoding/hex"
	"testing"
)

func TestGenerate(t *testing.T) {

	// From http://sshvis.appspot.com/

	var tests = []struct {
		input  string
		output string
	}{
		{"f5f8a4b7ac83b0b9aa9ad4306269e738", "" +
			"+-----------------+\n" +
			"|                 |\n" +
			"|                 |\n" +
			"|          .      |\n" +
			"|  .      . o     |\n" +
			"|.* .    S . o    |\n" +
			"|+ B   .    +     |\n" +
			"| E o   + .. o    |\n" +
			"|...   o . .o .   |\n" +
			"|o.......  .oo    |\n" +
			"+-----------------+",
		},
		{"5a46a29f198bdb5a988ccef224acafbf", "" +
			"+-----------------+\n" +
			"|                 |\n" +
			"|                 |\n" +
			"|      . .        |\n" +
			"|     . o         |\n" +
			"|    . . S        |\n" +
			"|.  o = O         |\n" +
			"|..o = B          |\n" +
			"|o=   +           |\n" +
			"|o*E.o..          |\n" +
			"+-----------------+",
		},
		{"fc94b0c1e5b0987c5843997697ee9fb7", "" +
			"+-----------------+\n" +
			"|       .=o.  .   |\n" +
			"|     . *+*. o    |\n" +
			"|      =.*..o     |\n" +
			"|       o + ..    |\n" +
			"|        S o.     |\n" +
			"|         o  .    |\n" +
			"|          .  . . |\n" +
			"|              o .|\n" +
			"|               E.|\n" +
			"+-----------------+",
		},
	}

	for _, tt := range tests {
		data, _ := hex.DecodeString(tt.input)

		sshhash := NewSSH()

		sshhash.Write(data)

		s := sshhash.String()

		if s != tt.output {
			t.Errorf("ssh visual hash test failed: got:\n%s\nwanted:\n%s", s, tt.output)
		}
	}

}
