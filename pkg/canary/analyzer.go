package canary

// Analyzer manages canary deployments
type Analyzer struct {
    threshold float64
}

func NewAnalyzer(threshold float64) *Analyzer {
    return &Analyzer{threshold: threshold}
}
