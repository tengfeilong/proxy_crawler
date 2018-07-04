package models

import (
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
)


type MovieInfo struct{ 
  Id int64
  Movie_id int64
  Movie_name string
  Movie_pic string
  Movie_director string
  Movie_writer string
  Movie_country string
  Movie_language string
  Movie_main_character string
  Movie_type string
  Movie_on_time string
  Movie_span string
  Movie_grade string
  Ping_jia string
}


func AddMovie(movie_info *MovieInfo)(int64,error){
	movie_info.Id = 0
	id,err := db.Insert(movie_info)
	return id,err
}

func GetMovieDirector(movieHtml string) string{
	if movieHtml == ""{
		return ""
	}
	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieName(movieHtml string)string{
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieMainCharacters(movieHtml string)string{
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	mainCharacters := ""
	for _,v := range result{
		mainCharacters += v[1] + "/"
	}

	return strings.Trim(mainCharacters, "/")
}

func GetMovieGrade(movieHtml string)string{
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}
	return string(result[0][1])
}

func GetMovieGenre(movieHtml string)string{
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	movieGenre := ""
	for _,v := range result{
		movieGenre += v[1] + "/"
	}
	return strings.Trim(movieGenre, "/")
}

func GetMovieOnTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieRunningTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}


func GetMovieUrls(movieHtml string)[]string{
	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/.*?)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	var movieSets []string
	for _,v := range result{
		movieSets = append(movieSets, v[1])
	}

	return movieSets
}
