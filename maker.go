package prompts

import "fmt"

type Order struct {
	Type TypeN
	Acts []Act

	// Character
	Char Character
}

func New() *Order {
	return &Order{
		Acts: Load(),
	}
}

// Character set UniqueName for ActorCharacterfromMovieBookAnything and english.
type Character struct {
	Series    string
	Character string
}

func (o *Order) Prompt(toJapanese bool) (actor, prompt string) {
	s := o.Acts[o.Type].Prompt
	if o.Type == ActorCharacterfromMovieBookAnything {
		s = fmt.Sprintf(s,
			o.Char.Character,
			o.Char.Series,
			o.Char.Character,
			o.Char.Character,
			o.Char.Character,
			o.Char.Character,
			o.Char.Character)
	}
	if toJapanese {
		s += "上記命令に以後日本語で答えてください。"
	}
	return o.Acts[o.Type].Actor, s
}
