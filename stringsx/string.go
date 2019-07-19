package stringsx

import "fmt"

func Hello(name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}

func HelloByLang(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "zh":
		return fmt.Sprintf("你好, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	default:
		return fmt.Sprintf("unknow , %s!", name), nil
	}
}
