package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadWinnerInputMsg = "invalid winner input, expect format of 'Player wins'"

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayersInput := c.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}

	c.game.Start(numberOfPlayers, c.out)

	winnerInput := c.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(c.out, BadWinnerInputMsg)
		return
	}

	c.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
