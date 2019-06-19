package scheduling

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmesserli/vros/config"
	"github.com/whiteshtef/clockwork"
)

type scheduledChecker struct {
	Config *config.Config
}

func ScheduleJobs(config *config.Config) {
	checker := scheduledChecker{
		Config: config,
	}

	s := clockwork.NewScheduler()
	s.Schedule().Every().Minute().Do(checker.checkActionNeeded)
}

func (c scheduledChecker) checkActionNeeded() {
	now := time.Now()

	var dowCode int // Monday = 1, ..., Sunday = 7
	if now.Weekday() == time.Sunday {
		dowCode = 7
	} else {
		dowCode = int(now.Weekday())
	}

	hour, minute := now.Hour(), now.Minute()
	offsetDuration, _ := time.ParseDuration(fmt.Sprintf("%vm", c.Config.Verlesung.ReminderMinutesBefore))
	reminderTime := now.Add(offsetDuration)
	rHour, rMinute := reminderTime.Hour(), reminderTime.Minute()

	for _, verlesung := range c.Config.Verlesung.Entries {
		dayMatch := false
		for _, day := range verlesung.Days {
			if int(day) == dowCode {
				dayMatch = true
				break
			}
		}

		if !dayMatch {
			continue
		}

		hourMinute := strings.Split(verlesung.Time, ":")
		vHour, _ := strconv.Atoi(hourMinute[0])
		vMinute, _ := strconv.Atoi(hourMinute[1])

		if vHour == hour && vMinute == minute {
			// Time to send the report
			fmt.Printf("The presence report for '%s' should be sent now (%s)", verlesung.Name, verlesung.Time)
		}

		if vHour == rHour && vMinute == rMinute {
			// Time to send the reminders
			fmt.Printf("The reminders for '%s' should be sent now (%s)", verlesung.Name, verlesung.Time)
		}
	}
}
