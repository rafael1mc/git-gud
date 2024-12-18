package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" ---== Execution Started ==--- \n ")
	currBranch, err := currentBranch()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current Branch:", currBranch, "\n Make sure it's the default branch.")

	commitInDate(time.Now(), true)
	today := today()

	// initial date
	currDate := lastDayLastYear()

	// setup file that will be changed
	err = updateFile()
	if err != nil {
		panic(err)
	}
	err = add(tmpFileName)
	if err != nil {
		panic(err)
	}

	for {
		currDate = currDate.Add(1 * time.Hour * 24)
		if !currDate.Before(today) {
			break
		}

		weekday := currDate.Weekday()
		if weekday == time.Sunday || weekday == time.Saturday {
			continue
		}

		err = createHistory(currDate)
		if err != nil {
			panic(err) // TODO: add log instead to not fail everything?
		}
	}

	fmt.Println("Applying all date")
	err = rebaseDates()
	if err != nil {
		panic(err)
	}

	fmt.Println("Pushing")
	err = pushCurrent(true)
	if err != nil {
		panic(err)
	}
}

func createHistory(date time.Time) error {
	commitCount := RandInt(1, 25) // 1-25 commits
	fmt.Printf("(%d) Creating history in %v\n", commitCount, date)

	for range commitCount {
		err := updateFile()
		if err != nil {
			return err
		}
		date = date.Add(1 * time.Minute)
		err = commitInDate(date, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func today() time.Time {
	return time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		0,
		0,
		0,
		0,
		time.Now().Location(),
	)
}

func lastDayLastYear() time.Time {
	return time.Date(
		time.Now().Year(),
		1,
		0,
		12,
		0,
		0,
		0,
		time.Now().Location(),
	) // one less day than Jan 1st
}
