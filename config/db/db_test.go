package db

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

const filePath = "../../application.yaml"

func TestInitDb(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("load configre file failed: %s \n", err))
	}
	InitDbConfig()
	assert.NotNil(t, DB)
}
