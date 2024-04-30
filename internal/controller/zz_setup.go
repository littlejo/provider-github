// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	branch "github.com/littlejo/provider-github/internal/controller/branch/branch"
	providerconfig "github.com/littlejo/provider-github/internal/controller/providerconfig"
	repository "github.com/littlejo/provider-github/internal/controller/repository/repository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		branch.Setup,
		providerconfig.Setup,
		repository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
