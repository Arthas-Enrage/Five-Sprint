package actioninfo

import (
	"fmt"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}


func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
    for i, data := range dataset {
        if err := dp.Parse(data); err != nil {
            fmt.Printf("Ошибка парсинга данных (строка %d): %v\n", i+1, err)
            continue
        }

        info, err := dp.ActionInfo()
        if err != nil {
            fmt.Printf("Ошибка получения информации (строка %d): %v\n", i+1, err)
            continue
        }

        fmt.Printf("%s\n", info)
    }
}
