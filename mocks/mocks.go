package mocks

import (
	"github.com/stretchr/testify/mock"
)

// ***************************** Config mock obj ***************

type CfgTheMockedObjIF interface {
	GetSM9CliMode() bool
	GetSM9SmtpHost() string
	GetSM9SmtpPort() string
	GetSM9SmtpFrom() string
	GetBasicAuthUser() string
	GetBasicAuthPass() string
	GetSM9Host() string
	GetSM9Port() string
	GetSM9Token() string
	GetSM9URI() string
	InitConf(file string) error
	GetServerPort() uint16
	GetServerToken() string
}

func GetCfgMockedObj() CfgTheMockedObjIF {
	return &CfgTheMockedObject{}
}

type CfgTheMockedObject struct {
	mock.Mock
}

func (mo *CfgTheMockedObject) GetSM9CliMode() bool {
	args := mo.Called()
	return args.Get(0).(bool)
}

func (mo *CfgTheMockedObject) GetSM9SmtpHost() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) GetSM9SmtpPort() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) GetSM9SmtpFrom() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) GetBasicAuthUser() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) GetBasicAuthPass() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) GetSM9Host() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) GetSM9Port() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) GetSM9Token() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) GetSM9URI() string {
	args := mo.Called()
	return args.Get(0).(string)
}

func (mo *CfgTheMockedObject) InitConf(file string) error {
	args := mo.Called()
	return args.Error(1)
}

func (mo *CfgTheMockedObject) GetServerPort() uint16 {
	args := mo.Called()
	return args.Get(0).(uint16)
}

func (mo *CfgTheMockedObject) GetServerToken() string {
	args := mo.Called()
	return args.Get(0).(string)
}
