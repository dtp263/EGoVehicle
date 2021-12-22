package vehicle

import (
	"fmt"

	"github.com/dtp263/bessie/src/motor"
	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Vehicle interface {
	Drive(direction, speed int) error
}

type LRDriveVehicle struct {
	LeftMotor  motor.Motor
	RightMotor motor.Motor
}

type DriveArgs struct {
	XSpeed int `validate:"required,min=-1000,max=1000"`
	YSpeed int `validate:"required,min=-1000,max=1000"`
}

func (v LRDriveVehicle) DriveMotor(args DriveArgs, motor string) {

}

func (v LRDriveVehicle) Drive(args DriveArgs) error {
	validate.Struct(args)

	leftMotorSpeed := args.YSpeed * args.XSpeed
	rightMotorSpeed := (args.YSpeed * -1) * args.XSpeed

	fmt.Println(leftMotorSpeed)
	fmt.Println(rightMotorSpeed)

	// for left motor
	// if args.XSpeed > 0 {

	// }
	// const  args.

	v.LeftMotor.SetPower(leftMotorSpeed)
	v.RightMotor.SetPower(rightMotorSpeed)

	return nil
}
