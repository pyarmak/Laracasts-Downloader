package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/cheggaaa/pb"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"regexp"
	"time"
	"path"
)

func usage() {
	banner := fmt.Sprintf("%v USERNAME PASSWORD [DIRECTORY]", os.Args[0])
	log.Fatal(banner)
}

func parseOptions() config {
	if len(os.Args) < 3 {
		usage()
	}

	config := config{
		Username:  os.Args[1],
		Password:  os.Args[2],
	}

	if len(os.Args) > 3 {
		config.Directory = os.Args[3]
	} else {
		config.Directory = "."
	}

	return config
}

func main() {
	start := time.Now()
	// Check if username, password and directory (opt) is set
	config := parseOptions()
	scraper := newScraper(config)

	// First login to get more data in views
	log.Println("Logging in...")
	err := scraper.Login()

	if err != nil {
		log.Fatal(err)
	}

	// Find all lessons from /all
	log.Println("Building list of available lessons...")
	lessons, err := scraper.GetAllAvailableLessons()
	log.Printf("Found %v lessons", len(lessons))

	if err != nil {
		log.Fatal(err)
	}

	// Loop all lessons and download them
	for i, lesson := range lessons {
		log.Printf("Checking lesson %v/%v (%v)", i+1, len(lessons), lesson.Name)
		err = scraper.DownloadLesson(lesson)
		if err != nil {
			log.Printf("Skipping: %v", err)
		}
	}
	elapsed := time.Since(start)
  log.Printf("Finished in %v.", elapsed.String())
}

type config struct {
	Username  string
	Password  string
	Directory string
}

type scraper struct {
	// Name string
	BaseURL string
	Client  http.Client
	Username string
	Password string
	Directory string
}

type lesson struct {
	ID   int
	Name string
	URL  string
	Type string
	Series string
}

// Makes a string safe to use in a file name (e.g. for saving file atttachments)
func sanitize(text string) string {
	// Start with lowercase string
	fileName := strings.ToLower(text)
	fileName = path.Clean(path.Base(fileName))
	fileName = strings.Trim(fileName, " ")

	// Remove all other unrecognised characters - NB we do allow any printable characters
	legal, err := regexp.Compile(`[^\w\s\d\-_~,;\[\]\(\)]`)
	if err == nil {
		fileName = legal.ReplaceAllString(fileName, "")
	}

	// Remove any double dashes caused by existing - in name
	fileName = strings.Replace(fileName, "--", "-", -1)

	// NB this may be of length 0, caller must check
	return fileName
}

// Determine what the proper filename for a lesson should be
func (l *lesson) GetFilename(contentType string) (string, error) {
	basename := l.Name
	pieces := strings.Split(l.URL, "/")

	pieces = strings.Split(contentType, "/")
	extension := pieces[len(pieces)-1]

	return fmt.Sprintf("%v-%v.%v", strconv.Itoa(l.ID), basename, extension), nil
}

func newScraper(config config) scraper {
	s := scraper{}
	s.BaseURL = "https://laracasts.com"

	s.Username = config.Username
	s.Password = config.Password
	s.Directory = config.Directory

	jar, _ := cookiejar.New(nil)
	s.Client = http.Client{
		Jar: jar,
	}
	return s
}

func AppendIfMissing(slice []string, i string) []string {
    for _, ele := range slice {
        if ele == i {
            return slice
        }
    }
    return append(slice, i)
}

// Find all lesson on /all
// To get the lessonID you have to be logged in
func (s *scraper) GetAllAvailableLessons() ([]lesson, error) {
	episodes := []lesson{}

	url := s.BaseURL + "/all"
	additionalEpisodes, err := s.GetAvailableLessons(url)
	if err != nil {
		log.Fatal(err)
	}
	episodes = append(episodes, additionalEpisodes...)
	resp, err := s.Client.Get(url)

	if err != nil {
		return nil, err
	}

	// Find all pagination pages
	re := regexp.MustCompile(`(?s)(https://laracasts\.com/all\?page=\d)`)
	defer resp.Body.Close()
  contents, err := ioutil.ReadAll(resp.Body)
	pages := re.FindAllString(string(contents), -1)
	uniquePages := []string{}
	for _,page := range pages {
		uniquePages = AppendIfMissing(uniquePages, page)
	}

	for _,url := range uniquePages {
  	additionalEpisodes, err := s.GetAvailableLessons(url)
		if err != nil {
			log.Fatal(err)
		}
		episodes = append(episodes, additionalEpisodes...)
	}

	return episodes, nil
}

func (s *scraper) GetAvailableLessons(url string) ([]lesson, error) {
	episodes := []lesson{}

	resp, err := s.Client.Get(url)

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)

	// Find all links to lessons
	links := doc.Find(".container a.js-lesson-title")
	links.Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		name := s.Text()

		// Find the lessonID
		p := s.ParentsFiltered("li")
		input := p.Find("[name='lesson-id']")
		str, _ := input.Attr("value")
		lessonID, _ := strconv.Atoi(str)
		typ, _ := p.Find("[name='type']").Attr("value")
		typ = strings.ToLower(typ)
		typ = strings.Replace(typ, "laracasts\\", "", -1)
		series := p.Find(".list-group-item__details__sub").Children().First().Text()

		lesson := lesson{}
		lesson.ID = lessonID
		lesson.URL = href
		lesson.Name = sanitize(name)
		lesson.Type = typ
		lesson.Series = sanitize(series)

		episodes = append(episodes, lesson)
		// log.Println("Series name for " + lesson.Name + ": " + lesson.Series)
	})
	return episodes, nil
}

// Login to laracasts
func (s *scraper) Login() error {
	u := s.BaseURL + "/sessions"
	resp, err := s.Client.PostForm(u,
		url.Values{"email": {s.Username}, "password": {s.Password}})
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("login return wrong status code: %v, expected %v. Is your username/password correct?",
			resp.StatusCode, 200)
	}

	return nil
}

// Download a specific lesson and put it in a directory
func (s *scraper) DownloadLesson(lesson lesson) error {
	url := s.BaseURL + "/downloads/" + strconv.Itoa(lesson.ID) + "?type=" + lesson.Type
	path, basepath := "", ""

	resp, err := s.Client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return nil
	}

	headers := resp.Header
	filename, err := lesson.GetFilename(headers["Content-Type"][0])

	if len(lesson.Series) > 0 {
		basepath = s.Directory + "/" + lesson.Series
		os.MkdirAll(basepath, 0777)
		path = basepath + "/" + filename
	} else {
		path = s.Directory + "/" + filename
	}
	// log.Println(path)

	// Open the destination, return an error when the file already exists
	dest, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	defer dest.Close()

	if err != nil {
		// Nog a "already exists" error: blow up
		if !os.IsExist(err) {
			log.Fatal(err)
		}

		// OpenFile() + os.O_EXCL doesn't return a File
		dest, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		defer dest.Close()

		// Check if the video sizes online and local are the same
		fileInfo, _ := dest.Stat()
		if (fileInfo.Size() == resp.ContentLength) {
			return fmt.Errorf("%v already exists (and is the same size)", filename)
		}
		return nil
	}

	// log.Println("Starting downloading", url, "-", err)

	// if (true) {
	// 	_, err := io.Copy(dest, resp.Body)
	// 	if err != nil {

	// 		log.Println("Error while downloading", url, "-", err)
	// 		return nil
	// 	}

	// 	return nil
	// }

	// Create new progressbar
	bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES)
	bar.ShowSpeed = true
	bar.SetRefreshRate(time.Millisecond * 10)
	bar.Start()

	// create multi writer
	writer := io.MultiWriter(dest, bar)

	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return nil
	}
	bar.Finish()

	return nil
}
