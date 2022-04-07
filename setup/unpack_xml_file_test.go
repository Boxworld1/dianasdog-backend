// @Title  unpackfile
// @Description  Read XML files in folder and unpack them
// @Author  于沛楠
// @Update  2022/3/16
package setup

import (
	"strconv"
	"testing"
)

// test for function: unpackXMLFile
func TestUnpackXMLFile(t *testing.T) {
	itemList, itemCount, _, resourceName, err := UnpackXMLFile("../data/testcase/testcase_car.xml", "car")
	carKey := []string{"集度汽车新能源", "indi新能源", "北京汽车新能源", "蓝旗亚新能源", "wayray新能源"}

	if err != nil { //wrong error
		t.Error(err)
	}
	if itemCount != 5 { //wrong itemCount
		t.Error("Wrong count of cars: " + strconv.Itoa(itemCount))
	}
	if resourceName != "car" {
		t.Error("Wrong resourcename of cars: " + resourceName)
	}
	key := ""
	for i := 0; i < itemCount; i++ {
		key = itemList[i].SelectElement("key").Text()
		if key != carKey[i] { //wrong content of itemList
			t.Error("Wrong unpack of ./car/car_test.xml")
		}
	}

	itemList, itemCount, _, _, err = UnpackXMLFile("./cars/car_test.xml", "car")
	if err == nil || itemCount != 0 || itemList != nil { // error miss
		t.Error("read wrong filename but not send error")
	}
}
