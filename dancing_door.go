// Package dancingdoor provides a function to quickly generate easily remembered names
package dancingdoor

import (
	"math/rand"
	"strings"
)

var colors = []string{
	"white", "black", "yellow", "red", "blue", "brown", "green",
	"purple", "orange", "silver", "scarlet", "rainbow", "indigo",
	"ivory", "navy", "olive", "teal", "pink", "magenta", "maroon",
	"sienna", "gold", "golden",
}

var adjectives = []string{
	"abandoned", "aberrant", "accidentally", "aggressive", "aimless",
	"alien", "angry", "appropriate", "barbaric", "beacon", "big", "bitter",
	"bleeding", "brave", "brutal", "cheerful", "dancing", "dangerous",
	"dead", "deserted", "digital", "dirty", "disappointed", "discarded",
	"dreaded", "eastern", "eastern", "elastic", "empty", "endless",
	"essential", "eternal", "everyday", "fierce", "flaming", "flying",
	"forgotten", "forsaken", "freaky", "frozen", "full", "furious", "ghastly",
	"global", "gloomy", "grim", "gruesome", "gutsy", "helpless", "hidden",
	"hideous", "homeless", "hungry", "insane", "intense", "intensive",
	"itchy", "liquid", "lone", "lost", "meaningful", "modern",
	"monday's", "morbid", "moving", "needless", "nervous", "new", "next",
	"ninth", "nocturnal", "northernmost", "official", "old", "permanent",
	"persistent", "pointless", "pure", "quality", "random", "rare", "raw",
	"reborn", "remote", "restless", "rich", "risky", "rocky", "rough",
	"running", "rusty", "sad", "saturday's", "screaming", "serious",
	"severe", "silly", "skilled", "sleepy", "sliding", "small", "solid",
	"steamy", "stony", "stormy", "straw", "strawberry", "streaming",
	"strong", "subtle", "supersonic", "surreal", "tainted", "temporary", "third", "tidy",
	"timely", "unique", "vital", "western", "wild", "wooden", "worthy", "bitter",
	"boiling", "brave", "cloudy", "cold", "confidential", "dreadful", "dusty", "eager",
	"early", "grotesque", "harsh", "heavy", "hollow", "hot", "husky", "icy",
	"late", "lonesome", "long", "lucky", "massive", "maximum", "minimum",
	"mysterious", "outstanding", "rapid", "rebel", "scattered", "shiny",
	"solid", "square", "steady", "steep", "sticky", "stormy", "strong",
	"sunday's", "swift", "tasty",
}

var defaultSuffixes = []string{
	"alarm", "albatross", "anaconda", "antique", "artificial", "autopsy",
	"autumn", "avenue", "backpack", "balcony", "barbershop", "boomerang",
	"bulldozer", "butter", "canal", "cloud", "clown", "coffin", "comic",
	"compass", "cosmic", "crayon", "creek", "crossbow", "dagger", "dinosaur",
	"dog", "donut", "door", "doorstop", "electrical", "electron", "eyelid",
	"firecracker", "fish", "flag", "flannel", "flea", "frostbite", "gravel",
	"haystack", "helium", "kangaroo", "lantern", "leather", "limousine",
	"lobster", "locomotive", "logbook", "longitude", "metaphor", "microphone",
	"monkey", "moose", "morning", "mountain", "mustard", "neutron", "nitrogen",
	"notorious", "obscure", "ostrich", "oyster", "parachute", "peasant",
	"pineapple", "plastic", "postal", "pottery", "proton", "puppet", "railroad",
	"rhinestone", "roadrunner", "rubber", "scarecrow", "scoreboard", "scorpion",
	"shower", "skunk", "sound", "street", "subdivision", "summer", "sunshine",
	"tea", "temple", "test", "tire", "tombstone", "toothbrush", "torpedo",
	"toupee", "trendy", "trombone", "tuba", "tuna", "tungsten", "vegetable",
	"venom", "vulture", "waffle", "warehouse", "waterbird", "weather", "weeknight",
	"windshield", "winter", "wrench", "xylophone", "alpha", "arm", "beam", "beta",
	"bird", "breeze", "burst", "cat", "cobra", "crystal", "drill", "eagle",
	"emerald", "epsilon", "finger", "fist", "foot", "fox", "galaxy", "gamma",
	"hammer", "heart", "hook", "hurricane", "iron", "jazz", "jupiter", "knife",
	"lama", "laser", "lion", "mars", "mercury", "moon", "moose", "neptune",
	"omega", "panther", "planet", "pluto", "plutonium", "poseidon", "python",
	"ray", "sapphire", "scissors", "screwdriver", "serpent", "sledgehammer",
	"smoke", "snake", "space", "spider", "star", "steel", "storm", "sun",
	"swallow", "tiger", "uranium", "venus", "viper", "wrench", "yard", "zeus",
}

var defaultCorpus = append(colors, adjectives...) // Concatenate the two slices

const defaultSeparator = " "

// Options holds the parameters for generating a code name.
type Options struct {
	Corpus    []string
	Suffixes  []string
	Separator string
}

var defaultOptions = Options{
	Corpus:    defaultCorpus,
	Suffixes:  defaultSuffixes,
	Separator: defaultSeparator,
}

// MakeOptions yields a ready to use Options struct with defaults provided.
func MakeOptions() *Options {
	options := defaultOptions
	return &options
}

// WithCorpus returns a copy of the provided options structure with corpus set to the provided string array.
func (opts *Options) WithCorpus(corpus []string) *Options {
	opts.Corpus = corpus
	return opts
}

// WithSuffixes returns a copy of the provided options structure with suffixes set to the provided string array.
func (opts *Options) WithSuffixes(suffixes []string) *Options {
	opts.Suffixes = suffixes
	return opts
}

// WithSeparator returns a copy of the provided options structure with separator set to the provided string.
func (opts *Options) WithSeparator(separator string) *Options {
	opts.Separator = separator
	return opts
}

func ensureDefaults(opts *Options) *Options {
	if len(opts.Corpus) == 0 {
		opts.Corpus = defaultCorpus
	}
	if len(opts.Suffixes) == 0 {
		opts.Suffixes = defaultSuffixes
	}
	if opts.Separator == "" {
		opts.Separator = defaultSeparator
	}
	return opts
}

// CodenameElements generates a name composed of two or three parts.
// The parts are returned as an array of strings.
func CodenameElements(options *Options) []string {

	name1 := options.Corpus[rand.Intn(len(options.Corpus))]
	name2 := options.Corpus[rand.Intn(len(options.Corpus))]
	suffix := options.Suffixes[rand.Intn(len(options.Suffixes))]
	i := rand.Intn(100)

	var names []string

	switch {
	case i <= 15:
		names = []string{name1, name2, suffix}
	case i > 15 && i < 31:
		names = []string{suffix, name1}
	default:
		names = []string{name1, suffix}
	}
	return names
}

// Codename generates a name composed of two or three parts.
// The parts are joined with the separator configured in Options or a space
// by default and returned as a single string
func Codename(opts ...*Options) string {
	var options *Options
	if len(opts) == 0 {
		options = MakeOptions()
	} else {
		options = ensureDefaults(opts[0])
	}
	return strings.Join(CodenameElements(options), options.Separator)
}
