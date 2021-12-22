package motor

import "github.com/sirupsen/logrus"

type Motor interface {
	SetPower(power int) error
}

type KunrayMotor struct{}

func (m KunrayMotor) SetPower(power int) error {
	logrus.Infof("set motor power to %d", power)
	return nil
}
