package main

import (
	"log"
	"strings"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// commons is a component that displays a gamebook. A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type commons struct {
	app.Compo
	main app.Value
	storyboard app.Value
	button app.Value
	button2 app.Value
	scenes []Scene
	currentScene int
}

// Scene represents a single scene in the interactive novel
type Scene struct {
	ID 			 int
	Text         string
	Image 		 string
	FirstButton  string
	FirstLink    int
	SecondButton string
	SecondLink   int
}

func (c *commons) OnMount(ctx app.Context) {
	c.main = app.Window().Get("document").Call("querySelector", "#main")
	c.storyboard = app.Window().Get("document").Call("querySelector", "#story")
	c.button = app.Window().Get("document").Call("querySelector", "#button1")
	c.button2 = app.Window().Get("document").Call("querySelector", "#button2")

    c.loadScenes()

	c.currentScene = 0
	c.goToScene(c.currentScene)
}

// The Render method is where the component appearance is defined.
func (c *commons) Render() app.UI {
	return app.Div().ID("main").Body(
		app.Div().ID("story").Class("quote").Body(),
		app.Div().Class("button").Body(
			app.Span().ID("button1").OnClick(c.button1Clicked),
			app.Span().ID("button2").OnClick(c.button2Clicked),
		),
	)
}

func (c *commons) button1Clicked(ctx app.Context, e app.Event) {
	c.currentScene = c.scenes[c.currentScene].FirstLink;
  	c.goToScene(c.currentScene); 
}

func (c *commons) button2Clicked(ctx app.Context, e app.Event) {
	c.currentScene = c.scenes[c.currentScene].SecondLink;
  	c.goToScene(c.currentScene); 
}

func (c *commons) loadScenes() {
	c.scenes = []Scene{
		{
			ID: 		  0,
			Text:         "<h2> THE COMMONS - RECLAIM EARTH</h2>",
			Image: 		  "web/0.jpeg",
			FirstButton:  "Start game",
			FirstLink:    2,
			SecondButton: "About",
			SecondLink:   1,
		},
		{
			ID: 		  1,
			Text:         "<p>CREDITS</p><p>This is a gamebook created by Stateless Minds.</p><p>Visit https://github.com/stateless-minds/the-commons-reclaim-earth/</p>",
			SecondButton: "Start game",
			SecondLink:   2,
		},
		{
			ID: 		  2,
			Text:         "<p>THE GREAT RESET</p> Welcome traveler. The year is 2050. The great reset ended with a spectacular fiasco due to extreme energy needs which could not be sustainably met.",
			Image: 		  "web/2.jpeg",
			FirstButton:  "Previous",
			FirstLink:    0,
			SecondButton: "Next",
			SecondLink:   3,
		},
		{
			ID:		      3,
			Text:         "<p>ENERGY</p> The combined consumption of AI, robotics, data centers, electric cars and crypto currencies was so high and growing in geometric progression that hundreds of new power plants were needed on top of renewables.",
			Image: 		  "web/3.jpeg",
			FirstButton:  "Previous",
			FirstLink:    2,
			SecondButton: "Next",
			SecondLink:   4,
		},
		{
			ID: 		  4,
			Text:         "<p>COLLAPSE</p> Global supply chains and centralized systems collapsed due to the energy crisis. Elites resorted to bunkers and left the plebs to themselves.",
			Image: 		  "web/4.jpeg",
			FirstButton:  "Previous",
			FirstLink:    3,
			SecondButton: "Next",
			SecondLink:   5,
		},
		{
			ID: 		  5,
			Text:         "<p>CHAOS</p> Chaos errupted. People were so unprepared that the reaction was spontaneous and uncoordinated. Shops were looted until all supplies were exhausted.",
			Image: 		  "web/5.jpeg",
			FirstButton:  "Previous",
			FirstLink:    4,
			SecondButton: "Next",
			SecondLink:   6,
		},
		{
			ID:           6,
			Text:         "<p>IT'S ALL OVER</p>Institutions ceased to exist. Money didn't mean a thing. Private property was raided and abolished. Countries and borders were a thing of the past.",
			Image: 		  "web/6.jpeg",
			FirstButton:  "Previous",
			FirstLink:    5,
			SecondButton: "Next",
			SecondLink:   7,
		},
		{
			ID:           7,
			Text:         "<p>A NEW BEGINNING</p>Society has to be reestablished from scratch with our current level of knowledge, technology and understanding. The future of the world depends on you, your decisions and your fellow men.",
			Image: 		  "web/7.jpeg",
			FirstButton:  "Previous",
			FirstLink:    6,
			SecondButton: "Next",
			SecondLink:   8,
		},
		{
			ID:  	      8,
			Text:         "<p>STATE OR SELF-ORGANIZATION</p>You have to decide whether to reestablish the state or self-organize in a flat structure.",
			Image: 		  "web/8.jpeg",
			FirstButton:  "State",
			FirstLink:    9,
			SecondButton: "Self-organize",
			SecondLink:   10,
		},
		{
			ID: 		  9,
			Text:         "<p>STATE</p>Reestablishing states lead to new dictatorships and declaring martial law in order to maintain order. People lived in constant emergency state.",
			Image: 		  "web/8.jpeg",
			SecondButton: "Next",
			SecondLink:   14,
		},
		{	
			ID:           10,
			Text:         "<p>SELF-ORGANIZATION</p>Representative democracy didn't work out. It made people powerless by creating a new type of profession called politician. Direct is where you make the decisions. Liquid includes direct plus temporary delegation and revocation at anytime by topic.",
			Image: 		  "web/8.jpeg",
			FirstButton:  "Direct",
			FirstLink:    11,
			SecondButton: "Liquid",
			SecondLink:   12,
		},
		{
			ID: 		  11,
			Text:         "<p>DIRECT DEMOCRACY</p>Due to the dependency on activity the percentage of abastained is high. Most people don't have the knowledge and experience to vote in all areas of expertise.",
			Image: 		  "web/9.jpeg",
			SecondButton: "Next",
			SecondLink:   13,
		},
		{
			ID: 	      12,
			Text:         "<p>LIQUID DEMOCRACY</p>Voting directly and delegating/revoking voting rights to trustees and experts proved to be the best combination. Liquid democracy increased participation and encouraged learning by doing.",
			Image: 		  "web/10.jpeg",
			SecondButton: "Next",
			SecondLink:   13,
		},
		{
			ID:           13,
			Text:         "<p>VOTING IN THE CYBER AGE</p>With the new forms of democracy voting became a daily routine. Thanks to a simple peer-to-peer app which everyone has on their computing devices we didn't need any central servers or institutions to handle this.",
			Image: 		  "web/11.jpeg",
			FirstButton:  "Previous",
			FirstLink:    7,
			SecondButton: "Next",
			SecondLink:   14,
		},
		{
			ID:           14,
			Text:         "<p>PROPERTY</p>Besides the already tried forms of property - public and private we rediscovered a new one - the Commons. Now you have to decide which model to use.",
			Image: 		  "web/12.jpeg",
			FirstButton:  "Public/Private",
			FirstLink:    15,
			SecondButton: "The Commons",
			SecondLink:   16,
		},
		{
			ID: 		  15,
			Text:         "<p>PUBLIC/PRIVATE PARTNERSHIP</p>Public/private partnerships enhanced the joint venture power of corporations and states in a technocratic feudal dictatorship.",
			Image: 		  "web/12.jpeg",
			SecondButton: "Next",
			SecondLink:   17,
		},
		{
			ID:           16,
			Text:         "<p>THE COMMONS</p>With no states, governments and corporations Earth was declared a Commons and not treated as a mere resource to be exhausted anymore. All fences and borders were taken down. The Commons flourished once again as it used to be between the XIII and XVII centuries.",
			Image: 		  "web/12.jpeg",
			SecondButton: "Next",
			SecondLink:   18,
		},
		{
			ID:           17,
			Text:         "<p>MONEY</p>Money was restored as means of trade. It lead to the creation of programmed electronic money and made people live in a digital camp with constant surveillance and restrictions.",
			Image: 		  "web/13.jpeg",
			SecondButton: "Next",
			SecondLink:   19,
		},
		{
			ID: 		  18,
			Text:         "<p>MONEYLESS</p>The Commons made money obsolete. A sharing economy thrived. Previous malls and supermarkets turned into depots for taking and returning goods. All supply and demand was coordinated by an app.",
			Image: 		  "web/13.jpeg",
			SecondButton: "Next",
			SecondLink:   20,
		},
		{
			ID:           19,
			Text:         "<p>PRIVATE GOODS</p>Using money kept the class status of goods. Products were made with planned obsolescence and addiction features. They kept people disconnected and competing with each other.",
			Image: 		  "web/14.jpeg",
			SecondButton: "Next",
			SecondLink:   21,
		},
		{
			ID: 		  20,
			Text:         "<p>PUBLIC GOODS</p>Production changed from goods for private use to public ones. Ex-corporations turned cooperatives. Craftsmansip flourished. Personal computers were replaced by public terminals. Smart phones were replaced by public phones. Fridges, ovens, storage and other utilities became public services.",
			Image: 		  "web/14.jpeg",
			SecondButton: "Next",
			SecondLink:   22,
		},
		{
			ID:           21,
			Text:         "<p>PRIVATE PROPERTY</p>Private goods lead to keeping private housing. Price speculation made housing unaffordable for most people. To have a roof over their head they took debts for the rest of their lives which kept them enslaved and obedient.",
			Image: 		  "web/15.jpeg",
			SecondButton:  "Next",
			SecondLink:   23,
		},
		{
			ID: 		  22,
			Text:         "<p>SHARED HOUSING</p>Because everything was available for use anytime we freed up our wardrobes and spaces. We didn't need much else but a bed at home to rest in privacy. Everything else including food preparation could be done at communal spaces.",
			Image: 		  "web/15.jpeg",
			SecondButton: "Next",
			SecondLink:   24,
		},
		{
			ID: 		  23,
			Text:         "<p>LAWS</p>Private property requires laws in order to protect it from the rest. Laws were made to protect the wealthiest and most powerful at the expense of everybody else.",
			Image: 		  "web/16.jpeg",
			SecondButton: "Next",
			SecondLink:   25,
		},
		{
			ID: 		  24,
			Text:         "<p>LAWLESS</p>Without property and hierarchy laws became obsolete. Without the preconditions of ownership and division 99% of crime vanished. Any disputes were resolved through the liquid democracy app.",
			Image: 		  "web/16.jpeg",
			SecondButton: "Next",
			SecondLink:   25,
		},
		{
			ID: 		  25,
			Text:         "<p>MEDICINE</p>We are facing now a major decision to be voted. Would you rather go with healthcare or traditional healers?",
			Image: 		  "web/28.jpeg",
			FirstButton:  "Doctors",
			FirstLink:    26,
			SecondButton: "Healers",
			SecondLink:   27,
		},
		{
			ID: 		  26,
			Text:         "<p>HEALTHCARE</p>Life expectancy went up. Quality of life improved. But there were side effects with so much central power.",
			Image: 		  "web/29.jpeg",
			SecondButton: "Next",
			SecondLink:   28,
		},
		{
			ID:           27,
			Text:         "<p>HEALERS</p>We remained mortal thus maintaining a healthy natural balance of population. Life was considered a journey and not a goal.",
			Image: 		  "web/30.jpeg",
			SecondButton: "Next",
			SecondLink:   31,
		},
		{
			ID: 		  28,
			Text:         "<p>LOCKDOWNS</p>Lockdowns were imposed by the state and people were isolated as part of the new normal.",
			Image: 		  "web/17.jpeg",
			SecondButton: "Next",
			SecondLink:   29,
		},
		{
			ID: 		  29,
			Text:         "<p>OVERPOPULATION</p>With increased life length and less people dying from diseases the planet was overcrowded.",
			Image: 		  "web/17.jpeg",
			SecondButton: "Next",
			SecondLink:   30,
		},
		{
			ID: 		  30,
			Text:         "<p>GENOCIDE</p>Population was reduced by means of artificial food, stress, pandemics and wars.",
			Image: 		  "web/17.jpeg",
			SecondButton: "Next",
			SecondLink:   32,
		},
		{
			ID: 		  31,
			Text:         "<p>NOMADS</p>Since housing was on a first come first serve basis unless voted otherwise we all became nomads. We moved frequently driven by passion projects and community participation.",
			Image: 		  "web/17.jpeg",
			SecondButton: "Next",
			SecondLink:   33,
		},
		{
			ID: 		  32,
			Text:         "<p>ELITE AUTOMATION</p>Automation was used to make people redundant and easy to control.",
			Image: 		  "web/18.jpeg",
			SecondButton: "Next",
			SecondLink:   34,
		},
		{
			ID: 		  33,
			Text:         "<p>DEMOCRATIC AUTOMATION</p>With democractically decided automation where needed the definition of work diminished. All production was open doors without schedules and we can participate in anything anytime.",
			Image: 		  "web/19.jpeg",
			SecondButton: "Next",
			SecondLink:   35,
		},
		{
			ID: 	      34,
			Text:         "<p>MASS MEDIA</p>Mass media kept brainwashing people, keeping them busy and obedient, hyped on the next science discovery and new gadget.",
			Image: 		  "web/22.jpeg",
			SecondButton: "Next",
			SecondLink:   36,
		},
		{
			ID: 	      35,
			Text:         "<p>GOSSIP NEWS</p>Mass media got replaced by gossip apps where people reported what they observed personally. No more control of information, opinion and censorship.",
			Image: 		  "web/22.jpeg",
			SecondButton: "Next",
			SecondLink:   37,
		},
		{
			ID: 		  36,
			Text:         "<p>EDUCATION</p>Educational systems kept producing more workers which had no other role but to produce and consume till the end of their lives.",
			Image: 		  "web/23.jpeg",
			SecondButton: "Next",
			SecondLink:   38,
		},
		{
			ID: 		  37,
			Text:         "<p>FREE KNOWLEDGE</p>Education was replaced by voluntary free knowledge sourced from the internet and by real life practice. There were no ranks and certificates.",
			Image: 		  "web/23.jpeg",
			SecondButton: "Next",
			SecondLink:   39,
		},
		{
			ID: 		  38,
			Text:         "<p>FAMILY VALUES</p>Traditional family remained the pillar of society. People were living in personal bubbles tied to bloodlines and inheritance thus reproducing power and isolation.",
			Image: 		  "web/26.jpeg",
			SecondButton: "Next",
			SecondLink:   40,
		},
		{
			ID: 		  39,
			Text:         "<p>OPEN FAMILY</p>People stopped marrying and had open relationships with everyone. Sex was no longer a scarcity and a taboo. Crimes of revenge and jealousy were a thing of the past.",
			Image: 		  "web/26.jpeg",
			SecondButton: "Next",
			SecondLink:   41,
		},
		{
			ID: 		  40,
			Text:         "<p>OWN KIDS</p>Tradition with having children lead to population control. Parents were forbidden to have kids and instead newborns were breeded in laboratories.",
			Image: 		  "web/27.jpeg",
			SecondButton: "Next",
			SecondLink:   42,
		},
		{
			ID:           41,
			Text:         "<p>KIDS OF THE WORLD</p>We no longer had kids by necessity or tradition. All of them were raised as our own and there were no orphans. We voted to maintain a sustainable population of 1 billion.",
			Image: 		  "web/27.jpeg",
			SecondButton: "Next",
			SecondLink:   42,
		},
		{
			ID: 		  42,
			Text:         "<p>CITIES</p>Do you build new cities or keep expanding existing megalopolises.",
			Image: 		  "web/31.jpeg",
			FirstButton:  "Expand",
			FirstLink:    43,
			SecondButton: "Create new",
			SecondLink:   44,
		},
		{
			ID: 		  43,
			Text:         "<p>MEGALOPOLIS</p>Expanding traditional cities became more complex. Traffic congestions, pollution and overcrowding became the new normal.",
			Image: 		  "web/32.jpeg",
			SecondButton: "Next",
			SecondLink:   50,
		},
		{
			ID:           44,
			Text:         "<p>AUTONOMOUS CITIES</p>New autonomous cities were created from open-source blueprints. They were made of mobile houses, had bike paths instead of roads and no central systems.",
			Image: 		  "web/33.jpeg",
			SecondButton: "Next",
			SecondLink:   45,
		},
		{
			ID: 		  45,
			Text:         "<p>COMPOSTING</p>Composting replaced central disposal systems.",
			Image: 		  "web/34.jpeg",
			SecondButton: "Next",
			SecondLink:   46,
		},
		{
			ID: 		  46,
			Text:         "<p>GREENHOUSES</p>Growing local food in greenhouses replaced the imports of fresh food.",
			Image: 		  "web/35.jpeg",
			SecondButton: "Next",
			SecondLink:   47,
		},
		{
			ID: 	      47,
			Text:         "<p>ONE-MAN E-VEHICLES</p>One-man vehicles replaced cars. Bicycles and e-scooters were used for city transportation. E-carts for deliveries and light trams for supply.",
			Image: 		  "web/36.jpeg",
			SecondButton: "Next",
			SecondLink:   48,
		},
		{
			ID: 		  48,
			Text:         "<p>THE DOME</p>The autonomous cities were covered with domes to keep them air-conditioned with artificial watering and sun systems. This allowed year-round food and good weather for outdoor activities and living.",
			Image: 		  "web/37.jpeg",
			SecondButton: "Next",
			SecondLink:   49,
		},
		{
			ID: 		  49,
			Text:         "<p>DISASTER-FREE</p>Mobile houses and bike paths allowed the city to be moved within weeks in case of earthquakes. The dome protected the city from fires, pollution and floods.",
			Image: 		  "web/38.jpeg",
			SecondButton: "Next",
			SecondLink:   50,
		},
		{
			ID: 		  50,
			Text:         "<p>PARADISE OR OBLIVION</p>Dear reader, I hope this game managed to explain the interconnections between our beliefs, decisions and reality and to truly show you that a better world is only possible if we change ourselves.",
			Image: 		  "web/0.jpeg",
			FirstButton:  "Play again",
			FirstLink:    0,
			SecondButton: "About",
			SecondLink:   1,
		},
	}
}

func (c *commons) goToScene(sceneId int) {
	if c.scenes[sceneId].Image != "" {
		c.main.Set("style", "background: rgba(0, 0, 0, .65) url("+ c.scenes[sceneId].Image+"); background-size: cover;")
	} else {
		c.main.Set("style", "background: none")
	}

	if c.scenes[sceneId].Image != "web/0.jpeg" {
		c.storyboard.Set("innerHTML", "<div id=\"logo\"><span>THE COMMONS - RECLAIM EARTH<span></div> <span>" + strings.Replace(c.scenes[sceneId].Text, " ", " </span><span>", -1) + "</span>")
	} else {
		c.storyboard.Set("innerHTML", "<span>" + strings.Replace(c.scenes[sceneId].Text, " ", " </span><span>", -1) + "</span>")
	}

	if c.scenes[sceneId].FirstButton == "" {
		c.button.Set("innerHTML", "")
	} else {
		c.button.Set("innerHTML", "<span>" + c.scenes[sceneId].FirstButton + "</span>")
	}

	c.button2.Set("innerHTML", "<span>" + c.scenes[sceneId].SecondButton + "</span>")
  }

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", func() app.Composer { return &commons{} })

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "The Commons - Reclaimed Earth",
		Description: "Reclaim the Earth or vanish",
		Styles: []string{
			"/web/app.css", // Loads app.css file.
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
