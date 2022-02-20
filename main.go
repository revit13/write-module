// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"fybrik.io/fybrik/manager/controllers/utils"
	"fybrik.io/fybrik/pkg/logging"
	moduleutils "fybrik.io/write-module/utils"
)

func main() {

	logger := logging.LogInit(logging.MODULE, "WriteModule")
	uuid := moduleutils.GetAppUUID()

	logger.Info().Msg("Module")
	log := logger.With().Str(utils.FybrikAppUUID, uuid).Logger()
	log.Info().Msg("hello")

	fmt.Println("hello world")
}
