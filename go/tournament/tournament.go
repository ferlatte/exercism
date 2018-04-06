package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
}

func (t *team) points() int {
	return (t.wins * 3) + (t.draws)
}

type tally map[string]*team

// Tally tallys up a tournament
func Tally(r io.Reader, w io.Writer) error {
	cr := makeReader(r)
	t := make(tally)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		t, err = updateTally(t, record[0], record[1], record[2])
		if err != nil {
			return err
		}
	}
	return formatTally(t, w)
}

func updateTally(t tally, first, second, result string) (tally, error) {
	if t[first] == nil {
		t[first] = &team{name: first}
	}
	if t[second] == nil {
		t[second] = &team{name: second}
	}
	t[first].matchesPlayed++
	t[second].matchesPlayed++
	switch result {
	case "win":
		t[first].wins++
		t[second].losses++
	case "loss":
		t[first].losses++
		t[second].wins++
	case "draw":
		t[first].draws++
		t[second].draws++
	default:
		return nil, errors.New("Invalid result: " + result)
	}
	return t, nil
}

func formatTally(t tally, w io.Writer) error {
	teams := make([]team, 0)
	for _, v := range t {
		teams = append(teams, *v)
	}
	sort.Slice(teams, func(i, j int) bool {
		t1 := teams[i]
		t2 := teams[j]
		if t1.points() == t2.points() {
			return t1.name < t2.name
		}
		return t1.points() > t2.points()
	})
	_, err := io.WriteString(w, formatLine("Team", "MP",
		" W", " D", " L", " P"))
	if err != nil {
		return err
	}
	for _, team := range teams {
		_, err := io.WriteString(w, formatTeamLine(team))
		if err != nil {
			return err
		}
	}
	return nil
}

func formatTeamLine(t team) string {
	mp := fmt.Sprintf("%2d", t.matchesPlayed)
	w := fmt.Sprintf("%2d", t.wins)
	d := fmt.Sprintf("%2d", t.draws)
	l := fmt.Sprintf("%2d", t.losses)
	p := fmt.Sprintf("%2d", t.points())
	return formatLine(t.name, mp, w, d, l, p)
}

// Return a single formatted line
func formatLine(team, mp, w, d, l, p string) string {
	return team + strings.Repeat(" ", 31-len(team)) + "| " +
		mp + " | " +
		w + " | " +
		d + " | " +
		l + " | " +
		p + "\n"
}

func makeReader(r io.Reader) csv.Reader {
	cr := csv.NewReader(r)
	cr.Comma = ';'
	cr.Comment = '#'
	cr.FieldsPerRecord = 3
	return *cr
}
