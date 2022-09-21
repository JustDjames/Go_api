package main

import (
	"os"
	"testing"
)

func Test_ArgSet(t *testing.T) {
	arg := "Argument_value"
	expected := "Argument_value"
	result := var_parser(arg, "", "")

	if result != expected {
		t.Errorf("Test 'Test_ArgSet' failed, expected -> %v, got -> %v ", expected, result)
	} else {
		t.Logf("Test 'Test_ArgSet' Passed, expected -> %v, got -> %v ", expected, result)
	}
}

func Test_EnvSet(t *testing.T) {
	env := "EnvironmentVar_value"
	expected := "EnvironmentVar_value"

	t.Log("setting Env Var")

	err := os.Setenv("ENV_VAR", env)

	if err != nil {
		t.Errorf("Error setting ENV_VAR, Error: %s", err)
	}

	result := var_parser("", "ENV_VAR", "")

	if result != expected {
		t.Errorf("Test 'Test_EnvSet' failed, expected -> %v, got -> %v ", expected, result)
	} else {
		t.Logf("Test 'Test_EnvSet' Passed, expected -> %v, got -> %v ", expected, result)
	}

	t.Log("unsetting Env Var")

	err = os.Unsetenv("ENV_VAR")

	if err != nil {
		t.Errorf("Error unsetting ENV_VAR, Error: %s", err)
	}
}

func Test_DefaultSet(t *testing.T) {
	def := "Default_value"
	expected := "Default_value"

	result := var_parser("", "", def)

	if result != expected {
		t.Errorf("Test 'Test_DefaultSet' failed, expected -> %v, got -> %v ", expected, result)
	} else {
		t.Logf("Test 'Test_DefaultSet' Passed, expected -> %v, got -> %v ", expected, result)
	}
}

func Test_EnvOverDef(t *testing.T) {
	env := "EnvironmentVar_value"
	def := "Default_value"
	expected := "EnvironmentVar_value"

	t.Log("setting Env Var")

	err := os.Setenv("ENV_VAR", env)

	if err != nil {
		t.Errorf("Error setting ENV_VAR, Error: %s", err)
	}

	result := var_parser("", "ENV_VAR", def)

	if result != expected {
		t.Errorf("Test 'Test_EnvOverDef' failed, expected -> %v, got -> %v ", expected, result)
	} else {
		t.Logf("Test 'Test_EnvOverDef' Passed, expected -> %v, got -> %v ", expected, result)
	}
}

func Test_ArgOverAll(t *testing.T) {
	arg := "Argument_value"
	env := "EnvironmentVar_value"
	def := "Default_value"
	expected := "Argument_value"

	t.Log("setting Env Var")

	err := os.Setenv("ENV_VAR", env)

	if err != nil {
		t.Errorf("Error setting ENV_VAR, Error: %s", err)
	}

	result := var_parser(arg, "ENV_VAR", def)

	if result != expected {
		t.Errorf("Test 'Test_ArgOverAll' failed, expected -> %v, got -> %v ", expected, result)
	} else {
		t.Logf("Test 'Test_ArgOverAll' Passed, expected -> %v, got -> %v ", expected, result)
	}

	err = os.Unsetenv("ENV_VAR")

	if err != nil {
		t.Errorf("Error unsetting ENV_VAR, Error: %s", err)
	}
}
