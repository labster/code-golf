package routes

import "math/rand"

func pangramGrep() (args []string, out string) {
	all := [][]string{
		{"1", "A large fawn jumped quickly over white zinc boxes."},
		{"1", "A wizard's job is to vex chumps quickly in fog."},
		{"1", "Bright vixens jump; dozy fowl quack."},
		{"0", "Five or six big jet planes zoomed past the tower."},
		{"1", "Fix problem quickly with galvanized jets."},
		{"1", "Go, lazy fat vixen; be shrewd, jump quick."},
		{"1", "Jumpy halfling dwarves pick quartz box."},
		{"0", "My ex pub quiz crowd gave happy thanks."},
		{"1", "Pack my box with five dozen liquor jugs."},
		{"1", "The five boxing wizards jump quickly."},
		{"1", "The jay, pig, fox, zebra and my wolves quack!"},
		{"0", "The quick brown fox jumps over a lazy cat."},
		{"1", "Waxy and quivering, jocks fumble the pizza."},
		{"1", "When zombies arrive, quickly fax judge Pat."},
	}

	// Shuffle the whole set.
	for i := range all {
		j := rand.Intn(i + 1)
		all[i], all[j] = all[j], all[i]
	}

	for _, v := range all {
		args = append(args, v[1])

		if v[0] == "1" {
			out += v[1] + "\n"
		}
	}

	// Drop the trailing newline.
	out = out[:len(out)-1]

	return
}
