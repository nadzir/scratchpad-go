package jobdb

import (
	"database/sql"
	"fmt"
	"time"

	// "log"

	// mysql driver
	simplejson "github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	dbDriver            = "mysql"
	dbUser              = "root"
	dbPass              = "password"
	dbName              = "analysis"
	createJobTableQuery = `
	CREATE TABLE IF NOT EXISTS jobsListing (
		id int(11) unsigned NOT NULL AUTO_INCREMENT,
		source varchar(30) DEFAULT NULL,
		crawledURL varchar(500)    DEFAULT NULL,
		jobURL varchar(500)    DEFAULT NULL,
		jobTitle varchar(500)    DEFAULT NULL,
		companyName varchar(500)    DEFAULT NULL,
		description longtext   ,
		postingDate varchar(50)    DEFAULT NULL,
		closingDate varchar(50)    DEFAULT NULL,
		minSalary int DEFAULT NULL,
		maxSalary int DEFAULT NULL,
		salaryType int DEFAULT NULL,
		crawledAt DATE DEFAULT NULL,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;`

	insertJobTableQuery = `
	INSERT INTO jobsListing(
		source,
		crawledURL,
		jobURL,
		jobTitle,
		companyName,
		description,
		postingDate,
		closingDate,
		crawledAt)   
		SELECT ?,?,?,?,?,?,?,?,?
		FROM jobsListing
		WHERE source = ?
		and jobURL = ?
		and crawledAt = ?
		`
)

func dbConn() (db *sql.DB) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db.SetMaxOpenConns(20) // Sane default
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(time.Nanosecond)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// CreateJobTable : Create Job Table
func CreateJobTable() {
	db := dbConn()
	tables, err := db.Query(createJobTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer tables.Close()
}

type JobInfo struct {
	Source      string
	CrawledURL  string
	JobURL      string
	JobTitle    string
	CompanyName string
	Description string
	PostingDate string
	ClosingDate string
}

// Log : JobInfo loggin
func (j *JobInfo) Log() {
	log.WithFields(log.Fields{
		// "source": j.Source,
		// "crawled url": j.CrawledURL,
		"job title": j.JobTitle,
		"company":   j.CompanyName,
		"job url":   j.JobURL,
		// "desc":         j.Description,
		// "posting date": j.PostingDate,
		// "closing date": j.ClosingDate,
	}).Info(j.Source)
}

// InsertJobTable : Inserting to Job Table
func InsertJobTable(jobInfo JobInfo) {
	db := dbConn()
	stmt, err := db.Prepare(insertJobTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, queryErr := stmt.Exec(
		jobInfo.Source,
		jobInfo.CrawledURL,
		jobInfo.JobURL,
		jobInfo.JobTitle,
		jobInfo.CompanyName,
		jobInfo.Description,
		newNullString(jobInfo.PostingDate),
		newNullString(jobInfo.ClosingDate),
		time.Now().Local(),
		jobInfo.Source,
		jobInfo.JobURL,
		time.Now().Local(),
	)
	if queryErr != nil {
		log.Warn(queryErr)
	}

	stmt.Close()
}

// SelectTotalJobCount : Selecting from Job Table
func SelectTotalJobCount(source string, date string) string {
	var count string

	selectQuery := `
	select count(distinct jobTitle, companyName, description)
	from jobsListing`

	selectQuery = ValidJobTitle(selectQuery)

	if source != "" {
		selectQuery = ConditionalSource(selectQuery, source)
	}
	if date != "" {
		selectQuery = ConditionalCrawledAt(selectQuery, date)
	}

	db := dbConn()
	err := db.QueryRow(selectQuery).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

// SelectPopularJob : Selecting from Job Table
func SelectPopularJob(source string, date string) *simplejson.Json {
	selectQuery := `
	select jobTitle, count(jobTitle)
	from jobsListing
	`

	selectQuery = ValidJobTitle(selectQuery)

	if source != "" {
		selectQuery = ConditionalSource(selectQuery, source)
	}
	if date != "" {
		selectQuery = ConditionalCrawledAt(selectQuery, date)
	}

	selectQuery = fmt.Sprintf(`
	%s
	group by jobTitle
	order by count(jobTitle) desc
	limit 10`, selectQuery)

	var jobTitle, count string
	json := simplejson.New()

	db := dbConn()
	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&jobTitle, &count)
		if err != nil {
			log.Fatal(err)
		}
		json.Set(jobTitle, count)
	}
	return json
}

// SelectPopularCompany : Selecting from Job Table
func SelectPopularCompany(source, date string) *simplejson.Json {
	selectQuery := `
	select companyName, count(jobTitle)
	from jobsListing
	`

	selectQuery = ValidJobTitle(selectQuery)

	if source != "" {
		selectQuery = ConditionalSource(selectQuery, source)
	}
	if date != "" {
		selectQuery = ConditionalCrawledAt(selectQuery, date)
	}

	selectQuery = fmt.Sprintf(`
	%s
	group by companyName
	order by count(companyName) desc
	limit 10`, selectQuery)

	var company, count string
	json := simplejson.New()

	db := dbConn()
	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&company, &count)
		if err != nil {
			log.Fatal(err)
		}
		json.Set(company, count)
	}
	return json
}

func newNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func UpdateSalary(id, minSalary, maxSalary string) {
	updateQuery := `
	UPDATE jobslisting
	set minSalary = ?,
	maxSalary = ?
	where id = ?
	`
	db := dbConn()
	stmt, err := db.Prepare(updateQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, queryErr := stmt.Exec(
		newNullString(minSalary),
		newNullString(maxSalary),
		id,
	)
	if queryErr != nil {
		log.Warn(queryErr)
	}

	stmt.Close()
}
