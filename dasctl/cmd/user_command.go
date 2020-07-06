package cmd

import (
	"fmt"
	"strings"

	"github.com/bgentry/speakeasy"
	"github.com/spf13/cobra"
)

var (
	userShowDetail bool
)

// NewUserCommand returns the cobra command for "user".
func NewUserCommand() *cobra.Command {
	ac := &cobra.Command{
		Use:   "user <subcommand>",
		Short: "User related commands",
	}

	ac.AddCommand(newUserAddCommand())
	ac.AddCommand(newUserDeleteCommand())
	ac.AddCommand(newUserGetCommand())
	ac.AddCommand(newUserListCommand())
	ac.AddCommand(newUserChangePasswordCommand())
	ac.AddCommand(newUserGrantServCommand())
	ac.AddCommand(newUserRevokeServCommand())

	return ac
}

var (
	passwordInteractive bool
	passwordFromFlag    string
	noPassword          bool
	eps                 []string
	er                  error
)

func newUserAddCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "add <user name or user:password> [options]",
		Short: "Adds a new user",
		Run:   userAddCommandFunc,
	}

	cmd.Flags().BoolVar(&passwordInteractive, "interactive", true, "Read password from stdin instead of interactive terminal")
	cmd.Flags().StringVar(&passwordFromFlag, "new-user-password", "", "Supply password from the command line flag")
	cmd.Flags().BoolVar(&noPassword, "no-password", false, "Create a user without password (CN based auth only)")

	//eps, er = cmd.Flags().GetStringSlice("fd")

	eps = *cmd.Flags().StringSliceP("fd", "r", nil, "test arr")
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(eps)
	fmt.Println("==========")

	return &cmd
}

func newUserDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <user name>",
		Short: "Deletes a user",
		Run:   userDeleteCommandFunc,
	}
}

func newUserGetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get <user name> [options]",
		Short: "Gets detailed information of a user",
		Run:   userGetCommandFunc,
	}

	cmd.Flags().BoolVar(&userShowDetail, "detail", false, "Show permissions of roles granted to the user")

	return &cmd
}

func newUserListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists all users",
		Run:   userListCommandFunc,
	}
}

func newUserChangePasswordCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "passwd <user name> [options]",
		Short: "Changes password of user",
		Run:   userChangePasswordCommandFunc,
	}

	cmd.Flags().BoolVar(&passwordInteractive, "interactive", true, "If true, read password from stdin instead of interactive terminal")

	return &cmd
}

func newUserGrantServCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "grant-serv <user name> <service name>",
		Short: "Grants a role to a user",
		Run:   userGrantServCommandFunc,
	}
}

func newUserRevokeServCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "revoke-serv <user name> <service name>",
		Short: "Revokes a role from a user",
		Run:   userRevokeServCommandFunc,
	}
}

// userAddCommandFunc executes the "user add" command.
func userAddCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("user add command requires user name as its argument"))
	}

	var password string
	var user string

	if !noPassword {
		if passwordFromFlag != "" {
			user = args[0]
			password = passwordFromFlag
		} else {
			splitted := strings.SplitN(args[0], ":", 2)
			if len(splitted) < 2 {
				user = args[0]
				if !passwordInteractive {
					fmt.Scanf("%s", &password)
				} else {
					password = readPasswordInteractive(args[0])
				}
			} else {
				user = splitted[0]
				password = splitted[1]
				if len(user) == 0 {
					ExitWithError(ExitBadArgs, fmt.Errorf("empty user name is not allowed"))
				}
			}
		}
	} else {
		user = args[0]
	}

	fmt.Printf("User %s created\n", user)
}

// userDeleteCommandFunc executes the "user delete" command.
func userDeleteCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("user delete command requires user name as its argument"))
	}

	fmt.Println("删除用户成功", args[0])
}

// userGetCommandFunc executes the "user get" command.
func userGetCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("user get command requires user name as its argument"))
	}

	name := args[0]

	if userShowDetail {
		fmt.Printf("User: %s\n", name)
	} else {
		fmt.Println("用户信息", name)
	}
}

// userListCommandFunc executes the "user list" command.
func userListCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("user list command requires no arguments"))
	}

	fmt.Println("用户列表")
}

// userChangePasswordCommandFunc executes the "user passwd" command.
func userChangePasswordCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("user passwd command requires user name as its argument"))
	}

	var password string

	if !passwordInteractive {
		fmt.Scanf("%s", &password)
	} else {
		password = readPasswordInteractive(args[0])
	}

	fmt.Println("修改密码成功", args[0], password)
}

// userGrantServCommandFunc executes the "user grant-serv" command.
func userGrantServCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		ExitWithError(ExitBadArgs, fmt.Errorf("user grant-serv command requires user name and service name as its argument"))
	}

	fmt.Println(args[0], args[1])
}

// userRevokeServCommandFunc executes the "user revoke-serv" command.
func userRevokeServCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		ExitWithError(ExitBadArgs, fmt.Errorf("user revoke-serv requires user name and service name as its argument"))
	}

	fmt.Println(args[0], args[1])
}

func readPasswordInteractive(name string) string {
	prompt1 := fmt.Sprintf("Password of %s: ", name)
	password1, err1 := speakeasy.Ask(prompt1)
	if err1 != nil {
		ExitWithError(ExitBadArgs, fmt.Errorf("failed to ask password: %s", err1))
	}

	if len(password1) == 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("empty password"))
	}

	prompt2 := fmt.Sprintf("Type password of %s again for confirmation: ", name)
	password2, err2 := speakeasy.Ask(prompt2)
	if err2 != nil {
		ExitWithError(ExitBadArgs, fmt.Errorf("failed to ask password: %s", err2))
	}

	if password1 != password2 {
		ExitWithError(ExitBadArgs, fmt.Errorf("given passwords are different"))
	}

	return password1
}
