// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"fybrik.io/fybrik/manager/controllers/utils"
	"fybrik.io/fybrik/pkg/logging"
)

func main() {

	//Log    zerolog.Logger
	logger := logging.LogInit(logging.MODULE, "WriteModule")
	uuid = utils.FybrikAppUUID
	logger.Info().Str("clusterGroup", clusterGroup).Str("orgID", me.OrgId).Msg("Initializing Razee local")
	log := r.Log.With().Str(logging.CONTROLLER, "Blueprint").Str(utils.FybrikAppUUID, uuid).Str("blueprint", req.NamespacedName.String()).Logger()
	fmt.Println("hello world")
}
