package client

import (
	"errors"
	"net/url"
	"regexp"
)

const (
	RADIO = iota
	TEXTAREA
	CHECKBOX
	SELECT
)

type Question struct {
	ID         string
	Type       int
	Selections []string
}

type Form struct {
	URL       string
	Questions []Question
	IoNF      string
	PostedFNS string
	ValCode   string
}

func ParseForm(content string) (f Form, err error) {
	actionPattern := regexp.MustCompile(`action="(\w+\.aspx\?c=\d+)"`)
	action := actionPattern.FindStringSubmatch(content)
	if len(action) < 2 {
		err = errors.New("No matching actions")
		//fmt.Println("current action", content)
		return f, err
	}
	f.URL = "https://tellburgerking.com.cn/" + action[1]

	ioNFPattern := regexp.MustCompile(`name="IoNF" value="(\d+)"`)
	IoNF := ioNFPattern.FindStringSubmatch(content)
	if len(IoNF) >= 2 {
		f.IoNF = IoNF[1]
	}

	postedFNSPattern := regexp.MustCompile(`name="PostedFNS" value="(.+?)"`)
	PostedFNS := postedFNSPattern.FindStringSubmatch(content)
	if len(PostedFNS) >= 2 {
		f.PostedFNS = PostedFNS[1]
	}

	f.Questions = make([]Question, 0, 0)

	radioPattern := regexp.MustCompile(`type="radio" name="(\w+\d+)" value="(\d+)"`)
	radios := radioPattern.FindAllStringSubmatch(content, -1)
	radioQuestions := make(map[string][]string)
	for _, radio := range radios {
		_, ok := radioQuestions[radio[1]]
		if !ok {
			radioQuestions[radio[1]] = make([]string, 0, 0)
		}
		radioQuestions[radio[1]] = append(radioQuestions[radio[1]], radio[2])
	}
	for k, v := range radioQuestions {
		q := Question{
			ID:         k,
			Type:       RADIO,
			Selections: v,
		}
		f.Questions = append(f.Questions, q)
	}

	textareaPattern := regexp.MustCompile(`textarea name"(\w+\d+)"`)
	textareas := textareaPattern.FindAllStringSubmatch(content, -1)
	for _, textarea := range textareas {
		q := Question{
			ID:         textarea[1],
			Type:       TEXTAREA,
			Selections: []string{},
		}
		f.Questions = append(f.Questions, q)
	}

	checkboxPattern := regexp.MustCompile(`type="checkbox" name="(\w+?\d+)"`)
	checkboxes := checkboxPattern.FindAllStringSubmatch(content, -1)
	for _, checkbox := range checkboxes {
		q := Question{
			ID:         checkbox[1],
			Type:       CHECKBOX,
			Selections: []string{},
		}
		f.Questions = append(f.Questions, q)
	}

	selectPattern := regexp.MustCompile(`select id="(\w+?\d+)"`)
	selects := selectPattern.FindAllStringSubmatch(content, -1)
	for _, select_ := range selects {
		q := Question{
			ID:         select_[1],
			Type:       SELECT,
			Selections: []string{},
		}
		f.Questions = append(f.Questions, q)
	}

	textPattern := regexp.MustCompile(`type="text" name="(\w+?\d+)"`)
	texts := textPattern.FindAllStringSubmatch(content, -1)
	for _, text := range texts {
		q := Question{
			ID:         text[1],
			Type:       TEXTAREA,
			Selections: []string{},
		}
		f.Questions = append(f.Questions, q)
	}

	valCodePattern := regexp.MustCompile(`class="ValCode"\D*(\d+)`)
	valCodes := valCodePattern.FindAllStringSubmatch(content, -1)
	for _, valCode := range valCodes {
		f.ValCode = valCode[1]
	}
	return
}

func (f *Form) FormatForm() url.Values {
	values := make(url.Values)
	values.Set("IoNF", f.IoNF)
	values.Set("PostedFNS", f.PostedFNS)

	for _, question := range f.Questions {
		switch question.Type {
		case RADIO:
			values.Set(question.ID, question.Selections[0])
		case TEXTAREA:
			values.Set(question.ID, "")
		case CHECKBOX:

		case SELECT:
			values.Set(question.ID, "9")
		}
	}
	return values
}
