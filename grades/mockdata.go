package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "Sean",
			LastName:  "Martin",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 72,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 65,
				},
				Grade{
					Title: "Test 1",
					Type:  GradeTest,
					Score: 87,
				},
			}},
		Student{
			ID:        2,
			FirstName: "Alex",
			LastName:  "Pearson",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 72,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 65,
				},
				Grade{
					Title: "Test 1",
					Type:  GradeTest,
					Score: 87,
				},
			}},
		Student{
			ID:        3,
			FirstName: "Amy",
			LastName:  "Bacon",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 72,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 65,
				},
				Grade{
					Title: "Test 1",
					Type:  GradeTest,
					Score: 87,
				},
			}},
	}
}
