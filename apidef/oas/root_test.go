package oas

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/TykTechnologies/tyk/apidef"
)

func TestXTykAPIGateway(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var emptyXTykAPIGateway XTykAPIGateway

		var convertedAPI apidef.APIDefinition
		emptyXTykAPIGateway.ExtractTo(&convertedAPI)

		var resultXTykAPIGateway XTykAPIGateway
		resultXTykAPIGateway.Fill(convertedAPI)

		assert.Equal(t, emptyXTykAPIGateway, resultXTykAPIGateway)
	})

	t.Run("filled", func(t *testing.T) {
		initialAPI := apidef.APIDefinition{}
		Fill(t, &initialAPI, 0)

		xTykAPIGateway := XTykAPIGateway{}
		xTykAPIGateway.Fill(initialAPI)

		var convertedAPI apidef.APIDefinition
		xTykAPIGateway.ExtractTo(&convertedAPI)

		bytes, _ := json.MarshalIndent(convertedAPI, "", "  ")
		fmt.Println(string(bytes))

		assert.Equal(t, initialAPI, convertedAPI)
	})
}

func TestInfo(t *testing.T) {
	var emptyInfo Info

	var convertedAPI apidef.APIDefinition
	emptyInfo.ExtractTo(&convertedAPI)

	var resultInfo Info
	resultInfo.Fill(convertedAPI)

	assert.Equal(t, emptyInfo, resultInfo)
}

func TestState(t *testing.T) {
	var emptyState State

	var convertedAPI apidef.APIDefinition
	emptyState.ExtractTo(&convertedAPI)

	var resultState State
	resultState.Fill(convertedAPI)

	assert.Equal(t, emptyState, resultState)
}

func TestUpstream(t *testing.T) {
	var emptyUpstream Upstream

	var convertedAPI apidef.APIDefinition
	emptyUpstream.ExtractTo(&convertedAPI)

	var resultUpstream Upstream
	resultUpstream.Fill(convertedAPI)

	assert.Equal(t, emptyUpstream, resultUpstream)
}

func TestServer(t *testing.T) {
	var emptyServer Server

	var convertedAPI apidef.APIDefinition
	emptyServer.ExtractTo(&convertedAPI)

	var resultServer Server
	resultServer.Fill(convertedAPI)

	assert.Equal(t, emptyServer, resultServer)
}

func TestListenPath(t *testing.T) {
	var emptyListenPath ListenPath

	var convertedAPI apidef.APIDefinition
	emptyListenPath.ExtractTo(&convertedAPI)

	var resultListenPath ListenPath
	resultListenPath.Fill(convertedAPI)

	assert.Equal(t, emptyListenPath, resultListenPath)
}

// Fill populates the given input with non-default values. Index is where to start incrementing values.
func Fill(t *testing.T, input interface{}, index int) {
	v := reflect.ValueOf(input).Elem()

	switch kind := v.Type().Kind(); kind {
	case reflect.String:
		v.SetString(fmt.Sprintf("%d", index))
	case reflect.Bool:
		v.SetBool(true)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(index))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(index))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(float64(index))
	case reflect.Slice:
		if v.Type() == reflect.TypeOf(json.RawMessage{}) {
			v.Set(reflect.ValueOf(json.RawMessage(`{"test-key":"test-value"}`)))
		} else {
			newSlice := reflect.MakeSlice(v.Type(), 3, 3)

			for i := 0; i < 3; i++ {
				newValue := reflect.New(v.Type().Elem()).Elem()
				Fill(t, newValue.Addr().Interface(), index+i)
				newSlice.Index(i).Set(newValue)
			}

			v.Set(newSlice)
		}
	case reflect.Interface:
		v.Set(reflect.ValueOf(1))
	case reflect.Map:

		if v.Type() == reflect.TypeOf(map[string]apidef.AuthConfig{}) {
			v.Set(reflect.ValueOf(FillTestAuthConfigs(t, index)))
		} else {

			newMap := reflect.MakeMapWithSize(v.Type(), 0)
			for i := 0; i < 3; i++ {
				newKey := reflect.New(v.Type().Key()).Elem()
				Fill(t, newKey.Addr().Interface(), index+i)

				newValue := reflect.New(v.Type().Elem()).Elem()
				Fill(t, newValue.Addr().Interface(), index+i)

				newMap.SetMapIndex(newKey, newValue)
			}

			v.Set(newMap)
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if v.Type().Field(i).Tag.Get("json") == "-" || v.Type().Field(i).Tag.Get("json") == "" {
				continue
			}

			Fill(t, fv.Addr().Interface(), index+i+1)
		}

	case reflect.Ptr:
		newValue := reflect.New(v.Type().Elem()).Elem()
		Fill(t, newValue.Addr().Interface(), index)
		v.Set(newValue.Addr())
	default:
		t.Fatalf("uncovered kind in API definition: %s", kind.String())
	}
}

func FillTestAuthConfigs(t *testing.T, index int) map[string]apidef.AuthConfig {
	authConfigs := make(map[string]apidef.AuthConfig)

	a := apidef.AuthConfig{}
	Fill(t, &a, index)
	authConfigs["authToken"] = a

	return authConfigs
}
