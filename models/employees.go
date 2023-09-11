package models

type Employee struct {
	Id                      int     `json:"id"`
	FullName                string  `json:"full_name"`
	Inn                     string  `json:"inn"`
	JobTitle                string  `json:"job_title"`
	BirthDate               string  `json:"birth_date"`
	WorkExperience          string  `json:"work_experience"`
	EnterpriseID            int     `json:"enterprise_id"`
	Gender                  string  `json:"gender"`
	BirthCountry            string  `json:"birth_country"`
	BirthCity               string  `json:"birth_city"`
	Nationality             string  `json:"nationality"`
	Citizenship             string  `json:"citizenship"`
	IsResident              string  `json:"is_resident"`
	MaritalStatus           string  `json:"marital_status"`
	WorkPhone               string  `json:"work_phone"`
	MobilePhone             string  `json:"mobile_phone"`
	Email                   string  `json:"email"`
	GraduatedSchool         string  `json:"graduated_school"`
	SchoolGraduatingYear    string  `json:"school_graduating_year"`
	GraduatedUniversity     string  `json:"graduated_university"`
	GraduatedUniversityYear string  `json:"graduated_university_year"`
	Speciality              string  `json:"speciality"`
	EducationType           string  `json:"education_type"`
	GovWorkExperience       string  `json:"gov_work_experience"`
	IsDeputy                string  `json:"is_deputy"`
	IsActiveEmployee        string  `json:"is_active_employee"`
	Biography               string  `json:"biography"`
	IsKeyEmployee           string  `json:"is_key_employee"`
	ParticipateInProjects   string  `json:"participate_in_projects"`
	Subdivision             string  `json:"subdivision"`
	IsTempEmployee          string  `json:"is_temp_employee"`
	PositionEntryDate       string  `json:"position_entry_date"`
	PositionDismissalDate   string  `json:"position_dismissal_date"`
	Salary                  float32 `json:"salary"`
}
