package main
import (
	"fmt"
	"os"
	"database/sql"
	"github.com/tmcausla/gatoraggregator/internal/config"
	"github.com/tmcausla/gatoraggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	/*fmt.Println("old config:") // USE FOR DEBUGGING
	homeDir, err := os.UserHomeDir() // //////////////////////
	if err != nil { // //////////////////////
		fmt.Println("error getting home directory: %w", err) // //////////////////////
		return // //////////////////////
	} // //////////////////////
	data, err := os.ReadFile(homeDir + "/.gatorconfig.json") // //////////////////////
	if err != nil { // //////////////////////
		fmt.Println("error reading file") // //////////////////////
		return // //////////////////////
	} // //////////////////////
	fmt.Println(string(data))*/ // USE FOR DEBUGGING

	mainConfig, err := config.Read()
	if err != nil {
		fmt.Println("error creating struct")
		return
	}
	mainState := state{
		cfg: &mainConfig,
	}
	db, err := sql.Open("postgres", mainState.cfg.DbURL)
	defer db.Close()

	if err != nil {
		fmt.Printf("error opening database from DbURL: %v", err)
		os.Exit(1)
	}
	mainState.db = database.New(db)

	cmdList := commands{
		cmds: map[string]func(*state, command) error{},
	}

	cmdList.register("login", handlerLogin)
	cmdList.register("register", handlerRegister)
	cmdList.register("reset", handlerReset)
	cmdList.register("users", handlerUsers)

	userArgs := os.Args
	if len(userArgs) < 2 {
		fmt.Println("no commands provided")
		os.Exit(1)
	}

	cmdName := userArgs[1]
	cmdArgs := userArgs[2:]
	userCmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = cmdList.run(&mainState, userCmd)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}



	/*fmt.Println("new config:") // USE FOR DEBUGGING
	data, err = os.ReadFile(homeDir + "/.gatorconfig.json") // //////////////////////
	if err != nil { // //////////////////////
		fmt.Println("error reading file") // //////////////////////
		return // //////////////////////
	} // //////////////////////
	fmt.Println(string(data))*/ // USE FOR DEBUGGING

	os.Exit(0)
}