package app

import (
	"github.com/spf13/cobra"
	"github.com/scmn-dev/secman/pkg/pipe/edit"
)

func EditCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Update or change a value in a password.",
		Long: "Update or change a value in a password.",
		Aliases: []string{"modify", "change"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				PwOpts.Password = args[0]
			}

			edit.Edit(&PwOpts)

			return nil
		},
	}

	cmd.Flags().BoolVarP(&PwOpts.Logins, "logins", "l", false, "Edit password from logins type.")
	cmd.Flags().BoolVarP(&PwOpts.CreditCards, "credit-cards", "c", false, "Edit password from credit cards type.")
	cmd.Flags().BoolVarP(&PwOpts.Emails, "emails", "e", false, "Edit password from emails type.")
	cmd.Flags().BoolVarP(&PwOpts.Notes, "notes", "n", false, "Edit password from notes type.")
	cmd.Flags().BoolVarP(&PwOpts.Servers, "servers", "s", false, "Edit password from servers type.")

	return cmd
}
