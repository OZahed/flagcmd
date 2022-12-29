package docs

// import (
// 	"flag"
// 	"fmt"
// 	"time"

// 	"github.com/OZahed/flagcmd"
// )

// var (
// 	fs = flag.NewFlagSet("add", flag.ContinueOnError)

// 	sb = &flagcmd.SubCommand{
// 		SubCommandName: "add",
// 		SortDesc:       `adds new Tood list`,
// 		LongDesc:       `you can have new Todo Lists `,
// 		Usage: `add -t "title" -du 2 -p high -c Category
// OPT list:
// 	-t: title
// 	-du: due time
// 	-p: priority
// 	-c: category
// `,
// 		FlagSet: fs,
// 		Handler: func(sc *flagcmd.SubCommand) error {
// 			fmt.Println("Called add sub command ")
// 			return nil
// 		},
// 		ErrorHandler: func(sce flagcmd.SubCommandError) {
// 			fmt.Println(sce.Error())
// 		},
// 	}
// 	sbList = &flagcmd.SubCommand{
// 		SubCommandName: "list",
// 		SortDesc:       `lists all Todo stack`,
// 		LongDesc:       ``,
// 		Usage: `list n
// OPT list:
// 	n: list of the tasks for n days from now (if n is negative number, it will show open todos with passed
// 	due time)
// `,
// 		FlagSet: fs,
// 		Handler: func(sc *flagcmd.SubCommand) error {
// 			fmt.Println("Called add sub command ")
// 			return nil
// 		},
// 		ErrorHandler: func(sce flagcmd.SubCommandError) {
// 			fmt.Println(sce.Error())
// 		},
// 	}
// )

// func init() {

// 	if err := flagcmd.RegisterSubCommand(sb); err != nil {
// 		fmt.Println(err)
// 	}

// 	if err := flagcmd.RegisterSubCommand(sbList); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func main() {
// 	flagcmd.SetAppName("Applications")

// 	title := fs.String("t", "", "title")
// 	due := fs.Duration("du", 1, "due time duration")
// 	priority := fs.String("p", "low", "priority")
// 	category := fs.String("c", "NA", "category")

// 	if err := flagcmd.Parse(); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	dueTime := time.Now().Add(*due * 24 * time.Hour)
// 	fmt.Printf("Values:\ntitle: %s\npriority: %s\nCategory: %s\nDueTime: %s\n",
// 		*title,
// 		*priority,
// 		*category,
// 		dueTime.Format(time.RFC3339),
// 	)
// }
