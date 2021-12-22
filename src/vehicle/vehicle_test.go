package vehicle

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type VehicleTestSuite struct {
	suite.Suite
}

// func (s *VehicleTestSuite) SetupSuite() {
// }

func (s *VehicleTestSuite) TestCalculateMotorSpeed() {
	// testVehicle := LRDriveVehicle{
	// 	LeftMotor: ,
	// }
}

func TestVehicleTestSuite(t *testing.T) {
	suite.Run(t, new(VehicleTestSuite))
}
