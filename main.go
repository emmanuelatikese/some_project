package main

import (
	"fmt"
	"log"
	other_funcs "master_algo/utils"
	"os"
	"strings"
)

var (
	filename string
)

func init(){
	// get the file path with os and check if it a file.

	if len(os.Args) < 2 {
		fmt.Println("No argument was provided");
		return;
	}

	filename := os.Args[1];
	// checking for mp3
	if strings.HasSuffix(filename, ".mp3"){
			// if mp3 convert to wav and remain the file name
			filename, err := other_funcs.ConvertMP3ToWav(filename);
			if err != nil {
				log.Println(err);
				return;
			}

			_,err = os.Open(filename)
			if !(os.IsExist(err))  {
				fmt.Println("The filename:", filename, "has been converted to wav file.")
			}
	}



}

func main (){
	// function to convert wav to x(n) of sample number (N) return list of x(n)
	// function to calculate magnitude of each frequencies 0 - 20K.
	// function to convert magnitude of each frequencies to db
	// function to create a csv file.
	// if possible to can run a python file in golang if possible to get the actual virtual.
	fmt.Println("Hello golang");
}