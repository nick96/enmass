package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nick96/enmass/witheditor"
)

// GroupActionType representats of the type of action that can be done
type GroupActionType int

// Actions that can be performed on groups
const (
	NOGROUPACTIONTYPE GroupActionType = iota
	GET
	SEND
	CHECK
	STAT
)

// GroupAction representation of the actions that can be done.
// Combined with the GroupActionType, this tells us what we're doing.
type GroupAction int

// Actions performed
const (
	NOGROUPACTION GroupAction = iota
	EMAIL
	TEXT
)

// CmdLineArgs represents the command line arguments that can be
// given.
type CmdLineArgs struct {
	CredentialsPath string
	TokenPath       string
	GroupName       string
	ActionType      GroupActionType
	Action          GroupAction
}

// Parse command line arguments
func parseArgs(args []string) CmdLineArgs {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Group and actions not specified")
		Usage()
		os.Exit(1)
	} else if len(args) == 1 {
		fmt.Fprintln(os.Stderr, "Action on group must be specified")
		Usage()
		os.Exit(1)
	}

	var actionType GroupActionType
	var action GroupAction
	var credentialsPath string
	var tokenPath string
	var groupName string

	groupName = args[0]

	switch strings.ToLower(args[1]) {
	case "check":
		actionType = CHECK
	case "stat":
		actionType = STAT
	case "get":
		actionType = GET
	case "send":
		actionType = SEND
	default:
		fmt.Fprintf(os.Stderr, "%s is not a known argument\n", args[1])
		Usage()
		os.Exit(1)
	}

	if (actionType == CHECK || actionType == STAT) && len(args) > 2 {
		fmt.Fprintln(os.Stderr, "Ignoring trailing junk after argument",
			strings.Join(args[2:], " "))
	} else if actionType == GET || actionType == SEND {
		if len(args) > 2 {
			switch strings.ToLower(args[2]) {
			case "email":
				action = EMAIL
			case "text":
				action = TEXT
			default:
				fmt.Fprintf(os.Stderr, "%s not not a know argument for get or send\n",
					args[2])
			}
		} else {
			fmt.Fprintln(os.Stderr, "get and send actions required additional arguments")
			Usage()
			os.Exit(1)
		}
	}

	if len(args) > 3 {
		credentialsSet := false
		tokenSet := false
		for _, arg := range args[3:] {
			if strings.HasPrefix(arg, "--credentials") {
				if !credentialsSet {
					credentialsSet = true
					credentialsPath = strings.TrimPrefix("--credentials", arg)
				} else {
					fmt.Fprintf(os.Stderr, "Credentials path is already set")
					Usage()
					os.Exit(1)
				}
			} else if strings.HasPrefix(arg, "-c") {
				if !credentialsSet {
					credentialsSet = true
					credentialsPath = strings.TrimPrefix("-c", arg)
				} else {
					fmt.Fprintf(os.Stderr, "Credentials path is already set")
					Usage()
					os.Exit(1)
				}
			} else if strings.HasPrefix(arg, "--Token") {
				if !tokenSet {
					tokenSet = true
					credentialsPath = strings.TrimPrefix("--Token", arg)
				} else {
					fmt.Fprintf(os.Stderr, "Token path is already set")
					Usage()
					os.Exit(1)
				}
			} else if strings.HasPrefix(arg, "-t") {
				if !tokenSet {
					tokenSet = true
					credentialsPath = strings.TrimPrefix("-t", arg)
				} else {
					fmt.Fprintf(os.Stderr, "Token path is already set")
					Usage()
					os.Exit(1)
				}
			} else {
				fmt.Fprintf(os.Stderr, "%s is not a known option\n", arg)
				Usage()
				os.Exit(1)
			}
		}
	} else {
		credentialsPath = "credentials.json"
		tokenPath = "token.json"
	}

	return CmdLineArgs{
		GroupName:       groupName,
		ActionType:      actionType,
		Action:          action,
		CredentialsPath: credentialsPath,
		TokenPath:       tokenPath,
	}
}

// Usage prints out tool usage
func Usage() {
	fmt.Fprint(os.Stderr, `Usage: enmass <group> [arguments] [options]
  group is the name of the contact group we will be operating on.

  Arguments:
    check: Check that all of the members of <group> have at least one
           email and phone number
    stat: Print out the stats of a group (nuber of members, emails and
          phone numbers, etc.)
    get: Perform a get operation on the group (this requires a second
         argument of "email" or "text")
    send: Perform a send operation on the group (this requires a
          second argument of "email" or "text")

  Options:
    --credentials/-c: Path to the credentials file
    --Token/-t:       Path to the Token file

  The usage of the "get" and "send" arguments requires an additional
  argument of "email" or "text". This is used to specify what
  operation is to be performed. "email" indicates the operation is to
  be on emails and "text" on phone numbers. The possible combinations
  are:

    get email:  Get all the emails for the group
    get text:   Get all the phone numbers for the group
    send email: Send an email to all the members of the group
    send text:  Send a text message to all the members of the group
`)
}

// Read a message from standard input until the EOF character is read.
func readMessage() (string, error) {
	return witheditor.WithEditor(
		"vim",
		os.Getenv("ENMASS_SIGNATURE"),
		"Enter the message you wish to send to the group.")

}
