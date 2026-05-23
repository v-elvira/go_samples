package main

import (
	"errors"
	"fmt"
	"os"
)

// label - уникальное наименование
type label string

// command - команда, которую можно выполнять в игре
type command label

// список доступных команд
var (
	eat  = command("eat")
	take = command("take")
	talk = command("talk to")
)

// thing - объект, который существует в игре
type thing struct {
	name    label
	actions map[command]string
}

// supports() возвращает true, если объект
// поддерживает команду action
func (t thing) supports(action command) bool {
	_, ok := t.actions[action]
	return ok
}

// String() возвращает описание объекта
func (t thing) String() string {
	return string(t.name)
}

// полный список объектов в игре
var (
	apple = thing{"apple", map[command]string{
		eat:  "mmm, delicious!",
		take: "you have an apple now",
	}}
	bob = thing{"bob", map[command]string{
		talk: "Bob says hello",
	}}
	coin = thing{"coin", map[command]string{
		take: "you have a coin now",
	}}
	mirror = thing{"mirror", map[command]string{
		take: "you have a mirror now",
		talk: "mirror does not answer",
	}}
	mushroom = thing{"mushroom", map[command]string{
		eat:  "tastes funny",
		take: "you have a mushroom now",
	}}
)

// step описывает шаг игры: сочетание команды и объекта
type step struct {
	cmd command
	obj thing
}

// isValid() возвращает true, если объект
// совместим с командой
func (s step) isValid() bool {
	return s.obj.supports(s.cmd)
}

// String() возвращает описание шага
func (s step) String() string {
	return fmt.Sprintf("%s %s", s.cmd, s.obj)
}

// начало решения

// invalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type invalidStepError struct {
	// better to form error message in Error() method with fmt.Sprintf and mini info-fields stored in struct
	message string
	advice  string
}

// notEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type notEnoughObjectsError struct {
	message string
	advice  string
}

// commandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type commandLimitExceededError struct {
	message string
	advice  string
}

// objectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type objectLimitExceededError struct {
	message string
	advice  string
}

// gameOverError - ошибка, которая произошла в игре
type gameOverError struct {
	// количество шагов, успешно выполненных
	// до того, как произошла ошибка
	nSteps int
	err    error
}

func (e gameOverError) Error() string {
	return e.err.Error()
}

func (e gameOverError) Unwrap() error {
	return e.err
}

func (e invalidStepError) Error() string {
	// better to form error message in Error() method with fmt.Sprintf and mini info-fields stored in struct
	return e.message
}

func (e notEnoughObjectsError) Error() string {
	return e.message
}

func (e commandLimitExceededError) Error() string {
	return e.message
}

func (e objectLimitExceededError) Error() string {
	return e.message
}

// player - игрок
type player struct {
	// количество съеденного
	nEaten int
	// количество диалогов
	nDialogs int
	// инвентарь
	inventory []thing
}

// has() возвращает true, если у игрока
// в инвентаре есть предмет obj
func (p *player) has(obj thing) bool {
	for _, got := range p.inventory {
		if got.name == obj.name {
			return true
		}
	}
	return false
}

// do() выполняет команду cmd над объектом obj
// от имени игрока
func (p *player) do(cmd command, obj thing) error {
	// действуем в соответствии с командой
	switch cmd {
	case eat:
		if p.nEaten > 1 {
			return commandLimitExceededError{"you don't want to eat anymore", "eat less"}
		}
		p.nEaten++
	case take:
		if p.has(obj) {
			return objectLimitExceededError{fmt.Sprintf("you already have a %s", obj), fmt.Sprintf("don't be greedy, 1 %s is enough", obj)}
		}
		p.inventory = append(p.inventory, obj)
	case talk:
		if p.nDialogs > 0 {
			return commandLimitExceededError{"you don't want to talk anymore", "talk to less"}
		}
		p.nDialogs++
	}
	return nil
}

// newPlayer создает нового игрока
func newPlayer() *player {
	return &player{inventory: []thing{}}
}

// game описывает игру
type game struct {
	// игрок
	player *player
	// объекты игрового мира
	things map[label]int
	// количество успешно выполненных шагов
	nSteps int
}

// has() проверяет, остались ли в игровом мире указанные предметы
func (g *game) has(obj thing) bool {
	count := g.things[obj.name]
	return count > 0
}

// execute() выполняет шаг step
func (g *game) execute(st step) error {
	// not my:
	// defer func() { err = gameOverError{nSteps: g.nSteps, err: err}.Wrap() }() // +Wrap() in gameOverError:

	/*
		// func (e gameOverError) Wrap() error {
		// if e.err == nil {
		// 	return nil
		// }
		// return e
		// }
	*/

	// проверяем совместимость команды и объекта
	if !st.isValid() {
		return gameOverError{nSteps: g.nSteps, err: invalidStepError{message: fmt.Sprintf("cannot %s", st), advice: fmt.Sprintf("things like '%s %s' are impossible", st.cmd, st.obj)}}
	}

	// когда игрок берет или съедает предмет,
	// тот пропадает из игрового мира
	if st.cmd == take || st.cmd == eat {
		if !g.has(st.obj) {
			return gameOverError{nSteps: g.nSteps, err: notEnoughObjectsError{message: fmt.Sprintf("there are no %ss left", st.obj), advice: fmt.Sprintf("be careful with scarce %ss", st.obj)}}
		}
		g.things[st.obj.name]--
	}

	// выполняем команду от имени игрока
	if err := g.player.do(st.cmd, st.obj); err != nil {
		return gameOverError{nSteps: g.nSteps, err: err}
	}

	g.nSteps++
	return nil
}

// newGame() создает новую игру
func newGame() *game {
	p := newPlayer()
	things := map[label]int{
		apple.name:    2,
		coin.name:     3,
		mirror.name:   1,
		mushroom.name: 1,
	}
	return &game{p, things, 0}
}

// giveAdvice() возвращает совет, который
// поможет игроку избежать ошибки err в будущем
func giveAdvice(err error) string {
	// better, not my (extra advice interface) -> see BELOW

	var type1 invalidStepError
	var type2 notEnoughObjectsError
	var type3 commandLimitExceededError
	var type4 objectLimitExceededError

	switch {
	case errors.As(err, &type1):
		return type1.advice
	case errors.As(err, &type2):
		return type2.advice
	case errors.As(err, &type3):
		return type3.advice
	case errors.As(err, &type4):
		return type4.advice
	}
	// if advice = error message, no need for advice field and code above
	return err.Error()
}

// конец решения

/*
// better giveAdvice(), not my (extra advice interface):

type advice interface {
	Advice() string
}

func (e notEnoughObjectsError) Advice() string {
	return "be careful with scarce mushrooms"
}

func giveAdvice(err error) string {
	for {
		if advice, isAdvice := err.(advice); isAdvice {
			return advice.Advice()
		}

		err = errors.Unwrap(err)

		if err == nil {
			break
		}
	}

	return err.Error()
}
*/

func main() {
	gm := newGame()
	/*
		steps := []step{
			{eat, apple},
			{talk, bob},
			{take, coin},
			{eat, mushroom},
		}
	*/
	steps := []step{
		{talk, bob},
		{eat, mirror},
	}

	for _, st := range steps {
		if err := tryStep(gm, st); err != nil {
			os.Exit(1)
		}
	}
	fmt.Println("You win!")
}

// tryStep() выполняет шаг игры и печатает результат
func tryStep(gm *game, st step) error {
	fmt.Printf("trying to %s %s... ", st.cmd, st.obj.name)
	if err := gm.execute(st); err != nil {
		fmt.Println("FAIL", err)
		return err
	}
	fmt.Println("OK")
	return nil
}
