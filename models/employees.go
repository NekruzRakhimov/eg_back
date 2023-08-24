package models

type Employee struct {
	Id             int    `json:"id"`
	FullName       string `json:"full_name"`
	Inn            string `json:"inn"`
	JobTitle       string `json:"job_title"`
	BirthDate      string `json:"birth_date"`
	WorkExperience string `json:"work_experience"`
}
