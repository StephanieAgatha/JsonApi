package entities

type MovieDetails struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	RelatedMovie []MovieDetails `json:"related_movie"`
}

var MovieData = []MovieDetails{
	{
		Id:          "001",
		Name:        "Transformers The Last Knight",
		Description: "Michael Bay was still on board for the fifth Transformers film, which marked the last live-action chapter in this particular timeline run.",
		RelatedMovie: []MovieDetails{
			{
				Id:          "012",
				Name:        "Transformers Age of Extinction",
				Description: "Mark Wahlberg joined the Transformers franchise in Age of Extinction, as Cade Yeager, a father and struggling inventor swept up in the latest feud between Autobots",
			},
			{
				Id:          "013",
				Name:        "Transformers Dark of the moon",
				Description: "Megan Fox was out and Rosie Huntington-Whiteley was in as Shia LeBeouf's new girlfriend",
			},
		},
	},
	{
		Id:          "002",
		Name:        "Spiderman No Way Home",
		Description: "Spider-Man: No Way Home picks up right after the post-credits of Spider-Man: Far From Home and is the most recent mainline Spider-Man film to be released.",
		RelatedMovie: []MovieDetails{
			{
				Id:          "014",
				Name:        "Spirderman Far For Home",
				Description: "Taking place after the events of Avengers: Endgame and dealing with much of what happened in that film",
			},
			{
				Id:          "015",
				Name:        "Avengers End Game",
				Description: "While certain plot reasons keep Tom Holland’s Spider-Man out of Avengers: Endgame for much of the moviejson",
			},
		},
	},
}
