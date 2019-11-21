package bot

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strconv"
)

type Menu struct {
	Greetings string
	Menu      map[int]string
}

type Config struct {
	Greetings string
	End       string
	Start     string
	Employee  map[int]int
	Channel int64
	Token     string
	Timeout   int
	Menu      map[int]Menu
	current   int
	answers	  map[int]int
}

func (menu Menu) string() string {
	result := ""
	for i := 0; i < 100; i++ {
		if val, ok := menu.Menu[i]; ok {
			result += fmt.Sprintf("\n%v: %v\n", i, val)
		}
	}
	return result
}
func (menu Menu) copy() Menu {
	newMenu :=  Menu{}
	newMenu.Greetings = menu.Greetings
	mn := make(map[int]string, 0)
	for key, value := range menu.Menu {
		k := key
		val := value
		mn[k] = val
	}
	newMenu.Menu = mn
	return newMenu
}


func Parse(yamlByte []byte) (config *Config, err error) {
	var conf Config
	config = &conf
	if err = yaml.Unmarshal(yamlByte, &config); err != nil {
		return
	}
	conf.current = 999
	config.answers = make(map[int]int, 0)
	for key, _ := range config.Menu {
		config.answers[key] = 0
	}
	return
}

func (config *Config) getMenu(input int) string {
	menu := config.Menu[1]
	if res, ok :=config.Menu[input]; ok {
		menu = res
	}
	return menu.Greetings + menu.string()

}

func (config *Config) Copy() *Config{
	newConfig := Config{}
	newConfig.Greetings = config.Greetings
	newConfig.End = config.End
	newConfig.Start = config.Start
	newConfig.Channel = config.Channel
	newConfig.Timeout = config.Timeout
	newConfig.current = config.current
	emp := make(map[int]int, 0)
	for key, value := range config.Employee {
		k := key
		val := value
		emp[k] = val
	}
	newConfig.Employee = emp
	ans := make(map[int]int, 0)
	for key, value := range config.answers {
		k := key
		val := value
		ans[k] = val
	}
	newConfig.answers = ans

	mn := make(map[int]Menu, 0)
	for key, menu := range config.Menu {
		k := key
		mn[k] = menu.copy()
	}
	newConfig.Menu = mn
	return &newConfig
}

func (config *Config) next(answer int) int {
	config.answers[config.current] = answer
	if _, ok := config.Menu[config.current+1]; ok {
		config.current = config.current + 1
	}
	return config.current
}


func (config *Config) Answer(input string) (string, *int64) {
	result := ""
	if config.current == 999 {
		result += config.Greetings
		config.current = 0
	} else if input == config.Start {
		config.current = 0
        result += config.getMenu(config.current);
	} else {
		i, err := strconv.Atoi(input)
		if err == nil {
			if worker, ok := config.Employee[config.current]; ok {
				if worker == i {
					emp := config.Channel
					return config.End, &emp
				}
			}
			result += config.getMenu(config.next(i));
		} else {
			result += config.getMenu(config.current);
		}
    }
	return result, nil
}
