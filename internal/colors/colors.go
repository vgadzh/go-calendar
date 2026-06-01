package colors

const (
	// Reset color
	ResetColor = "\033[0m"

	// Regular Colors
	Black  = "\033[0;30m"
	Red    = "\033[0;31m"
	Green  = "\033[0;32m"
	Yellow = "\033[0;33m"
	Blue   = "\033[0;34m"
	Purple = "\033[0;35m"
	Cyan   = "\033[0;36m"
	White  = "\033[0;37m"

	// Faint
	FaintBlack  = "\033[2;30m"
	FaintRed    = "\033[2;31m"
	FaintGreen  = "\033[2;32m"
	FaintYellow = "\033[2;33m"
	FaintBlue   = "\033[2;34m"
	FaintPurple = "\033[2;35m"
	FaintCyan   = "\033[2;36m"
	FaintWhite  = "\033[2;37m"

	// Bold
	BoldBlack  = "\033[1;30m"
	BoldRed    = "\033[1;31m"
	BoldGreen  = "\033[1;32m"
	BoldYellow = "\033[1;33m"
	BoldBlue   = "\033[1;34m"
	BoldPurple = "\033[1;35m"
	BoldCyan   = "\033[1;36m"
	BoldWhite  = "\033[1;37m"

	// Underline
	UnderlineBlack  = "\033[4;30m"
	UnderlineRed    = "\033[4;31m"
	UnderlineGreen  = "\033[4;32m"
	UnderlineYellow = "\033[4;33m"
	UnderlineBlue   = "\033[4;34m"
	UnderlinePurple = "\033[4;35m"
	UnderlineCyan   = "\033[4;36m"
	UnderlineWhite  = "\033[4;37m"

	// High Intensity
	IntensiveBlack  = "\033[0;90m"
	IntensiveRed    = "\033[0;91m"
	IntensiveGreen  = "\033[0;92m"
	IntensiveYellow = "\033[0;93m"
	IntensiveBlue   = "\033[0;94m"
	IntensivePurple = "\033[0;95m"
	IntensiveCyan   = "\033[0;96m"
	IntensiveWhite  = "\033[0;97m"

	// Bold High Intensity
	BoldIntensiveBlack  = "\033[1;90m"
	BoldIntensiveRed    = "\033[1;91m"
	BoldIntensiveGreen  = "\033[1;92m"
	BoldIntensiveYellow = "\033[1;93m"
	BoldIntensiveBlue   = "\033[1;94m"
	BoldIntensivePurple = "\033[1;95m"
	BoldIntensiveCyan   = "\033[1;96m"
	BoldIntensiveWhite  = "\033[1;97m"

	// Background
	OnBlack  = "\033[40m"
	OnRed    = "\033[41m"
	OnGreen  = "\033[42m"
	OnYellow = "\033[43m"
	OnBlue   = "\033[44m"
	OnPurple = "\033[45m"
	OnCyan   = "\033[46m"
	OnWhite  = "\033[47m"

	// High Intensity backgrounds
	OnIntensiveBlack  = "\033[0;100m"
	OnIntensiveRed    = "\033[0;101m"
	OnIntensiveGreen  = "\033[0;102m"
	OnIntensiveYellow = "\033[0;103m"
	OnIntensiveBlue   = "\033[0;104m"
	OnIntensivePurple = "\033[0;105m"
	OnIntensiveCyan   = "\033[0;106m"
	OnIntensiveWhite  = "\033[0;107m"
)

// GetColoredString
func GetColoredString(text string, colors ...string) string {
	var colorList string
	for _, color := range colors {
		colorList += color
	}
	return colorList + text + ResetColor
}
