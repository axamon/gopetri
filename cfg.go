package gopetri

// Cfg is config for petri network.
type Cfg struct {
	Start       string                   `json:"start"`
	Finish      []string                 `json:"finish"`
	Places      []string                 `json:"places"`
	Transitions map[string]CfgTransition `json:"transitions"`
}

// CfgTransition is struct of transition.
type CfgTransition struct {
	From []string `json:"from"`
	To   []string `json:"to"`
}
