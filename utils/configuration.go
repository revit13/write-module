// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package moduleutils

import (
	"os"

	"github.com/onsi/ginkgo"
	"github.com/rs/zerolog"
)

// Attributes that are defined in a config map or the runtime environment
const (
	AppUUIDKey          string = "APP_UUID"
	LoggingVerbosityKey string = "LOGGING_VERBOSITY"
	PrettyLoggingKey    string = "PRETTY_LOGGING"
)

// AppUUID returns the application UUID which is used
// for logging
func GetAppUUID() string {
	return os.Getenv(AppUUIDKey)
}

func SetIfNotSet(key string, value string, t ginkgo.GinkgoTInterface) {
	if _, b := os.LookupEnv(key); !b {
		if err := os.Setenv(key, value); err != nil {
			t.Fatalf("Could not set environment variable %s", key)
		}
	}
}

func DefaultTestConfiguration(t ginkgo.GinkgoTInterface) {
	SetIfNotSet(AppUUIDKey, "12345678", t)
	SetIfNotSet(LoggingVerbosityKey, "-1", t)
	SetIfNotSet(PrettyLoggingKey, "true", t)
}

func logEnvVariable(log zerolog.Logger, key string) {
	value, found := os.LookupEnv(key)
	if found {
		log.Info().Msg("Env variable " + key + " set to \"" + value + "\"")
	} else {
		log.Info().Msg("Env variable " + key + " not set")
	}
}

func LogEnvVariables(log zerolog.Logger) {
	envVarArray := [...]string{AppUUIDKey, LoggingVerbosityKey, PrettyLoggingKey}

	log.Info().Msg("Manager configured with the following environment variables:")
	for _, envVar := range envVarArray {
		logEnvVariable(log, envVar)
	}
}
