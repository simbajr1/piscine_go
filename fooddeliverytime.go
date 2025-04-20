package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	// Define the menu with proper initialization of food structs
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	// Check if the item exists in the menu
	if item, exists := menu[order]; exists {
		return item.preptime
	}

	// Return 404 if the item doesn't exist
	return 404
}
