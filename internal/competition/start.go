package competition

import (
	"fmt"
	"time"
)

func (c *Competition) StartCompetition() error {
	for {
		printMenu()
		var command string
		fmt.Scan(&command)
		switch command {
		case update:
			c.log.Info(update)
			if err := c.update(); err != nil {
				return err
			}

		case test:
			if err := c.test(); err != nil {
				return err
			}

		case learn:
			if err := c.learn(); err != nil {
				return err
			}

		case restore:
			if err := c.restore(); err != nil {
				return err
			}

		case updateFromBackUp:
			if err := c.updateFromBackUp(); err != nil {
				return err
			}

		case backup:
			if err := c.backup(); err != nil {
				return err
			}

		case translate:
			if err := c.translator(); err != nil {
				c.log.Error(err)
			}

		case exit:
			fmt.Println("You have to do it, your dream wait")
			return nil
		}
	}
}

func printMenu() {
	menu := []string{
		fmt.Sprintf("      Update library (newWords.txt): [%v]\n", update),
		fmt.Sprintf("      Test knowlige:   [%v]\n", test),
		fmt.Sprintf("      Learn words:     [%v]\n", learn),
		fmt.Sprintf("      Restore by json: [%v]\n", restore),
		fmt.Sprintf("      Update by json:  [%v]\n", updateFromBackUp),
		fmt.Sprintf("      Backup:          [%v]\n", backup),
		fmt.Sprintf("      Init map words:  [%v]\n", translate),
		fmt.Sprintf("          Exit:        [%v]\n", exit),
	}

	for _, pos := range menu {
		fmt.Println(pos)
		time.Sleep(20 * time.Millisecond)
	}
}
