package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on version control", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()
	// switched for specifications
	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 58, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a", 70, Assignment)
	gc.AddGrade("e", 70, Exam)
	gc.AddGrade("s", 70, Essay)
	if got := gc.GetFinalGrade(); got != "C" {
		t.Fatalf("want C, got %s", got)
	}
}

func TestGetGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a", 60, Assignment)
	gc.AddGrade("e", 60, Exam)
	gc.AddGrade("s", 60, Essay)
	if got := gc.GetFinalGrade(); got != "D" {
		t.Fatalf("want D, got %s", got)
	}
}

// Boundary equality checks: exact cutoffs should map to higher letter
func TestBoundaries_EqualCutoffs(t *testing.T) {
	cases := []struct {
		score int
		want  string
	}{
		{60, "D"},
		{70, "C"},
		{80, "B"},
		{90, "A"},
	}
	for _, c := range cases {
		gc := NewGradeCalculator()
		gc.AddGrade("a", c.score, Assignment)
		gc.AddGrade("e", c.score, Exam)
		gc.AddGrade("s", c.score, Essay)
		if got := gc.GetFinalGrade(); got != c.want {
			t.Fatalf("score=%d => want %s, got %s", c.score, c.want, got)
		}
	}
}

// Multiple items per category to exercise averaging logic
func TestMultipleItemsPerCategory_AverageUsed(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 80, Assignment)
	gc.AddGrade("a2", 100, Assignment)
	gc.AddGrade("e1", 70, Exam)
	gc.AddGrade("e2", 75, Exam)
	gc.AddGrade("e3", 85, Exam)
	gc.AddGrade("s1", 100, Essay)
	if got := gc.GetFinalGrade(); got != "B" {
		t.Fatalf("want B, got %s", got)
	}
}

// Missing categories must count as 0
func TestMissingCategories_CountAsZero(t *testing.T) {
	ga := NewGradeCalculator()
	ga.AddGrade("a", 100, Assignment)
	if got := ga.GetFinalGrade(); got != "F" {
		t.Fatalf("only assignments 100 => want F, got %s", got)
	}
	ge := NewGradeCalculator()
	ge.AddGrade("e", 100, Exam)
	if got := ge.GetFinalGrade(); got != "F" {
		t.Fatalf("only exams 100 => want F, got %s", got)
	}
	gs := NewGradeCalculator()
	gs.AddGrade("s", 100, Essay)
	if got := gs.GetFinalGrade(); got != "F" {
		t.Fatalf("only essays 100 => want F, got %s", got)
	}

	gz := NewGradeCalculator()
	if got := gz.GetFinalGrade(); got != "F" {
		t.Fatalf("no grades => want F, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatal("Assignment.String mismatch")
	}
	if Exam.String() != "exam" {
		t.Fatal("Exam.String mismatch")
	}
	if Essay.String() != "essay" {
		t.Fatal("Essay.String mismatch")
	}
}
