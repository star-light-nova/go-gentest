package cmd

const (
	ROOT_HELP_SHORT = `Generates '_test.go' files for '.go' files.`
	ROOT_HELP_LONG  = `Welcome to the '_test.go' generator.

The name literary means 'Go Generate Test'.
Right now, the only available command is 'start'

Please enter this command to know more about:
    > gentest start --help
`
	START_HELP_SHORT = `Command that allows to generate '_test.go' files in the current directory.`
	START_HELP_LONG  = `Actual implementation of the generator.

    Will create '<file_name>_test.go' file right after each '<file_name>.go' file.
    > gentest start

If you want to know more about the functionality, please read "README.md#flags".

Command example:
    Fully prints out program output to the current terminal session,
    so that it won't affect to your current directory.
    Consider this as a safe-run with no effect (*can be run in combination with other flags), but with informative output.
    > gentest start --dry-run

    If you want to generate everything in separate folder
    > gentest start --test-folder=test_folder

    If you have only one '.go' file to generate
    > gentest start --test-only=path/to/foo.go
`
)
