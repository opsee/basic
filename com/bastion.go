package com

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	BastionStateNew       = "new"
	BastionStateLaunching = "launching"
	BastionStateFailed    = "failed"
	BastionStateActive    = "active"
	BastionStateDisabled  = "disabled"
	BastionStateDeleted   = "deleted"
)

var BastionStates = []string{
	BastionStateNew,
	BastionStateLaunching,
	BastionStateFailed,
	BastionStateActive,
	BastionStateDisabled,
	BastionStateDeleted,
}

type BastionUser struct {
	Username string
	Key      string
}

type Bastion struct {
	Id               string         `json:"id"`
	CustomerId       string         `json:"customer_id" db:"customer_id"`
	ExecutionGroupId string         `json:"execution_group_id" db:"execution_group_id"`
	UserId           int            `json:"user_id" db:"user_id"`
	StackId          sql.NullString `json:"stack_id" db:"stack_id"`
	ImageId          sql.NullString `json:"image_id" db:"image_id"`
	InstanceId       sql.NullString `json:"instance_id" db:"instance_id"`
	GroupId          sql.NullString `json:"group_id" db:"group_id"`
	InstanceType     string         `json:"instance_type" db:"instance_type"`
	SubnetRouting    string         `json:"subnet_routing" db:"subnet_routing"`
	Region           string         `json:"region"`
	VPCId            string         `json:"vpc_id" db:"vpc_id"`
	SubnetId         string         `json:"subnet_id" db:"subnet_id"`
	State            string         `json:"state"`
	Connected        bool           `json:"connected"`
	Password         string         `json:"-"`
	PasswordHash     string         `json:"-" db:"password_hash"`
	CreatedAt        time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at" db:"updated_at"`
}

func NewBastion(userId int, customerId, executionGroupId, region, vpcId, subnetId, subnetRouting, instanceType string) (*Bastion, error) {
	pwbytes := make([]byte, 18)
	if _, err := rand.Read(pwbytes); err != nil {
		return nil, err
	}

	pw := base64.StdEncoding.EncodeToString(pwbytes)
	pwhash, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		return nil, err
	}

	return &Bastion{
		Password:         pw,
		PasswordHash:     string(pwhash),
		CustomerId:       customerId,
		ExecutionGroupId: executionGroupId,
		UserId:           userId,
		InstanceType:     instanceType,
		SubnetRouting:    subnetRouting,
		State:            BastionStateNew,
		Region:           region,
		VPCId:            vpcId,
		SubnetId:         subnetId,
		StackId:          sql.NullString{},
		ImageId:          sql.NullString{},
		InstanceId:       sql.NullString{},
		GroupId:          sql.NullString{},
	}, nil
}

func (bastion *Bastion) StackName() string {
	return "opsee-bastion-" + bastion.Id
}

func (bastion *Bastion) Name() string {
	return "Opsee Bastion " + bastion.Id
}

func (bastion *Bastion) Fail() *Bastion {
	bastion.State = BastionStateFailed
	return bastion
}

func (bastion *Bastion) Launch(stackId, imageId string) *Bastion {
	bastion.State = BastionStateLaunching
	bastion.StackId = sql.NullString{stackId, stackId != ""}
	bastion.ImageId = sql.NullString{imageId, imageId != ""}
	return bastion
}

func (bastion *Bastion) Activate(instanceId, groupId string) *Bastion {
	bastion.State = BastionStateActive
	bastion.InstanceId = sql.NullString{instanceId, instanceId != ""}
	bastion.GroupId = sql.NullString{groupId, groupId != ""}
	return bastion
}

func (bastion *Bastion) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(bastion.PasswordHash), []byte(password))
}
