package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/entraktest/internal/models"
	repomock "github.com/entraktest/internal/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func Test_Store(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	deviceRepository := repomock.NewDeviceRepositoryMock(mockCtrl)
	svc := NewDeviceService(deviceRepository)

	xml := `
	<root>
   <data>
      <element>
         <device>SGD-12344</device>
         <value>1234.266</value>
      </element>
   </data>
   <devices>
      <SGD-12344>Temperature Sensor</SGD-12344>
   </devices>
   <id>2314</id>
   <record_time>1008910273</record_time>
</root>`

	xmlFail := `
	<root>
   <data>
      <element>
         <device>SGD-12344</device>
         <value>1234.266</value>
      </element>
   </data>
   <devices>
      <G3112>Temperature Sensor</G3112>
   </devices>
   <id>2314</id>
   <record_time>1008910273</record_time>
</root>`

	t.Run("Success Inserted", func(t *testing.T) {
		device := models.Device{
			ID:       0,
			Device:   "SGD-12344",
			Value:    1234.266,
			Unit:     "C",
			Datetime: time.Unix(1008910273, 0),
		}
		deviceRepository.EXPECT().Store(device).Return(nil).Times(1)
		err := svc.Store([]byte(xml))
		assert.Equal(t, err, nil)
	})

	t.Run("Fail Inserted", func(t *testing.T) {
		err := svc.Store([]byte(xmlFail))
		fmt.Println(err)
		assert.Equal(t, err, fmt.Errorf("Unmatch data request"))
	})
}

func Test_GetById(t *testing.T) {

}

func Test_GetAll(t *testing.T) {

}
