package main

import (
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/cli/generate"
	"github.com/sev-2/raiden/pkg/cli/imports"
	"github.com/sev-2/raiden/pkg/resource"
	"github.com/sev-2/raiden/pkg/utils"
	"github.com/spf13/cobra"
	"mediversepreip/internal/bootstrap"
)


func main() {
	f := resource.Flags{}

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			f.CheckAndActivateDebug(cmd)
			// load configuration
			if f.ProjectPath == "" {
				curDir, err := utils.GetCurrentDirectory()
				if err != nil {
					imports.ImportLogger.Error(err.Error())
					return
				}
				f.ProjectPath = curDir
			}

			config, err := raiden.LoadConfig(nil)
			if err != nil {
				imports.ImportLogger.Error(err.Error())
				return
			}

			// register app resource
			bootstrap.RegisterRpc()
			bootstrap.RegisterRoles()
			bootstrap.RegisterModels()
			bootstrap.RegisterStorages()

			if err = generate.Run(&f.Generate, config, f.ProjectPath, false); err != nil {
				imports.ImportLogger.Error(err.Error())
			}

			if err := resource.Import(&f, config); err != nil {
				imports.ImportLogger.Error(err.Error())
			}

			if !f.DryRun {
				imports.ImportLogger.Info("regenerate bootstrap file")
				if err = generate.Run(&f.Generate, config, f.ProjectPath, false); err != nil {
					imports.ImportLogger.Error(err.Error())
					return
				}
				imports.ImportLogger.Info("finish import process")
			}
		},
	}

	f.BindLog(cmd)
	cmd.Flags().StringVarP(&f.ProjectPath, "project-path", "p", "", "set project path")
	cmd.Flags().BoolVarP(&f.RpcOnly, "rpc-only", "", false, "import rpc only")
	cmd.Flags().BoolVarP(&f.RolesOnly, "roles-only", "r", false, "import roles only")
	cmd.Flags().BoolVarP(&f.ModelsOnly, "models-only", "m", false, "import models only")
	cmd.Flags().BoolVarP(&f.StoragesOnly, "storages-only", "", false, "import storages only")
	cmd.Flags().StringVarP(&f.AllowedSchema, "schema", "s", "", "set allowed schema to import, use coma separator for multiple schema")
	cmd.Flags().BoolVar(&f.DryRun, "dry-run", false, "run import in simulate mode without actual import resource as code")

	f.Generate.Bind(cmd)

	cmd.Execute()
}
