package main

import (
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/cli/generate"
	"github.com/sev-2/raiden/pkg/cli/apply"
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
					apply.ApplyLogger.Error(err.Error())
					return
				}
				f.ProjectPath = curDir
			}

			config, err := raiden.LoadConfig(nil)
			if err != nil {
				apply.ApplyLogger.Error(err.Error())
				return
			}

			if err := generate.Run(&f.Generate, config, f.ProjectPath, false); err != nil {
				apply.ApplyLogger.Error(err.Error())
				return
			}

			// register app resource
			bootstrap.RegisterRpc()
			bootstrap.RegisterRoles()
			bootstrap.RegisterModels()
			bootstrap.RegisterStorages()
			
			if err = resource.Apply(&f, config); err != nil {
				apply.ApplyLogger.Error(err.Error())
			}
		},
	}

	f.BindLog(cmd)
	cmd.Flags().StringVarP(&f.ProjectPath, "project-path", "p", "", "set project path")
	cmd.Flags().BoolVarP(&f.RpcOnly, "rpc-only", "", false, "apply rpc only")
	cmd.Flags().BoolVarP(&f.RolesOnly, "roles-only", "r", false, "apply roles only")
	cmd.Flags().BoolVarP(&f.ModelsOnly, "models-only", "m", false, "apply models only")
	cmd.Flags().BoolVarP(&f.StoragesOnly, "storages-only", "", false, "apply storages only")
	cmd.Flags().StringVarP(&f.AllowedSchema, "schema", "s", "", "set allowed schema to apply, use coma separator for multiple schema")
	cmd.Flags().BoolVar(&f.DryRun, "dry-run", false, "run apply in simulate mode without actual running apply change")

	f.Generate.Bind(cmd)

	cmd.Execute()
}
