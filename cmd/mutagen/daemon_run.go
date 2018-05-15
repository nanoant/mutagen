package main

import (
	"os"
	"os/signal"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/havoc-io/mutagen/cmd"
	"github.com/havoc-io/mutagen/pkg/agent"
	"github.com/havoc-io/mutagen/pkg/daemon"
	"github.com/havoc-io/mutagen/pkg/rpc"
	"github.com/havoc-io/mutagen/pkg/session"
	"github.com/havoc-io/mutagen/pkg/ssh"
)

func daemonRunMain(command *cobra.Command, arguments []string) {
	// Validate arguments.
	if len(arguments) != 0 {
		cmd.Fatal(errors.New("unexpected arguments provided"))
	}

	// TODO: Do we eventually want to encapsulate the construction of the daemon
	// RPC server into the daemon package, much like we do with endpoints? It
	// becomes a bit difficult to do cleanly. Also, I want the ability to have
	// different processes host the daemon (e.g. a GUI). In those cases, we may
	// want to add additional services that wouldn't be present in the CLI
	// daemon. So I'll leave things the way they are for now, but I'd like to
	// keep thinking about this for the future. One easy thing we could do is
	// move the daemon lock into the daemon service (and add a corresponding
	// shutdown method to the daemon service).

	// Attempt to acquire the daemon lock and defer its release. If there is a
	// crash, the lock will be released by the OS automatically, but on Windows
	// this may only happen after some unspecified period of time (though it
	// does seem to be basically instant).
	lock, err := daemon.AcquireLock()
	if err != nil {
		cmd.Fatal(errors.Wrap(err, "unable to acquire daemon lock"))
	}
	defer lock.Unlock()

	// Perform housekeeping.
	agent.Housekeep()
	session.HousekeepCaches()
	session.HousekeepStaging()

	// Create the RPC server.
	server := rpc.NewServer()

	// Create and register the daemon service.
	daemonService, daemonTermination := daemon.NewService()
	server.Register(daemonService)

	// Create and regsiter the SSH service.
	sshService := ssh.NewService()
	server.Register(sshService)

	// Create the and register the session service and defer its shutdown. We
	// want to do a clean shutdown because we don't want to information
	// generated during a synchronization cycle.
	sessionService, err := session.NewService(sshService)
	if err != nil {
		cmd.Fatal(errors.Wrap(err, "unable to create session service"))
	}
	server.Register(sessionService)
	defer sessionService.Shutdown()

	// Create the daemon listener and defer its closure.
	listener, err := daemon.NewListener()
	if err != nil {
		cmd.Fatal(errors.Wrap(err, "unable to create daemon listener"))
	}
	defer listener.Close()

	// Serve incoming connections in a separate Goroutine, watching for serving
	// failure.
	serverTermination := make(chan error, 1)
	go func() {
		serverTermination <- server.Serve(listener)
	}()

	// Wait for termination from a signal, the server, or the daemon server. We
	// treat daemon termination as a non-error.
	signalTermination := make(chan os.Signal, 1)
	signal.Notify(signalTermination, cmd.TerminationSignals...)
	select {
	case sig := <-signalTermination:
		cmd.Fatal(errors.Errorf("terminated by signal: %s", sig))
	case <-daemonTermination:
		return
	case err = <-serverTermination:
		cmd.Fatal(errors.Wrap(err, "premature server termination"))
	}
}

var daemonRunCommand = &cobra.Command{
	Use:    "run",
	Short:  "Runs the Mutagen daemon",
	Run:    daemonRunMain,
	Hidden: true,
}

var daemonRunConfiguration struct {
	help bool
}

func init() {
	// Bind flags to configuration. We manually add help to override the default
	// message, but Cobra still implements it automatically.
	flags := daemonRunCommand.Flags()
	flags.BoolVarP(&daemonRunConfiguration.help, "help", "h", false, "Show help information")
}