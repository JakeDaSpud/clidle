package main;

// Imports
import(
	"fmt"
	"math/rand"
);

/* Game Variables */

// State
var quit bool = false; // Flag to end app
var in_shop = false;

// Word Bank
// Words that can be selected to be typed
var word_bank = [...]string{
	"dog", 
	"cat", 
	"camel",
	"frank",
	"donut",
	"d'oughnouts"};
var word_bank_count = len(word_bank); // Amount of words in word_bank

// Word Variables
var current_correct_word = ""; // Current word to be typed from word_bank
var word_value int = 1; // Point value of getting a word correct

// Player Variables
var player_input string = "";
var total_points int = 0.0;
var broke_message bool = false;

// Player's save data
type SaveData struct {
	
}



/* Shop Upgrades */

// Points Per Second Upgrade
var unlocked_pps_upgrade bool = false; // Whether the player has bought the Points per second upgrade
var pps_cost int = 150;
var pps_value int = 0; // Amount of Points earned every pps_timer seconds
var pps_timer float32 = 0.0; // Time between earning another pps cycle

// Increase Word Points Upgrade
var unlocked_increase_upgrade bool = false;
var increase_level int = 0;
var increase_cost int = word_value * 10;

// Multiple Words Upgrade
var unlocked_multiple_upgrade bool = false; // Whether the player has bought the Multiple words upgrade
var max_correct_words int = 1;



func spacer() {
	fmt.Println(" ---------- ");
}

func clear() {
	for i := 0; i < 100; i++ {
		fmt.Print("\n");
	}
}

func choose_word() string {
	var selected_word_idx int = rand.Intn(word_bank_count);
	return word_bank[selected_word_idx];
}

func correct_word(attempt string) bool {
	if (attempt == current_correct_word) {
		fmt.Println("+", word_value);
		return true;
	} else {
		return false;
	}
}

func save() {
	fmt.Println("savegamestuffgoeshere");
}

func shop() {
	for (in_shop) {
		clear();
		shop_menu();
		fmt.Scanln(&player_input);
		
		// Set to false, otherwise it shows up when not wanted
		broke_message = false;
		
		if (player_input == "back") {
			in_shop = false;
			clear();
		} else if (player_input == "quit") {
			in_shop = false;
			quit = true;
			break;
		} else if (player_input == "1" && total_points >= increase_cost) { // Increase Upgrade
			broke_message = false;
			unlocked_increase_upgrade = true;
			
			if (increase_level == 0) {
				word_value += 1;
			} else {
				word_value += increase_level;
			}
			
			// Exchange Points
			total_points -= increase_cost;
			increase_level += 1;
			
			increase_cost *= 2;
		} else {
			broke_message = true;
		}
	}
}

func shop_menu() {
	fmt.Println("\n-= Shop =-");
	fmt.Println("Type the number to buy the item, the cost is on the right.");
	
	fmt.Print("Points: ", total_points);
	if (broke_message) {
		fmt.Println(" <-- Can't afford that...");
	} else {
		fmt.Print("\n");
	}
	fmt.Println("Word Value:", word_value);
	
	spacer();
	
	fmt.Print("1. [Lvl ");
	fmt.Print(increase_level);
	fmt.Println("] Increase Points per Word:", increase_cost);
	
	spacer();
	
	fmt.Println("Type 'back' to leave the shop or 'quit' to exit the game");
}

func main() {
	fmt.Println("CLIdle Starting, welcome!");
	
	for (!quit) {
		
		if (player_input == "shop") {
			in_shop = true;
			shop();
		} else {
			clear();
		}
		
		// Quit check
		if (player_input == "quit" || quit) {
			quit = true;
			break;
		}
		
		current_correct_word = "";
		for i := 0; i < max_correct_words; i++ {
			current_correct_word += choose_word();
		}
		
		fmt.Println("Type 'shop' to buy upgrades");
		fmt.Println("Word Value:", word_value);
		
		spacer();
		
		fmt.Println("Points:", total_points);
		fmt.Println("Word: " + current_correct_word);
		fmt.Print("Type: ");
		
		// TODO:
		// - Error checking
		fmt.Scanln(&player_input);
		
		if (correct_word(player_input)) {
			total_points += word_value;
		}
	}
	
	save();
	fmt.Println("Goodbye!");
}
