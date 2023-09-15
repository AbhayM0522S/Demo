package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// middleware ,helper-file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {
	fmt.Println("PROGRAM TO PERFORM BUILD API....")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "121", CourseName: "GO-LANG", CoursePrice: 259, Author: &Author{Fullname: "Abhay mandal", Website: "LCO.DEV"}})
	courses = append(courses, Course{CourseId: "122", CourseName: "CORE JAVA", CoursePrice: 567, Author: &Author{Fullname: "Kishan", Website: "CO.DEV"}})
	courses = append(courses, Course{CourseId: "123", CourseName: "SQL", CoursePrice: 578, Author: &Author{Fullname: "Kishore", Website: "LC.DEV"}})
	courses = append(courses, Course{CourseId: "124", CourseName: "PYTHON", CoursePrice: 448, Author: &Author{Fullname: "Venkatesh", Website: "P.DEV"}})
	//routing
	log.Fatal(http.ListenAndServe(":8000", r))
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>WELCOME TO THE PAGE OF BUILD API...</h1>"))

}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL COURSES....")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ONE COURSES....")
	w.Header().Set("Content-Type", "application/json")
	//grab id from the request
	params := mux.Vars(r)
	// loop through courses,find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("NO COURSE FOUND WITH GIVEN ID")
	return
}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE ONE COURSE....")
	//what if body is nil
	if r.Body == nil {
		json.NewEncoder(w).Encode("PLEASE SEND SOME DATA....")
	}
	//what about-{}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("NO DATA INSIDE JSON....")
		return
	}
	//generate a unique id convert into integer
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")
	//first -grab id from request
	params := mux.Vars(r)

	//loop ,id,remove,add with my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("DATA UPDATE SUCCESFULLY ")
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//send a id when id is not found...
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE THE COURSE")
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//loop ,id,remove(index,index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
