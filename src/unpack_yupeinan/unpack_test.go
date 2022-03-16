package unpack

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	resources = append(resources, "car")
	resources = append(resources, "poem")
	resources = append(resources, "cars")
	resources = append(resources, "poems")

	for i := 0; i < len(resources); i++ {
		filesPath, err := GetFiles("../" + resources[i])
		if err != nil {
			//if (err.Error() != "open ../" + resources[i] + ": The system cannot find the file specified.") {
			t.Error(err)
			//}
		} else {
			for j := 0; j < len(filesPath); j++ {
				err = Unpack(filesPath[j], resources[i])
				if err != nil {
					t.Error(err)
				}
			}
		}
	}

}
