package domain

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

// What the host will define
type Question struct {
	ID          string   `json:"id"`
	Prompt      string   `json:"prompt"`       // The question that will be shown to the user
	Choices     []string `json:"choices"`      // List of possible answers to the prompt
	AnswerIndex int      `json:"answer_index"` // The index of the correct solution
	TimeLimit   int      `json:"time_limit"`   // Time limit per question
}

// What the players will submit
type Answer struct {
	QuestionID  string `json:"question_id"`  // What question the players is answering
	ChoiceIndex int    `json:"choice_index"` // The actual answer they selected
}

type Quiz struct {
	ID               string     `json:"id"`
	Title            string     `json:"title"`               // The name of the quiz
	OwnerID          string     `json:"owner_id"`            // The UUID of the user who owns the quiz
	DefaultTimeLimit int        `json:"default_time_limit"`  // The time limit that will be set if the question does not contain its own timiing
	Questions        []Question `json:"questions,omitempty"` // The list of question a quiz has might be empty in some cases like when we want to ask an LLM for questions
}

type QuizChoice struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type GameRoom struct {
	ID                   string            `json:"id"`
	HostID               string            `json:"host_id"` // UUID of the host
	Players              map[string]string `json:"players"` // UUID -> username
	Quiz                 *Quiz             `json:"quiz,omitempty"`
	CurrentQuestionIndex int               `json:"current_question_index"`
	Started              bool              `json:"started"`
	Scores               map[string]int    `json:"scores"`

	QuizOptions []QuizChoice   `json:"quiz_options,omitempty"`
	QuizVotes   map[string]int `json:"quiz_vote,omitempty"`
	QuizChosen  string         `json:"quiz_chosen,omitempty"`
}
