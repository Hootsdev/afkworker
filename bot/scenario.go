package bot

import (
	"regexp"

	"worker/ocr"

	"github.com/fatih/color"
	"golang.org/x/exp/slices"
)

var locs map[string]Location

func (d *Daywalker) Snecnario(s *Scenario) (e error) {
	s.Tasks = d.Load(s.Path)
	if s.Pattern == "loop" {
		e = Loop(s, d)
	}
	if s.Pattern == "if" {
		e = Logic(s, d)
	}

	return nil
}

func KeywordHits(kw, ocr []string) int {
	res := 0
	for _, word := range kw {
		if slices.Contains(ocr, word) {
			res++
		}
	}
	return res
}

func Loop(s *Scenario, d *Daywalker) error {
	cnt := 1
	for {
		color.HiMagenta("Run LOOP scenario, #%v execution.", cnt)
		err := d.RunTasks(s.Tasks)
		if err != nil {
			color.HiRed("During #%v run something went wrong breaking the loop", cnt)
			return err
		}
		cnt++
	}
}

func Logic(s *Scenario, d *Daywalker) error {
	for {
		Decision(s, d)
	}
}

// TODO run different psm in tesseract --psm 6 ,11, 12 show something GOOD
func Decision(s *Scenario, d *Daywalker) error {
	switch currentloc := WhereIs(d); currentloc {
	case "campain":
		color.HiGreen("### RUN => CAMP ########\n")
		locs["campain"].Actions["BeginCampain"].run(d)
	case "battlescreen":
		color.HiRed("######## RUN => FIGHT ######\n")
		locs["battlescreen"].Actions["Fight"].run(d)
	case "battleresult":
		color.HiGreen("########### RUN => RETRY ####\n")
		locs["battleresult"].Actions["Retry"].run(d)
	case "campainBoss":
		color.HiGreen("####RUN => BEGIN BOSS #####\n")
		locs["campainBoss"].Actions["BeginBoss"].run(d)
	case "victory":
		color.HiGreen("####RUN => VICTORY, NEXT #####\n")
		locs["victory"].Actions["next"].run(d)
	}
	return nil
}

func Regex(s string) {
	r := regexp.MustCompile(`Stage:(?P<chapter>\d+)-(?P<stage>\d+)`)
	ch, stg := "", ""
	for k, v := range r.FindStringSubmatch(s) {
		switch k {
		case 1:
			ch = v
			break
		case 2:
			stg = v
			break
		}
	}
	if ch != "" {
		color.HiMagenta("##### STAGE: %v-%v ###########", ch, stg)
	}
}

func WhereIs(d *Daywalker) string {
	current := ocr.OCRFields(d.Peek())
	color.HiYellow("##### Where we? ##############################\n## %v ##\n", current)
	maxhits, loc := 0, ""
	for name, v := range locs {
		hits := KeywordHits(v.Keywords, current)
		if hits > maxhits {
			maxhits = hits
			loc = name
		}

	}
	if loc != "" {
		color.HiBlue("######## %v ########\n", loc)
	}

	return loc
}
