package prompts

type Order struct {
	Type typeN
	Acts []Act

	// Character
	Char Character
}

func New() *Order {
	return &Order{
		Acts: Load(),
	}
}

type Character struct {
	Series    string
	Character string
}

func (o *Order) Prompt(toJapanese bool) (actor, prompt string) {
	s := o.Acts[o.Type].Prompt
	if toJapanese {
		s += "上記命令に以後日本語で答えてください。"
	}
	return o.Acts[o.Type].Actor, s
}
