package esepunittests

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	n := gc.calculateNumericalGrade() // compare without rounding

	switch {
	case n >= 90.0:
		return "A"
	case n >= 80.0:
		return "B"
	case n >= 70.0:
		return "C"
	case n >= 60.0:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	g := Grade{Name: name, Grade: grade, Type: gradeType}
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, g)
	case Exam:
		gc.exams = append(gc.exams, g)
	case Essay:
		gc.essays = append(gc.essays, g)
	}
}

// Weights: Assignments 50%, Exams 35%, Essays 15%.
func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	const wA, wE, wS = 0.50, 0.35, 0.15

	a, aOK := computeAverage(gc.assignments)
	e, eOK := computeAverage(gc.exams)
	s, sOK := computeAverage(gc.essays) // correct: essays!

	if !aOK {
		a = 0
	}
	if !eOK {
		e = 0
	}
	if !sOK {
		s = 0
	}

	return wA*a + wE*e + wS*s
}

func computeAverage(grades []Grade) (float64, bool) {
	if len(grades) == 0 {
		return 0, false
	}
	sum := 0.0
	for _, g := range grades {
		sum += float64(g.Grade)
	}
	return sum / float64(len(grades)), true
}
