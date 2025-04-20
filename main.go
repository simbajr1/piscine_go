package piscine

func ReverseMenuIndex(menu []string) []string {
	reversedMenu := make([]string, len(menu))

	for i := range menu {
		reversedMenu[len(menu)-1-i] = menu[i]
	}
	return reversedMenu
}
