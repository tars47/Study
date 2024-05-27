// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a chat bot to talk to you. Read lines of input from stdin and respond
// with random responses to predefined keywords.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {

	red := color.New(color.FgHiBlue).Add(color.Italic)
	green := color.New(color.FgGreen).Add(color.Italic)

	var jks []joke

	if err := json.Unmarshal([]byte(jokes), &jks); err != nil {
		log.Fatal("Error Parsing exiting")
	}

	fmt.Println(`🤡 Joker Alert 🤡, Time to laugh like a small baby 👶! Press Enter for a joke...`)

	reader := bufio.NewReader(os.Stdin)

	for {

		_, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		idx := rand.Intn(len(jks))

		jk := jks[idx]

		fmt.Print("> ")

		for _, r := range jk.Setup {
			delay := rand.Intn(100) + 1
			time.Sleep(time.Duration(delay) * time.Millisecond)
			red.Print(string(r))
		}

		fmt.Println()

		time.Sleep(time.Duration(3000) * time.Millisecond)
		fmt.Print("> ")

		for _, r := range jk.Punchline {
			delay := rand.Intn(100) + 1
			time.Sleep(time.Duration(delay) * time.Millisecond)
			green.Print(string(r))
		}
		fmt.Println()

	}

}

type joke struct {
	Setup     string
	Punchline string
}

var jokes string = `[
	{
	  "type": "general",
	  "setup": "What did the fish say when it hit the wall?",
	  "punchline": "Dam."
	},
	{
	  "type": "general",
	  "setup": "How do you make a tissue dance?",
	  "punchline": "You put a little boogie on it."
	},
	{
	  "type": "general",
	  "setup": "What's Forrest Gump's password?",
	  "punchline": "1Forrest1"
	},
	{
	  "type": "general",
	  "setup": "What do you call a belt made out of watches?",
	  "punchline": "A waist of time."
	},
	{
	  "type": "general",
	  "setup": "Why can't bicycles stand on their own?",
	  "punchline": "They are two tired"
	},
	{
	  "type": "general",
	  "setup": "How does a train eat?",
	  "punchline": "It goes chew, chew"
	},
	{
	  "type": "general",
	  "setup": "What do you call a singing Laptop?",
	  "punchline": "A Dell"
	},
	{
	  "type": "general",
	  "setup": "How many lips does a flower have?",
	  "punchline": "Tulips"
	},
	{
	  "type": "general",
	  "setup": "How do you organize an outer space party?",
	  "punchline": "You planet"
	},
	{
	  "type": "general",
	  "setup": "What kind of shoes does a thief wear?",
	  "punchline": "Sneakers"
	},
	{
	  "type": "general",
	  "setup": "What's the best time to go to the dentist?",
	  "punchline": "Tooth hurty."
	},
	{
	  "type": "knock-knock",
	  "setup": "Knock knock. \n Who's there? \n A broken pencil. \n A broken pencil who?",
	  "punchline": "Never mind. It's pointless."
	},
	{
	  "type": "knock-knock",
	  "setup": "Knock knock. \n Who's there? \n Cows go. \n Cows go who?",
	  "punchline": "No, cows go moo."
	},
	{
	  "type": "knock-knock",
	  "setup": "Knock knock. \n Who's there? \n Little old lady. \n Little old lady who?",
	  "punchline": "I didn't know you could yodel!"
	},
	{
	  "type": "programming",
	  "setup": "What's the best thing about a Boolean?",
	  "punchline": "Even if you're wrong, you're only off by a bit."
	},
	{
	  "type": "programming",
	  "setup": "What's the object-oriented way to become wealthy?",
	  "punchline": "Inheritance"
	},
	{
	  "type": "programming",
	  "setup": "Where do programmers like to hangout?",
	  "punchline": "The Foo Bar."
	},
	{
	  "type": "programming",
	  "setup": "Why did the programmer quit his job?",
	  "punchline": "Because he didn't get arrays."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the two silk worms in a race?",
	  "punchline": "It ended in a tie."
	},
	{
	  "type": "general",
	  "setup": "What do you call a laughing motorcycle?",
	  "punchline": "A Yamahahahaha."
	},
	{
	  "type": "general",
	  "setup": "A termite walks into a bar and says...",
	  "punchline": "'Where is the bar tended?'"
	},
	{
	  "type": "general",
	  "setup": "What does C.S. Lewis keep at the back of his wardrobe?",
	  "punchline": "Narnia business!"
	},
	{
	  "type": "programming",
	  "setup": "Why do programmers always mix up Halloween and Christmas?",
	  "punchline": "Because Oct 31 == Dec 25"
	},
	{
	  "type": "programming",
	  "setup": "A SQL query walks into a bar, walks up to two tables and asks...",
	  "punchline": "'Can I join you?'"
	},
	{
	  "type": "programming",
	  "setup": "How many programmers does it take to change a lightbulb?",
	  "punchline": "None that's a hardware problem"
	},
	{
	  "type": "programming",
	  "setup": "If you put a million monkeys at a million keyboards, one of them will eventually write a Java program",
	  "punchline": "the rest of them will write Perl"
	},
	{
	  "type": "programming",
	  "setup": "['hip', 'hip']",
	  "punchline": "(hip hip array)"
	},
	{
	  "type": "programming",
	  "setup": "To understand what recursion is...",
	  "punchline": "You must first understand what recursion is"
	},
	{
	  "type": "programming",
	  "setup": "There are 10 types of people in this world...",
	  "punchline": "Those who understand binary and those who don't"
	},
	{
	  "type": "general",
	  "setup": "What did the duck say when he bought lipstick?",
	  "punchline": "Put it on my bill"
	},
	{
	  "type": "general",
	  "setup": "What happens to a frog's car when it breaks down?",
	  "punchline": "It gets toad away"
	},
	{
	  "type": "general",
	  "setup": "did you know the first French fries weren't cooked in France?",
	  "punchline": "they were cooked in Greece"
	},
	{
	  "type": "programming",
	  "setup": "Which song would an exception sing?",
	  "punchline": "Can't catch me - Avicii"
	},
	{
	  "type": "knock-knock",
	  "setup": "Knock knock. \n Who's there? \n Opportunity.",
	  "punchline": "That is impossible. Opportunity doesn’t come knocking twice!"
	},
	{
	  "type": "programming",
	  "setup": "Why do Java programmers wear glasses?",
	  "punchline": "Because they don't C#."
	},
	{
	  "type": "general",
	  "setup": "Why did the mushroom get invited to the party?",
	  "punchline": "Because he was a fungi."
	},
	{
	  "type": "general",
	  "setup": "Do you know what the word 'was' was initially?",
	  "punchline": "Before was was was was was is."
	},
	{
	  "type": "general",
	  "setup": "I'm reading a book about anti-gravity...",
	  "punchline": "It's impossible to put down"
	},
	{
	  "type": "general",
	  "setup": "If you're American when you go into the bathroom, and American when you come out, what are you when you're in there?",
	  "punchline": "European"
	},
	{
	  "type": "general",
	  "setup": "Want to hear a joke about a piece of paper?",
	  "punchline": "Never mind...it's tearable"
	},
	{
	  "type": "general",
	  "setup": "I just watched a documentary about beavers.",
	  "punchline": "It was the best dam show I ever saw"
	},
	{
	  "type": "general",
	  "setup": "If you see a robbery at an Apple Store...",
	  "punchline": "Does that make you an iWitness?"
	},
	{
	  "type": "general",
	  "setup": "A ham sandwhich walks into a bar and orders a beer. The bartender says...",
	  "punchline": "I'm sorry, we don't serve food here"
	},
	{
	  "type": "general",
	  "setup": "Why did the Clydesdale give the pony a glass of water?",
	  "punchline": "Because he was a little horse"
	},
	{
	  "type": "general",
	  "setup": "If you boil a clown...",
	  "punchline": "Do you get a laughing stock?"
	},
	{
	  "type": "general",
	  "setup": "Finally realized why my plant sits around doing nothing all day...",
	  "punchline": "He loves his pot."
	},
	{
	  "type": "general",
	  "setup": "Don't look at the eclipse through a colander.",
	  "punchline": "You'll strain your eyes."
	},
	{
	  "type": "general",
	  "setup": "I bought some shoes from a drug dealer.",
	  "punchline": "I don't know what he laced them with, but I was tripping all day!"
	},
	{
	  "type": "general",
	  "setup": "Why do chicken coops only have two doors?",
	  "punchline": "Because if they had four, they would be chicken sedans"
	},
	{
	  "type": "general",
	  "setup": "What do you call a factory that sells passable products?",
	  "punchline": "A satisfactory"
	},
	{
	  "type": "general",
	  "setup": "When a dad drives past a graveyard: Did you know that's a popular cemetery?",
	  "punchline": "Yep, people are just dying to get in there"
	},
	{
	  "type": "general",
	  "setup": "Why did the invisible man turn down the job offer?",
	  "punchline": "He couldn't see himself doing it"
	},
	{
	  "type": "general",
	  "setup": "How do you make holy water?",
	  "punchline": "You boil the hell out of it"
	},
	{
	  "type": "general",
	  "setup": "I had a dream that I was a muffler last night.",
	  "punchline": "I woke up exhausted!"
	},
	{
	  "type": "general",
	  "setup": "Why is peter pan always flying?",
	  "punchline": "Because he neverlands"
	},
	{
	  "type": "programming",
	  "setup": "How do you check if a webpage is HTML5?",
	  "punchline": "Try it out on Internet Explorer"
	},
	{
	  "type": "general",
	  "setup": "What do you call a cow with no legs?",
	  "punchline": "Ground beef!"
	},
	{
	  "type": "general",
	  "setup": "I dropped a pear in my car this morning.",
	  "punchline": "You should drop another one, then you would have a pair."
	},
	{
	  "type": "programming",
	  "setup": "Lady: How do I spread love in this cruel world?",
	  "punchline": "Random Dude: [...💘]"
	},
	{
	  "type": "programming",
	  "setup": "A user interface is like a joke.",
	  "punchline": "If you have to explain it then it is not that good."
	},
	{
	  "type": "knock-knock",
	  "setup": "Knock knock. \n Who's there? \n Hatch. \n Hatch who?",
	  "punchline": "Bless you!"
	},
	{
	  "type": "general",
	  "setup": "What do you call sad coffee?",
	  "punchline": "Despresso."
	},
	{
	  "type": "general",
	  "setup": "Why did the butcher work extra hours at the shop?",
	  "punchline": "To make ends meat."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the hungry clock?",
	  "punchline": "It went back four seconds."
	},
	{
	  "type": "general",
	  "setup": "Well...",
	  "punchline": "That’s a deep subject."
	},
	{
	  "type": "general",
	  "setup": "Did you hear the story about the cheese that saved the world?",
	  "punchline": "It was legend dairy."
	},
	{
	  "type": "general",
	  "setup": "Did you watch the new comic book movie?",
	  "punchline": "It was very graphic!"
	},
	{
	  "type": "general",
	  "setup": "I started a new business making yachts in my attic this year...",
	  "punchline": "The sails are going through the roof."
	},
	{
	  "type": "general",
	  "setup": "I got hit in the head by a soda can, but it didn't hurt that much...",
	  "punchline": "It was a soft drink."
	},
	{
	  "type": "general",
	  "setup": "I can't tell if i like this blender...",
	  "punchline": "It keeps giving me mixed results."
	},
	{
	  "type": "general",
	  "setup": "I couldn't get a reservation at the library...",
	  "punchline": "They were fully booked."
	},
	{
	  "type": "programming",
	  "setup": "I was gonna tell you a joke about UDP...",
	  "punchline": "...but you might not get it."
	},
	{
	  "type": "programming",
	  "setup": "The punchline often arrives before the set-up.",
	  "punchline": "Do you know the problem with UDP jokes?"
	},
	{
	  "type": "programming",
	  "setup": "Why do C# and Java developers keep breaking their keyboards?",
	  "punchline": "Because they use a strongly typed language."
	},
	{
	  "type": "general",
	  "setup": "What do you give to a lemon in need?",
	  "punchline": "Lemonaid."
	},
	{
	  "type": "general",
	  "setup": "Never take advice from electrons.",
	  "punchline": "They are always negative."
	},
	{
	  "type": "general",
	  "setup": "Hey, dad, did you get a haircut?",
	  "punchline": "No, I got them all cut."
	},
	{
	  "type": "general",
	  "setup": "What time is it?",
	  "punchline": "I don't know... it keeps changing."
	},
	{
	  "type": "general",
	  "setup": "A weasel walks into a bar. The bartender says, \"Wow, I've never served a weasel before. What can I get for you?\"",
	  "punchline": "Pop,goes the weasel."
	},
	{
	  "type": "general",
	  "setup": "Bad at golf?",
	  "punchline": "Join the club."
	},
	{
	  "type": "general",
	  "setup": "Can a kangaroo jump higher than the Empire State Building?",
	  "punchline": "Of course. The Empire State Building can't jump."
	},
	{
	  "type": "general",
	  "setup": "Can February march?",
	  "punchline": "No, but April may."
	},
	{
	  "type": "general",
	  "setup": "Can I watch the TV?",
	  "punchline": "Yes, but don’t turn it on."
	},
	{
	  "type": "general",
	  "setup": "Dad, can you put my shoes on?",
	  "punchline": "I don't think they'll fit me."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the bread factory burning down?",
	  "punchline": "They say the business is toast."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the chameleon who couldn't change color?",
	  "punchline": "They had a reptile dysfunction."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the cheese factory that exploded in France?",
	  "punchline": "There was nothing left but de Brie."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the cow who jumped over the barbed wire fence?",
	  "punchline": "It was udder destruction."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the guy who invented Lifesavers?",
	  "punchline": "They say he made a mint."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the guy whose whole left side was cut off?",
	  "punchline": "He's all right now."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the kidnapping at school?",
	  "punchline": "It's ok, he woke up."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the Mexican train killer?",
	  "punchline": "He had loco motives"
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the new restaurant on the moon?",
	  "punchline": "The food is great, but there’s just no atmosphere."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the runner who was criticized?",
	  "punchline": "He just took it in stride"
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the scientist who was lab partners with a pot of boiling water?",
	  "punchline": "He had a very esteemed colleague."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the submarine industry?",
	  "punchline": "It really took a dive..."
	},
	{
	  "type": "general",
	  "setup": "Did you hear that David lost his ID in prague?",
	  "punchline": "Now we just have to call him Dav."
	},
	{
	  "type": "general",
	  "setup": "Did you hear that the police have a warrant out on a midget psychic ripping people off?",
	  "punchline": "It reads \"Small medium at large.\""
	},
	{
	  "type": "general",
	  "setup": "Did you hear the joke about the wandering nun?",
	  "punchline": "She was a roman catholic."
	},
	{
	  "type": "general",
	  "setup": "Did you hear the news?",
	  "punchline": "FedEx and UPS are merging. They’re going to go by the name Fed-Up from now on."
	},
	{
	  "type": "general",
	  "setup": "Did you hear the one about the guy with the broken hearing aid?",
	  "punchline": "Neither did he."
	},
	{
	  "type": "general",
	  "setup": "Did you know crocodiles could grow up to 15 feet?",
	  "punchline": "But most just have 4."
	},
	{
	  "type": "general",
	  "setup": "What do ghosts call their true love?",
	  "punchline": "Their ghoul-friend"
	},
	{
	  "type": "general",
	  "setup": "Did you know that protons have mass?",
	  "punchline": "I didn't even know they were catholic."
	},
	{
	  "type": "general",
	  "setup": "Did you know you should always take an extra pair of pants golfing?",
	  "punchline": "Just in case you get a hole in one."
	},
	{
	  "type": "general",
	  "setup": "Do I enjoy making courthouse puns?",
	  "punchline": "Guilty"
	},
	{
	  "type": "general",
	  "setup": "Do you know where you can get chicken broth in bulk?",
	  "punchline": "The stock market."
	},
	{
	  "type": "general",
	  "setup": "Do you want a brief explanation of what an acorn is?",
	  "punchline": "In a nutshell, it's an oak tree."
	},
	{
	  "type": "general",
	  "setup": "Ever wondered why bees hum?",
	  "punchline": "It's because they don't know the words."
	},
	{
	  "type": "general",
	  "setup": "Have you ever heard of a music group called Cellophane?",
	  "punchline": "They mostly wrap."
	},
	{
	  "type": "general",
	  "setup": "Have you heard of the band 1023MB?",
	  "punchline": "They haven't got a gig yet."
	},
	{
	  "type": "general",
	  "setup": "Have you heard the rumor going around about butter?",
	  "punchline": "Never mind, I shouldn't spread it."
	},
	{
	  "type": "general",
	  "setup": "How are false teeth like stars?",
	  "punchline": "They come out at night!"
	},
	{
	  "type": "general",
	  "setup": "How can you tell a vampire has a cold?",
	  "punchline": "They start coffin."
	},
	{
	  "type": "general",
	  "setup": "How come a man driving a train got struck by lightning?",
	  "punchline": "He was a good conductor."
	},
	{
	  "type": "general",
	  "setup": "How come the stadium got hot after the game?",
	  "punchline": "Because all of the fans left."
	},
	{
	  "type": "general",
	  "setup": "How did Darth Vader know what Luke was getting for Christmas?",
	  "punchline": "He felt his presents."
	},
	{
	  "type": "general",
	  "setup": "How did the hipster burn the roof of his mouth?",
	  "punchline": "He ate the pizza before it was cool."
	},
	{
	  "type": "general",
	  "setup": "How do hens stay fit?",
	  "punchline": "They always egg-cercise!"
	},
	{
	  "type": "general",
	  "setup": "How do locomotives know where they're going?",
	  "punchline": "Lots of training"
	},
	{
	  "type": "general",
	  "setup": "How do the trees get on the internet?",
	  "punchline": "They log on."
	},
	{
	  "type": "general",
	  "setup": "How do you find Will Smith in the snow?",
	  "punchline": " Look for fresh prints."
	},
	{
	  "type": "general",
	  "setup": "How do you fix a broken pizza?",
	  "punchline": "With tomato paste."
	},
	{
	  "type": "general",
	  "setup": "How do you fix a damaged jack-o-lantern?",
	  "punchline": "You use a pumpkin patch."
	},
	{
	  "type": "general",
	  "setup": "How do you get a baby alien to sleep?",
	  "punchline": " You rocket."
	},
	{
	  "type": "general",
	  "setup": "How do you get two whales in a car?",
	  "punchline": "Start in England and drive West."
	},
	{
	  "type": "general",
	  "setup": "How do you know if there’s an elephant under your bed?",
	  "punchline": "Your head hits the ceiling!"
	},
	{
	  "type": "general",
	  "setup": "How do you make a hankie dance?",
	  "punchline": "Put a little boogie in it."
	},
	{
	  "type": "general",
	  "setup": "How good are you at Power Point?",
	  "punchline": "I Excel at it."
	},
	{
	  "type": "general",
	  "setup": "How do you organize a space party?",
	  "punchline": "You planet."
	},
	{
	  "type": "general",
	  "setup": "How do you steal a coat?",
	  "punchline": "You jacket."
	},
	{
	  "type": "general",
	  "setup": "How do you tell the difference between a crocodile and an alligator?",
	  "punchline": "You will see one later and one in a while."
	},
	{
	  "type": "general",
	  "setup": "How does a dyslexic poet write?",
	  "punchline": "Inverse."
	},
	{
	  "type": "general",
	  "setup": "How does a French skeleton say hello?",
	  "punchline": "Bone-jour."
	},
	{
	  "type": "general",
	  "setup": "How does a penguin build it’s house?",
	  "punchline": "Igloos it together."
	},
	{
	  "type": "general",
	  "setup": "How does a scientist freshen their breath?",
	  "punchline": "With experi-mints!"
	},
	{
	  "type": "general",
	  "setup": "How does the moon cut his hair?",
	  "punchline": "Eclipse it."
	},
	{
	  "type": "general",
	  "setup": "How many apples grow on a tree?",
	  "punchline": "All of them!"
	},
	{
	  "type": "general",
	  "setup": "How many bones are in the human hand?",
	  "punchline": "A handful of them."
	},
	{
	  "type": "general",
	  "setup": "How many hipsters does it take to change a lightbulb?",
	  "punchline": "Oh, it's a really obscure number. You've probably never heard of it."
	},
	{
	  "type": "general",
	  "setup": "How many kids with ADD does it take to change a lightbulb?",
	  "punchline": "Let's go ride bikes!"
	},
	{
	  "type": "general",
	  "setup": "How many optometrists does it take to change a light bulb?",
	  "punchline": "1 or 2? 1... or 2?"
	},
	{
	  "type": "general",
	  "setup": "How many seconds are in a year?",
	  "punchline": "12. January 2nd, February 2nd, March 2nd, April 2nd.... etc"
	},
	{
	  "type": "general",
	  "setup": "How many South Americans does it take to change a lightbulb?",
	  "punchline": "A Brazilian"
	},
	{
	  "type": "general",
	  "setup": "How many tickles does it take to tickle an octopus?",
	  "punchline": "Ten-tickles!"
	},
	{
	  "type": "general",
	  "setup": "How much does a hipster weigh?",
	  "punchline": "An instagram."
	},
	{
	  "type": "general",
	  "setup": "How was the snow globe feeling after the storm?",
	  "punchline": "A little shaken."
	},
	{
	  "type": "general",
	  "setup": "Is the pool safe for diving?",
	  "punchline": "It deep ends."
	},
	{
	  "type": "general",
	  "setup": "Is there a hole in your shoe?",
	  "punchline": "No… Then how’d you get your foot in it?"
	},
	{
	  "type": "general",
	  "setup": "What did the spaghetti say to the other spaghetti?",
	  "punchline": "Pasta la vista, baby!"
	},
	{
	  "type": "general",
	  "setup": "What’s 50 Cent’s name in Zimbabwe?",
	  "punchline": "200 Dollars."
	},
	{
	  "type": "general",
	  "setup": "Want to hear a chimney joke?",
	  "punchline": "Got stacks of em! First one's on the house"
	},
	{
	  "type": "general",
	  "setup": "Want to hear a joke about construction?",
	  "punchline": "Nah, I'm still working on it."
	},
	{
	  "type": "general",
	  "setup": "Want to hear my pizza joke?",
	  "punchline": "Never mind, it's too cheesy."
	},
	{
	  "type": "general",
	  "setup": "What animal is always at a game of cricket?",
	  "punchline": "A bat."
	},
	{
	  "type": "general",
	  "setup": "What are the strongest days of the week?",
	  "punchline": "Saturday and Sunday...the rest are weekdays."
	},
	{
	  "type": "general",
	  "setup": "What biscuit does a short person like?",
	  "punchline": "Shortbread. "
	},
	{
	  "type": "general",
	  "setup": "What cheese can never be yours?",
	  "punchline": "Nacho cheese."
	},
	{
	  "type": "general",
	  "setup": "What creature is smarter than a talking parrot?",
	  "punchline": "A spelling bee."
	},
	{
	  "type": "general",
	  "setup": "What did celery say when he broke up with his girlfriend?",
	  "punchline": "She wasn't right for me, so I really don't carrot all."
	},
	{
	  "type": "general",
	  "setup": "What did Michael Jackson name his denim store?",
	  "punchline": "   Billy Jeans!"
	},
	{
	  "type": "general",
	  "setup": "What did one nut say as he chased another nut?",
	  "punchline": " I'm a cashew!"
	},
	{
	  "type": "general",
	  "setup": "What did one plate say to the other plate?",
	  "punchline": "Dinner is on me!"
	},
	{
	  "type": "general",
	  "setup": "What did one snowman say to the other snow man?",
	  "punchline": "Do you smell carrot?"
	},
	{
	  "type": "general",
	  "setup": "What did one wall say to the other wall?",
	  "punchline": "I'll meet you at the corner!"
	},
	{
	  "type": "general",
	  "setup": "What did Romans use to cut pizza before the rolling cutter was invented?",
	  "punchline": "Lil Caesars"
	},
	{
	  "type": "general",
	  "setup": "What did the 0 say to the 8?",
	  "punchline": "Nice belt."
	},
	{
	  "type": "general",
	  "setup": "What did the beaver say to the tree?",
	  "punchline": "It's been nice gnawing you."
	},
	{
	  "type": "general",
	  "setup": "What did the big flower say to the littler flower?",
	  "punchline": "Hi, bud!"
	},
	{
	  "type": "general",
	  "setup": "What did the Buffalo say to his little boy when he dropped him off at school?",
	  "punchline": "Bison."
	},
	{
	  "type": "general",
	  "setup": "What did the digital clock say to the grandfather clock?",
	  "punchline": "Look, no hands!"
	},
	{
	  "type": "general",
	  "setup": "What did the dog say to the two trees?",
	  "punchline": "Bark bark."
	},
	{
	  "type": "general",
	  "setup": "What did the Dorito farmer say to the other Dorito farmer?",
	  "punchline": "Cool Ranch!"
	},
	{
	  "type": "general",
	  "setup": "What did the fish say when it swam into a wall?",
	  "punchline": "Damn!"
	},
	{
	  "type": "general",
	  "setup": "What did the grape do when he got stepped on?",
	  "punchline": "He let out a little wine."
	},
	{
	  "type": "general",
	  "setup": "What did the judge say to the dentist?",
	  "punchline": "Do you swear to pull the tooth, the whole tooth and nothing but the tooth?"
	},
	{
	  "type": "general",
	  "setup": "What did the late tomato say to the early tomato?",
	  "punchline": "I’ll ketch up"
	},
	{
	  "type": "general",
	  "setup": "What did the left eye say to the right eye?",
	  "punchline": "Between us, something smells!"
	},
	{
	  "type": "general",
	  "setup": "What did the mountain climber name his son?",
	  "punchline": "Cliff."
	},
	{
	  "type": "general",
	  "setup": "What did the ocean say to the beach?",
	  "punchline": "Thanks for all the sediment."
	},
	{
	  "type": "general",
	  "setup": "What did the ocean say to the shore?",
	  "punchline": "Nothing, it just waved."
	},
	{
	  "type": "general",
	  "setup": "Why don't you find hippopotamuses hiding in trees?",
	  "punchline": "They're really good at it."
	},
	{
	  "type": "general",
	  "setup": "What did the pirate say on his 80th birthday?",
	  "punchline": "Aye Matey!"
	},
	{
	  "type": "general",
	  "setup": "What did the Red light say to the Green light?",
	  "punchline": "Don't look at me I'm changing!"
	},
	{
	  "type": "general",
	  "setup": "What did the scarf say to the hat?",
	  "punchline": "You go on ahead, I am going to hang around a bit longer."
	},
	{
	  "type": "general",
	  "setup": "What did the shy pebble wish for?",
	  "punchline": "That she was a little boulder."
	},
	{
	  "type": "general",
	  "setup": "What did the traffic light say to the car as it passed?",
	  "punchline": "Don't look I'm changing!"
	},
	{
	  "type": "general",
	  "setup": "What did the Zen Buddist say to the hotdog vendor?",
	  "punchline": "Make me one with everything."
	},
	{
	  "type": "general",
	  "setup": "What do birds give out on Halloween?",
	  "punchline": "Tweets."
	},
	{
	  "type": "general",
	  "setup": "What do I look like?",
	  "punchline": "A JOKE MACHINE!?"
	},
	{
	  "type": "general",
	  "setup": "What do prisoners use to call each other?",
	  "punchline": "Cell phones."
	},
	{
	  "type": "general",
	  "setup": "What do vegetarian zombies eat?",
	  "punchline": "Grrrrrainnnnnssss."
	},
	{
	  "type": "general",
	  "setup": "What do you call a bear with no teeth?",
	  "punchline": "A gummy bear!"
	},
	{
	  "type": "general",
	  "setup": "What do you call a bee that lives in America?",
	  "punchline": "A USB."
	},
	{
	  "type": "general",
	  "setup": "What do you call a boomerang that won't come back?",
	  "punchline": "A stick."
	},
	{
	  "type": "general",
	  "setup": "What do you call a careful wolf?",
	  "punchline": "Aware wolf."
	},
	{
	  "type": "general",
	  "setup": "What do you call a cow on a trampoline?",
	  "punchline": "A milk shake!"
	},
	{
	  "type": "general",
	  "setup": "What do you call a cow with two legs?",
	  "punchline": "Lean beef."
	},
	{
	  "type": "general",
	  "setup": "What do you call a crowd of chess players bragging about their wins in a hotel lobby?",
	  "punchline": "Chess nuts boasting in an open foyer."
	},
	{
	  "type": "general",
	  "setup": "What do you call a dad that has fallen through the ice?",
	  "punchline": "A Popsicle."
	},
	{
	  "type": "general",
	  "setup": "What do you call a dictionary on drugs?",
	  "punchline": "High definition."
	},
	{
	  "type": "general",
	  "setup": "what do you call a dog that can do magic tricks?",
	  "punchline": "a labracadabrador"
	},
	{
	  "type": "general",
	  "setup": "What do you call a droid that takes the long way around?",
	  "punchline": "R2 detour."
	},
	{
	  "type": "general",
	  "setup": "What do you call a duck that gets all A's?",
	  "punchline": "A wise quacker."
	},
	{
	  "type": "general",
	  "setup": "What do you call a fake noodle?",
	  "punchline": "An impasta."
	},
	{
	  "type": "general",
	  "setup": "What do you call a fashionable lawn statue with an excellent sense of rhythmn?",
	  "punchline": "A metro-gnome"
	},
	{
	  "type": "general",
	  "setup": "What do you call a fat psychic?",
	  "punchline": "A four-chin teller."
	},
	{
	  "type": "general",
	  "setup": "What do you call a fly without wings?",
	  "punchline": "A walk."
	},
	{
	  "type": "general",
	  "setup": "What do you call a girl between two posts?",
	  "punchline": "Annette."
	},
	{
	  "type": "general",
	  "setup": "What do you call a group of disorganized cats?",
	  "punchline": "A cat-tastrophe."
	},
	{
	  "type": "general",
	  "setup": "What do you call a group of killer whales playing instruments?",
	  "punchline": "An Orca-stra."
	},
	{
	  "type": "general",
	  "setup": "What do you call a monkey in a mine field?",
	  "punchline": "A babooooom!"
	},
	{
	  "type": "general",
	  "setup": "What do you call a nervous javelin thrower?",
	  "punchline": "Shakespeare."
	},
	{
	  "type": "general",
	  "setup": "What do you call a pig that knows karate?",
	  "punchline": "A pork chop!"
	},
	{
	  "type": "general",
	  "setup": "What do you call a pig with three eyes?",
	  "punchline": "Piiig"
	},
	{
	  "type": "general",
	  "setup": "What do you call a pile of cats?",
	  "punchline": " A Meowtain."
	},
	{
	  "type": "general",
	  "setup": "What do you call a sheep with no legs?",
	  "punchline": "A cloud."
	},
	{
	  "type": "general",
	  "setup": "What do you call a troublesome Canadian high schooler?",
	  "punchline": "A poutine."
	},
	{
	  "type": "general",
	  "setup": "What do you call an alligator in a vest?",
	  "punchline": "An in-vest-igator!"
	},
	{
	  "type": "general",
	  "setup": "What do you call an Argentinian with a rubber toe?",
	  "punchline": "Roberto"
	},
	{
	  "type": "general",
	  "setup": "What do you call an eagle who can play the piano?",
	  "punchline": "Talonted!"
	},
	{
	  "type": "general",
	  "setup": "What do you call an elephant that doesn’t matter?",
	  "punchline": "An irrelephant."
	},
	{
	  "type": "general",
	  "setup": "What do you call an old snowman?",
	  "punchline": "Water."
	},
	{
	  "type": "general",
	  "setup": "What do you call cheese by itself?",
	  "punchline": "Provolone."
	},
	{
	  "type": "general",
	  "setup": "What do you call corn that joins the army?",
	  "punchline": "Kernel."
	},
	{
	  "type": "general",
	  "setup": "What do you call someone with no nose?",
	  "punchline": "Nobody knows."
	},
	{
	  "type": "general",
	  "setup": "What do you call two barracuda fish?",
	  "punchline": " A Pairacuda!"
	},
	{
	  "type": "general",
	  "setup": "What do you do on a remote island?",
	  "punchline": "Try and find the TV island it belongs to."
	},
	{
	  "type": "general",
	  "setup": "What do you do when you see a space man?",
	  "punchline": "Park your car, man."
	},
	{
	  "type": "general",
	  "setup": "What do you get hanging from Apple trees?",
	  "punchline": "Sore arms."
	},
	{
	  "type": "general",
	  "setup": "What do you get when you cross a bee and a sheep?",
	  "punchline": "A bah-humbug."
	},
	{
	  "type": "general",
	  "setup": "What do you get when you cross a chicken with a skunk?",
	  "punchline": "A fowl smell!"
	},
	{
	  "type": "general",
	  "setup": "What do you get when you cross a rabbit with a water hose?",
	  "punchline": "Hare spray."
	},
	{
	  "type": "general",
	  "setup": "What do you get when you cross a snowman with a vampire?",
	  "punchline": "Frostbite."
	},
	{
	  "type": "general",
	  "setup": "What do you give a sick lemon?",
	  "punchline": "Lemonaid."
	},
	{
	  "type": "general",
	  "setup": "What does a clock do when it's hungry?",
	  "punchline": "It goes back four seconds!"
	},
	{
	  "type": "general",
	  "setup": "What does a female snake use for support?",
	  "punchline": "A co-Bra!"
	},
	{
	  "type": "general",
	  "setup": "What does a pirate pay for his corn?",
	  "punchline": "A buccaneer!"
	},
	{
	  "type": "general",
	  "setup": "What does an angry pepper do?",
	  "punchline": "It gets jalapeño face."
	},
	{
	  "type": "general",
	  "setup": "What happens when you anger a brain surgeon?",
	  "punchline": "They will give you a piece of your mind."
	},
	{
	  "type": "general",
	  "setup": "What has ears but cannot hear?",
	  "punchline": "A field of corn."
	},
	{
	  "type": "general",
	  "setup": "What is a centipedes's favorite Beatle song?",
	  "punchline": " I want to hold your hand, hand, hand, hand..."
	},
	{
	  "type": "general",
	  "setup": "What is a tornado's favorite game to play?",
	  "punchline": "Twister!"
	},
	{
	  "type": "general",
	  "setup": "What is a vampire's favorite fruit?",
	  "punchline": "A blood orange."
	},
	{
	  "type": "general",
	  "setup": "What is a witch's favorite subject in school?",
	  "punchline": "Spelling!"
	},
	{
	  "type": "general",
	  "setup": "What is red and smells like blue paint?",
	  "punchline": "Red paint!"
	},
	{
	  "type": "general",
	  "setup": "What is the difference between ignorance and apathy?",
	  "punchline": "I don't know and I don't care."
	},
	{
	  "type": "general",
	  "setup": "What is the hardest part about sky diving?",
	  "punchline": "The ground."
	},
	{
	  "type": "general",
	  "setup": "What is the leading cause of dry skin?",
	  "punchline": "Towels"
	},
	{
	  "type": "general",
	  "setup": "What is the least spoken language in the world?",
	  "punchline": "Sign Language"
	},
	{
	  "type": "general",
	  "setup": "What is the tallest building in the world?",
	  "punchline": "The library, it’s got the most stories!"
	},
	{
	  "type": "general",
	  "setup": "What is this movie about?",
	  "punchline": "It is about 2 hours long."
	},
	{
	  "type": "general",
	  "setup": "What kind of award did the dentist receive?",
	  "punchline": "A little plaque."
	},
	{
	  "type": "general",
	  "setup": "What kind of bagel can fly?",
	  "punchline": "A plain bagel."
	},
	{
	  "type": "general",
	  "setup": "What kind of dinosaur loves to sleep?",
	  "punchline": "A stega-snore-us."
	},
	{
	  "type": "general",
	  "setup": "What kind of dog lives in a particle accelerator?",
	  "punchline": "A Fermilabrador Retriever."
	},
	{
	  "type": "general",
	  "setup": "What kind of magic do cows believe in?",
	  "punchline": "MOODOO."
	},
	{
	  "type": "general",
	  "setup": "What kind of music do planets listen to?",
	  "punchline": "Nep-tunes."
	},
	{
	  "type": "general",
	  "setup": "What kind of pants do ghosts wear?",
	  "punchline": "Boo jeans."
	},
	{
	  "type": "general",
	  "setup": "What kind of tree fits in your hand?",
	  "punchline": "A palm tree!"
	},
	{
	  "type": "general",
	  "setup": "What lies at the bottom of the ocean and twitches?",
	  "punchline": "A nervous wreck."
	},
	{
	  "type": "general",
	  "setup": "What musical instrument is found in the bathroom?",
	  "punchline": "A tuba toothpaste."
	},
	{
	  "type": "general",
	  "setup": "What time did the man go to the dentist?",
	  "punchline": "Tooth hurt-y."
	},
	{
	  "type": "general",
	  "setup": "What type of music do balloons hate?",
	  "punchline": "Pop music!"
	},
	{
	  "type": "general",
	  "setup": "What was a more important invention than the first telephone?",
	  "punchline": "The second one."
	},
	{
	  "type": "general",
	  "setup": "What was the pumpkin’s favorite sport?",
	  "punchline": "Squash."
	},
	{
	  "type": "general",
	  "setup": "What's black and white and read all over?",
	  "punchline": "The newspaper."
	},
	{
	  "type": "general",
	  "setup": "What's blue and not very heavy?",
	  "punchline": " Light blue."
	},
	{
	  "type": "general",
	  "setup": "What's brown and sticky?",
	  "punchline": "A stick."
	},
	{
	  "type": "general",
	  "setup": "What's orange and sounds like a parrot?",
	  "punchline": "A Carrot."
	},
	{
	  "type": "general",
	  "setup": "What's red and bad for your teeth?",
	  "punchline": "A Brick."
	},
	{
	  "type": "general",
	  "setup": "What's the best thing about elevator jokes?",
	  "punchline": "They work on so many levels."
	},
	{
	  "type": "general",
	  "setup": "What's the difference between a guitar and a fish?",
	  "punchline": "You can tune a guitar but you can't \"tuna\"fish!"
	},
	{
	  "type": "general",
	  "setup": "What's the difference between a hippo and a zippo?",
	  "punchline": "One is really heavy, the other is a little lighter."
	},
	{
	  "type": "general",
	  "setup": "What's the difference between a seal and a sea lion?",
	  "punchline": "An ion! "
	},
	{
	  "type": "general",
	  "setup": "What's the worst part about being a cross-eyed teacher?",
	  "punchline": "They can't control their pupils."
	},
	{
	  "type": "general",
	  "setup": "What's the worst thing about ancient history class?",
	  "punchline": "The teachers tend to Babylon."
	},
	{
	  "type": "general",
	  "setup": "What’s brown and sounds like a bell?",
	  "punchline": "Dung!"
	},
	{
	  "type": "general",
	  "setup": "What’s E.T. short for?",
	  "punchline": "He’s only got little legs."
	},
	{
	  "type": "general",
	  "setup": "What’s Forest Gump’s Facebook password?",
	  "punchline": "1forest1"
	},
	{
	  "type": "general",
	  "setup": "What’s the advantage of living in Switzerland?",
	  "punchline": "Well, the flag is a big plus."
	},
	{
	  "type": "general",
	  "setup": "What’s the difference between an African elephant and an Indian elephant?",
	  "punchline": "About 5000 miles."
	},
	{
	  "type": "general",
	  "setup": "When do doctors get angry?",
	  "punchline": "When they run out of patients."
	},
	{
	  "type": "general",
	  "setup": "When does a joke become a dad joke?",
	  "punchline": "When it becomes apparent."
	},
	{
	  "type": "general",
	  "setup": "When is a door not a door?",
	  "punchline": "When it's ajar."
	},
	{
	  "type": "general",
	  "setup": "Where did you learn to make ice cream?",
	  "punchline": "Sunday school."
	},
	{
	  "type": "general",
	  "setup": "Where do bees go to the bathroom?",
	  "punchline": " The BP station."
	},
	{
	  "type": "general",
	  "setup": "Where do hamburgers go to dance?",
	  "punchline": "The meat-ball."
	},
	{
	  "type": "general",
	  "setup": "Where do rabbits go after they get married?",
	  "punchline": "On a bunny-moon."
	},
	{
	  "type": "general",
	  "setup": "Where do sheep go to get their hair cut?",
	  "punchline": "The baa-baa shop."
	},
	{
	  "type": "general",
	  "setup": "Where do you learn to make banana splits?",
	  "punchline": "At sundae school."
	},
	{
	  "type": "general",
	  "setup": "Where do young cows eat lunch?",
	  "punchline": "In the calf-ateria."
	},
	{
	  "type": "general",
	  "setup": "Where does batman go to the bathroom?",
	  "punchline": "The batroom."
	},
	{
	  "type": "general",
	  "setup": "Where does Fonzie like to go for lunch?",
	  "punchline": "Chick-Fil-Eyyyyyyyy."
	},
	{
	  "type": "general",
	  "setup": "Where does Napoleon keep his armies?",
	  "punchline": "In his sleevies."
	},
	{
	  "type": "general",
	  "setup": "Where was the Declaration of Independence signed?",
	  "punchline": "At the bottom! "
	},
	{
	  "type": "general",
	  "setup": "Where’s the bin?",
	  "punchline": "I haven’t been anywhere!"
	},
	{
	  "type": "general",
	  "setup": "Which side of the chicken has more feathers?",
	  "punchline": "The outside."
	},
	{
	  "type": "general",
	  "setup": "Who did the wizard marry?",
	  "punchline": "His ghoul-friend"
	},
	{
	  "type": "general",
	  "setup": "Who is the coolest Doctor in the hospital?",
	  "punchline": "The hip Doctor!"
	},
	{
	  "type": "general",
	  "setup": "Why are fish easy to weigh?",
	  "punchline": "Because they have their own scales."
	},
	{
	  "type": "general",
	  "setup": "Why are fish so smart?",
	  "punchline": "Because they live in schools!"
	},
	{
	  "type": "general",
	  "setup": "Why are ghosts bad liars?",
	  "punchline": "Because you can see right through them!"
	},
	{
	  "type": "general",
	  "setup": "Why are graveyards so noisy?",
	  "punchline": "Because of all the coffin."
	},
	{
	  "type": "general",
	  "setup": "Why are mummys scared of vacation?",
	  "punchline": "They're afraid to unwind."
	},
	{
	  "type": "general",
	  "setup": "Why are oranges the smartest fruit?",
	  "punchline": "Because they are made to concentrate. "
	},
	{
	  "type": "general",
	  "setup": "Why are pirates called pirates?",
	  "punchline": "Because they arrr!"
	},
	{
	  "type": "general",
	  "setup": "Why are skeletons so calm?",
	  "punchline": "Because nothing gets under their skin."
	},
	{
	  "type": "general",
	  "setup": "Why can't a bicycle stand on its own?",
	  "punchline": "It's two-tired."
	},
	{
	  "type": "general",
	  "setup": "Why can't you use \"Beef stew\"as a password?",
	  "punchline": "Because it's not stroganoff."
	},
	{
	  "type": "general",
	  "setup": "Why can't your nose be 12 inches long?",
	  "punchline": "Because then it'd be a foot!"
	},
	{
	  "type": "general",
	  "setup": "Why can’t you hear a pterodactyl go to the bathroom?",
	  "punchline": "The p is silent."
	},
	{
	  "type": "general",
	  "setup": "Why couldn't the kid see the pirate movie?",
	  "punchline": "Because it was rated arrr!"
	},
	{
	  "type": "general",
	  "setup": "Why couldn't the lifeguard save the hippie?",
	  "punchline": "He was too far out, man."
	},
	{
	  "type": "general",
	  "setup": "Why did Dracula lie in the wrong coffin?",
	  "punchline": "He made a grave mistake."
	},
	{
	  "type": "general",
	  "setup": "Why did Sweden start painting barcodes on the sides of their battleships?",
	  "punchline": "So they could Scandinavian."
	},
	{
	  "type": "general",
	  "setup": "Why did the A go to the bathroom and come out as an E?",
	  "punchline": "Because he had a vowel movement."
	},
	{
	  "type": "general",
	  "setup": "Why did the barber win the race?",
	  "punchline": "He took a short cut."
	},
	{
	  "type": "general",
	  "setup": "Why did the belt go to prison?",
	  "punchline": "He held up a pair of pants!"
	},
	{
	  "type": "general",
	  "setup": "Why did the burglar hang his mugshot on the wall?",
	  "punchline": "To prove that he was framed!"
	},
	{
	  "type": "general",
	  "setup": "Why did the chicken get a penalty?",
	  "punchline": "For fowl play."
	},
	{
	  "type": "general",
	  "setup": "Why did the coffee file a police report?",
	  "punchline": "It got mugged."
	},
	{
	  "type": "general",
	  "setup": "Why did the cookie cry?",
	  "punchline": "Because his mother was a wafer so long"
	},
	{
	  "type": "general",
	  "setup": "Why did the cookie cry?",
	  "punchline": "It was feeling crumby."
	},
	{
	  "type": "general",
	  "setup": "Why did the cowboy have a weiner dog?",
	  "punchline": "Somebody told him to get a long little doggy."
	},
	{
	  "type": "general",
	  "setup": "Why did the fireman wear red, white, and blue suspenders?",
	  "punchline": "To hold his pants up."
	},
	{
	  "type": "general",
	  "setup": "Why did the girl smear peanut butter on the road?",
	  "punchline": "To go with the traffic jam."
	},
	{
	  "type": "general",
	  "setup": "Why did the half blind man fall in the well?",
	  "punchline": "Because he couldn't see that well!"
	},
	{
	  "type": "general",
	  "setup": "Why did the house go to the doctor?",
	  "punchline": "It was having window panes."
	},
	{
	  "type": "general",
	  "setup": "Why did the kid cross the playground?",
	  "punchline": "To get to the other slide."
	},
	{
	  "type": "general",
	  "setup": "Why did the man put his money in the freezer?",
	  "punchline": "He wanted cold hard cash!"
	},
	{
	  "type": "general",
	  "setup": "Why did the man run around his bed?",
	  "punchline": "Because he was trying to catch up on his sleep!"
	},
	{
	  "type": "general",
	  "setup": "Why did the melons plan a big wedding?",
	  "punchline": "Because they cantaloupe!"
	},
	{
	  "type": "general",
	  "setup": "Why did the octopus beat the shark in a fight?",
	  "punchline": "Because it was well armed."
	},
	{
	  "type": "general",
	  "setup": "Why did the opera singer go sailing?",
	  "punchline": "They wanted to hit the high Cs."
	},
	{
	  "type": "general",
	  "setup": "Why did the scarecrow win an award?",
	  "punchline": "Because he was outstanding in his field."
	},
	{
	  "type": "general",
	  "setup": "Why did the tomato blush?",
	  "punchline": "Because it saw the salad dressing."
	},
	{
	  "type": "general",
	  "setup": "Why did the tree go to the dentist?",
	  "punchline": "It needed a root canal."
	},
	{
	  "type": "general",
	  "setup": "Why did the worker get fired from the orange juice factory?",
	  "punchline": "Lack of concentration."
	},
	{
	  "type": "general",
	  "setup": "Why didn't the number 4 get into the nightclub?",
	  "punchline": "Because he is 2 square."
	},
	{
	  "type": "general",
	  "setup": "Why didn’t the orange win the race?",
	  "punchline": "It ran out of juice."
	},
	{
	  "type": "general",
	  "setup": "Why didn’t the skeleton cross the road?",
	  "punchline": "Because he had no guts."
	},
	{
	  "type": "general",
	  "setup": "Why do bananas have to put on sunscreen before they go to the beach?",
	  "punchline": "Because they might peel!"
	},
	{
	  "type": "general",
	  "setup": "Why do bears have hairy coats?",
	  "punchline": "Fur protection."
	},
	{
	  "type": "general",
	  "setup": "Why do bees have sticky hair?",
	  "punchline": "Because they use honey combs!"
	},
	{
	  "type": "general",
	  "setup": "Why do bees hum?",
	  "punchline": "Because they don't know the words."
	},
	{
	  "type": "general",
	  "setup": "Why do birds fly south for the winter?",
	  "punchline": "Because it's too far to walk."
	},
	{
	  "type": "general",
	  "setup": "Why do choirs keep buckets handy?",
	  "punchline": "So they can carry their tune"
	},
	{
	  "type": "general",
	  "setup": "Why do crabs never give to charity?",
	  "punchline": "Because they’re shellfish."
	},
	{
	  "type": "general",
	  "setup": "Why do ducks make great detectives?",
	  "punchline": "They always quack the case."
	},
	{
	  "type": "general",
	  "setup": "Why do mathematicians hate the U.S.?",
	  "punchline": "Because it's indivisible."
	},
	{
	  "type": "general",
	  "setup": "Why do pirates not know the alphabet?",
	  "punchline": "They always get stuck at \"C\"."
	},
	{
	  "type": "general",
	  "setup": "Why do pumpkins sit on people’s porches?",
	  "punchline": "They have no hands to knock on the door."
	},
	{
	  "type": "general",
	  "setup": "Why do scuba divers fall backwards into the water?",
	  "punchline": "Because if they fell forwards they’d still be in the boat."
	},
	{
	  "type": "general",
	  "setup": "Why do trees seem suspicious on sunny days?",
	  "punchline": "Dunno, they're just a bit shady."
	},
	{
	  "type": "general",
	  "setup": "Why do valley girls hang out in odd numbered groups?",
	  "punchline": "Because they can't even."
	},
	{
	  "type": "general",
	  "setup": "Why do wizards clean their teeth three times a day?",
	  "punchline": "To prevent bat breath!"
	},
	{
	  "type": "general",
	  "setup": "Why do you never see elephants hiding in trees?",
	  "punchline": "Because they're so good at it."
	},
	{
	  "type": "general",
	  "setup": "Why does a chicken coop only have two doors?",
	  "punchline": "Because if it had four doors it would be a chicken sedan."
	},
	{
	  "type": "general",
	  "setup": "Why does a Moon-rock taste better than an Earth-rock?",
	  "punchline": "Because it's a little meteor."
	},
	{
	  "type": "general",
	  "setup": "Why does it take longer to get from 1st to 2nd base, than it does to get from 2nd to 3rd base?",
	  "punchline": "Because there’s a Shortstop in between!"
	},
	{
	  "type": "general",
	  "setup": "Why does Norway have barcodes on their battleships?",
	  "punchline": "So when they get back to port, they can Scandinavian."
	},
	{
	  "type": "general",
	  "setup": "Why does Superman get invited to dinners?",
	  "punchline": "Because he is a Supperhero."
	},
	{
	  "type": "general",
	  "setup": "Why does Waldo only wear stripes?",
	  "punchline": "Because he doesn't want to be spotted."
	},
	{
	  "type": "programming",
	  "setup": "Knock-knock.",
	  "punchline": "A race condition. Who is there?"
	},
	{
	  "type": "programming",
	  "setup": "What's the best part about TCP jokes?",
	  "punchline": "I get to keep telling them until you get them."
	},
	{
	  "type": "programming",
	  "setup": "A programmer puts two glasses on his bedside table before going to sleep.",
	  "punchline": "A full one, in case he gets thirsty, and an empty one, in case he doesn’t."
	},
	{
	  "type": "programming",
	  "setup": "There are 10 kinds of people in this world.",
	  "punchline": "Those who understand binary, those who don't, and those who weren't expecting a base 3 joke."
	},
	{
	  "type": "general",
	  "setup": "Two guys walk into a bar . . .",
	  "punchline": "The first guy says \"Ouch!\" and the second says \"Yeah, I didn't see it either.\""
	},
	{
	  "type": "programming",
	  "setup": "What did the router say to the doctor?",
	  "punchline": "It hurts when IP."
	},
	{
	  "type": "programming",
	  "setup": "An IPv6 packet is walking out of the house.",
	  "punchline": "He goes nowhere."
	},
	{
	  "type": "programming",
	  "setup": "A DHCP packet walks into a bar and asks for a beer.",
	  "punchline": "Bartender says, \"here, but I’ll need that back in an hour!\""
	},
	{
	  "type": "programming",
	  "setup": "3 SQL statements walk into a NoSQL bar. Soon, they walk out",
	  "punchline": "They couldn't find a table."
	},
	{
	  "type": "general",
	  "setup": "I saw a nice stereo on Craigslist for $1. Seller says the volume is stuck on ‘high’",
	  "punchline": "I couldn’t turn it down."
	},
	{
	  "type": "programming",
	  "setup": "What’s the object-oriented way to become wealthy?",
	  "punchline": "Inheritance."
	},
	{
	  "type": "general",
	  "setup": "What do you call a bee that can't make up its mind?",
	  "punchline": "A maybe."
	 },
	 {
	  "type": "general",
	  "setup": "Why was Cinderalla thrown out of the football team?",
	  "punchline": "Because she ran away from the ball."
	},
	{
	  "type": "general",
	  "setup": "What kind of music do welders like?",
	  "punchline": "Heavy metal."
	},
	{
	  "type": "general",
	  "setup": "Why are “Dad Jokes” so good?",
	  "punchline": "Because the punchline is apparent."
	},
	{
	  "type": "programming",
	  "setup": "Why dot net developers don't wear glasses?",
	  "punchline": "Because they see sharp."
	},
	{
	  "type": "general",
	  "setup": "Why is seven bigger than nine?",
	  "punchline": "Because seven ate nine."
	},
	{
	  "type": "dad",
	  "setup": "Why do fathers take an extra pair of socks when they go golfing?",
	  "punchline": "In case they get a hole in one!"
	},
	{
	  "type": "general",
	  "setup": "What do you call a suspicious looking laptop?",
	  "punchline": "Asus"
	},
	{
	  "type": "programming",
	  "setup": "What did the Java code say to the C code?",
	  "punchline": "You've got no class."
	},
	{
	  "type": "programming",
	  "setup": "What is the most used language in programming?",
	  "punchline": "Profanity."
	},
	{
	  "type": "programming",
	  "setup": "Why do programmers always get Christmas and Halloween mixed up?",
	  "punchline": "Because DEC 25 = OCT 31"
	},
	{
	  "type": "programming",
	  "setup": "What goes after USA?",
	  "punchline": "USB."
	},
	{
	  "type": "dad",
	  "setup": "Why don't eggs tell jokes?",
	  "punchline": "Because they would crack each other up."
	},
	{
	  "type": "general",
	  "setup": "How do you make the number one disappear?",
	  "punchline": "Add the letter G and it’s “gone”!"
	},
	{
	  "type": "general",
	  "setup": "My older brother always tore the last pages of my comic books, and never told me why.",
	  "punchline": "I had to draw my own conclusions."
	},
	{
	  "type": "general",
	  "setup": "The Sergeant-Major growled at the young soldier: I didn’t see you at camouflage training this morning.",
	  "punchline": "Thank you very much, sir."
	},
	{
	  "type": "general",
	  "setup": "Why does Waldo only wear stripes?",
	  "punchline": "Because he doesn't want to be spotted."
	},
	{
	  "type": "programming",
	  "setup": "A male developer often gets called as a Dev, then what would you call a female developer?",
	  "punchline": "Devi."
	},
	{
	  "type": "general",
	  "setup": "Why did the kid throw the watch out the window?",
	  "punchline": "So time would fly."
	},
	{
	  "type": "programming",
	  "setup": "Where did the API go to eat?",
	  "punchline": "To the RESTaurant."
	},
	{
	  "type": "general",
	  "setup": "Why did the rooster cross the road?",
	  "punchline": "He heard that the chickens at KFC were pretty hot."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the Viking who was reincarnated?",
	  "punchline": "He was Bjorn again"
	},
	{
	  "type": "general",
	  "setup": "What does the mermaid wear to math class?",
	  "punchline": "Algae-bra."
	},
	{
	  "type": "general",
	  "setup": "Did you hear about the crime in the parking garage?",
	  "punchline": "It was wrong on so many levels."
	},
	{
	  "type": "programming",
	  "setup": "Hey, wanna hear a joke?",
	  "punchline": "Parsing HTML with regex."
	},
	{
	  "type": "general",
	  "setup": "Why didn't the skeleton go for prom?",
	  "punchline": "Because it had nobody."
	},
	{
	  "type": "general",
	  "setup": "A grocery store cashier asked if I would like my milk in a bag.",
	  "punchline": "I told her 'No, thanks. The carton works fine.'"
	},
	{
	  "type": "general",
	  "setup": "99.9% of the people are dumb!",
	  "punchline": "Fortunately I belong to the remaining 1%"
	},
	{
	  "type": "programming",
	  "setup": "I just got fired from my job at the keyboard factory.",
	  "punchline": "They told me I wasn't putting in enough shifts."
	},
	{
	  "type": "general",
	  "setup": "You see, mountains aren't just funny.",
	  "punchline": "They are hill areas."
	},
	{
	  "type": "general",
	  "setup": "What do elves post on Social Media?",
	  "punchline": "Elf-ies."
	}
  ]`
