package jokes

import (
	"crypto/rand"
	"math/big"

	"github.com/rs/zerolog/log"
)

var jokes = initJokes()

func initJokes() []string {
	return []string{
		"How do trains eat? Chew chew!",
		"How do you catch a unique rabbit? Unique up on it.",
		"How do you make a Kleenex dance? Put a little boogie in it.",
		"How do you organize a space party? You planet.",
		"How does a penguin build its house? Igloos it together.",
		"How does the moon cut his hair? Eclipse it.",
		"What do you call a bear with no teeth? A gummy bear.",
		"What do you call a bee that can't make up its mind? A Maybe.",
		"What do you call a can opener that doesn't work? A can't opener.",
		"What do you call a cow with no legs? Ground beef.",
		"What do you call a deer with no eyes? No-eye-deer.",
		"What do you call a dog magician? A labracadabrador.",
		"What do you call a fake noodle? An impasta.",
		"What do you call a fish with no eyes? A fsh.",
		"What do you call a sleeping bag? A nap sack.",
		"What do you call a sleeping bull? A bulldozer.",
		"What do you call a snowman with a six-pack? An abominable snowman.",
		"What do you call cheese that isn't yours? Nacho cheese.",
		"What do you get when you cross a detective and a dinosaur? Sherlock Bones.",
		"What do you get when you cross a snowman and a vampire? Frostbite.",
		"Why couldn't the bicycle stand up by itself? It was two-tired.",
		"Why did the banana go to the doctor? Because it wasn't peeling well.",
		"Why did the bicycle fall over? Because it was two-tired.",
		"Why did the book go to jail? Because it had a bad cover.",
		"Why did the chicken join a band? Because it had the drumsticks.",
		"Why did the coffee file a police report? It got mugged.",
		"Why did the computer go to the doctor? It had a virus.",
		"Why did the golfer bring two pairs of pants? In case he got a hole in one.",
		"Why did the grape stop in the middle of the road? Because it ran out of juice.",
		"Why did the math book look sad? Because it had too many problems.",
		"Why did the melons get married? Because they 'cantaloupe'.",
		"Why did the painting go to jail? For framing.",
		"Why did the pencil break up with the paper? Because it found out it was being erased.",
		"Why did the picture go to jail? Because it was framed.",
		"Why did the scarecrow become a comedian? Because he was outstanding in his field.",
		"Why did the skeleton go to the party alone? Because it had no body to go with.",
		"Why did the strawberry go to the gym? To get jammin'.",
		"Why did the student bring a ladder to school? He wanted to go to high school.",
		"Why did the tomato turn red? Because it saw the salad dressing.",
		"Why don't scientists trust atoms? Because they make up everything!",
	}
}

func GetJoke() (string, error) {
	random, err := rand.Int(rand.Reader, big.NewInt(int64(len(jokes))))
	if err != nil {
		log.Error().Err(err).Msg("Error generating a random number to fetch a joke")
		return "", err
	}

	return jokes[random.Int64()], nil
}
