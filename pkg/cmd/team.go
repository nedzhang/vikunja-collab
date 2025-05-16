// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-present Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public Licensee as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public Licensee for more details.
//
// You should have received a copy of the GNU Affero General Public Licensee
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"

	"code.vikunja.io/api/pkg/db"
	"code.vikunja.io/api/pkg/initialize"
	"code.vikunja.io/api/pkg/log"
	"code.vikunja.io/api/pkg/models"
	"github.com/spf13/cobra"
	"xorm.io/xorm"
)

// var (
// 	userFlagUsername              string
// 	userFlagEmail                 string
// 	userFlagPassword              string
// 	userFlagAvatar                = "default"
// 	userFlagResetPasswordDirectly bool
// 	userFlagEnableUser            bool
// 	userFlagDisableUser           bool
// 	userFlagDeleteNow             bool
// 	userFlagDeleteConfirm         bool
// )

func init() {

	// // User create flags
	// userCreateCmd.Flags().StringVarP(&userFlagUsername, "username", "u", "", "The username of the new user.")
	// _ = userCreateCmd.MarkFlagRequired("username")
	// userCreateCmd.Flags().StringVarP(&userFlagEmail, "email", "e", "", "The email address of the new user.")
	// _ = userCreateCmd.MarkFlagRequired("email")
	// userCreateCmd.Flags().StringVarP(&userFlagPassword, "password", "p", "", "The password of the new user. You will be asked to enter it if not provided through the flag.")
	// userCreateCmd.Flags().StringVarP(&userFlagAvatar, "avatar-provider", "a", "", "The avatar provider of the new user. Optional.")

	// // Team update flags
	// userUpdateCmd.Flags().StringVarP(&userFlagUsername, "username", "u", "", "The new username of the user.")
	// userUpdateCmd.Flags().StringVarP(&userFlagEmail, "email", "e", "", "The new email address of the user.")
	// userUpdateCmd.Flags().StringVarP(&userFlagAvatar, "avatar-provider", "a", "", "The new avatar provider of the new user.")

	// // Reset PW flags
	// userResetPasswordCmd.Flags().BoolVarP(&userFlagResetPasswordDirectly, "direct", "d", false, "If provided, reset the password directly instead of sending the user a reset mail.")
	// userResetPasswordCmd.Flags().StringVarP(&userFlagPassword, "password", "p", "", "The new password of the user. Only used in combination with --direct. You will be asked to enter it if not provided through the flag.")

	// // Change status flags
	// userChangeStatusCmd.Flags().BoolVarP(&userFlagDisableUser, "disable", "d", false, "Disable the user.")
	// userChangeStatusCmd.Flags().BoolVarP(&userFlagEnableUser, "enable", "e", false, "Enable the user.")

	// // User deletion flags
	// userDeleteCmd.Flags().BoolVarP(&userFlagDeleteNow, "now", "n", false, "If provided, deletes the user immediately instead of sending them an email first.")

	// // Bypass confirm prompt
	// userDeleteCmd.Flags().BoolVarP(&userFlagDeleteConfirm, "confirm", "c", false, "Bypasses any prompts confirming the deletion request, use with caution!")

	teamCmd.AddCommand(teamListCmd)
	rootCmd.AddCommand(teamCmd)
}

var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Manage team locally through the cli.",
}

// ListAllTeams returns all teams
func ListAllTeams(s *xorm.Session) (teams []*models.Team, err error) {
	err = s.Find(&teams)
	return
}

var teamListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of all teams.",
	PreRun: func(_ *cobra.Command, _ []string) {
		initialize.FullInit()
	},

	Run: func(_ *cobra.Command, _ []string) {
		s := db.NewSession()
		defer s.Close()

		team := &models.Team{
			ID: 1,
		}

		team.ReadOne(s, nil)

		fmt.Println("team 1: ", team)

		t := models.Team{}

		exists, err := s.
			Where("1 = 1").
			Get(&t)
		if err != nil {
			log.Fatalf("Error getting teams: %s", err)
		}
		if !exists {
			log.Fatal("********* NO TEAM FOUND!!! **********", t)
		}

		fmt.Println("Teams loaded: ", t)

		// // teams, err := ListAllTeams(s)

		// // if err != nil {
		// // 	_ = s.Rollback()
		// // 	log.Fatalf("Error getting teams: %s", err)
		// // }

		// // if err := s.Commit(); err != nil {
		// // 	log.Fatalf("Error getting teams: %s", err)
		// // }

		// table := tablewriter.NewWriter(os.Stdout)
		// table.SetHeader([]string{
		// 	"ID",
		// 	"TeamName",
		// 	"Creator.ID",
		// 	"Creator",
		// 	"Issuer",
		// 	"Created",
		// 	"Updated",
		// })

		// for _, u := range teams {
		// 	table.Append([]string{
		// 		strconv.FormatInt(u.ID, 10),
		// 		u.Name,
		// 		strconv.FormatInt(u.CreatedByID, 10),
		// 		u.CreatedBy.Name,
		// 		u.Issuer,
		// 		u.Created.Format(time.RFC3339),
		// 		u.Updated.Format(time.RFC3339),
		// 	})
		// }

		// table.Render()
	},
}
