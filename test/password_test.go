package test

import (
	"file_flow/common/helper"
	"testing"
)

func TestGenPassword(t *testing.T) {
	password, _ := helper.GenPassword("Aa123456")
	t.Log(password)
}
