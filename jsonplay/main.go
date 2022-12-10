package main

import (
	"fmt"
	"strings"
)

//todo: code for nested keys {"os": {"Windows": {"vers": "myBlah", "vers2": "return of blah"}, "Mac": "OSX", "Linux": "Ubuntu"},"compilers": "gcc", "stuff": "morestuff"}
//todo: will need to take an input of some sort of nested key.  ex os.windows : blah.  then out put would be "blah": "Windows OS"
//todo: where to put a comma for when I take in two inputs
//todo: convert this to a function

func main() {
	//I'm thinking this should accept (string map, json (as string)) then return a []byte

	//this is to build the new json
	var b strings.Builder

	//temp strings
	var tempStr string
	var tempStrArry []string
	var charCount, pos int
	var runes []rune

	unstructuredJson := `{"os": {"Windows": "Windows OS","Mac": "OSX","Linux": "Ubuntu"},"compilers": "gcc", "stuff": "morestuff"}`

	//we need to use the " " in the replace if they are not used, we run the risk of replacing something we do not want to replace.
	//todo:  if not sent in

	oldName := `"os"`
	newName := `"blah"`

	//oldName := `"compilers"`
	//newName := `"blah"`

	// this will change the json field names, but will not remove the suplus json we to not care about
	newUnstructuredJson := strings.Replace(unstructuredJson, oldName, newName, 1) //todo: does this need to be in the loop?  maybe...
	fmt.Println("new Json:" + newUnstructuredJson)

	runes = []rune(newUnstructuredJson)
	charCount = len(newUnstructuredJson) //i do not think this will change...

	//start of the new json
	b.WriteString(`{`)

	//testing to see if more json vaules are nested, or is this just single value
	tempStr = newName + (": {")
	fmt.Println("mytest:" + tempStr)
	if (strings.Contains(newUnstructuredJson, tempStr)) == true {
		//b.WriteString(newName)
		//b.WriteString(`: `)
		//b.WriteString(`{`)

		/*
			closeCount = strings.Count(mybetween, "}") //this will let me know if there are multiple nested fields.  and how many there are.
			if closeCount == 0 {
				b.WriteString(mybetween) //I'm sure this will change...
			} else {
				//need to do some more magic. I'm not thinking about it for now.
				//closeCount gives me the number of iterantion I need to loop through.  I have some substring craziness in my head I need to logic through ang go from there.
			}*/

		fmt.Println(charCount)
		pos = strings.Index(newUnstructuredJson, tempStr)
		fmt.Println(newUnstructuredJson)
		fmt.Println(tempStr)
		fmt.Println(pos)

		fmt.Println("pos:charCount")
		fmt.Println(string(runes[pos:charCount]))
		tempStr = string(runes[pos:charCount])

		tempStrArry = strings.Split(tempStr, `}`)

		fmt.Println("this one")
		fmt.Println(tempStrArry[0])
		b.WriteString(tempStrArry[0])

		b.WriteString(`}`)
	} else { //single value not nested

		fmt.Println(charCount)
		pos = strings.Index(newUnstructuredJson, newName)
		fmt.Println(pos)

		fmt.Println(string(runes[0:pos]))

		tempStr = string(runes[pos:charCount])
		fmt.Println(tempStr)
		//this will get rid of the comma
		tempStrArry = strings.Split(tempStr, `,`) // if this fails we can assum this is the last key in the json
		fmt.Println("this one")
		fmt.Println(tempStrArry[0])

		//todo: need logic for the if the last change to be made
		b.WriteString(tempStrArry[0])

		fmt.Println(string(runes[0:(charCount)]))

		/*
			tempStr = newName + (": ")
			fmt.Println("tempSring: " + tempStr)
			//does not work because between function is not pulling the , after the tempStr...
			mybetween = between(newUnstructuredJson, tempStr, ",") // if this fails it is the last item in the json and I need to get the between on } and not ,

			fmt.Println("stuff between: " + mybetween)
			b.WriteString(mybetween)
		*/
	}

	b.WriteString(`}`) //close json

	//return this string
	fmt.Println(b.String())
	/*
		//This can be used to validate the string builder is json (maybe need to confirm)  will need to err if not json and return error.
		jsondata, err := json.Marshal(b.String())
		if err != nil {
			fmt.Println(err)
			fmt.Println("json failed!")
		} else {
			fmt.Println("here is the byte array:")
			fmt.Println(jsondata)
		}*/

	//unstructuredJson := `{"os": {"Windows": "Windows OS","Mac": "OSX","Linux": "Ubuntu"},"compilers": "gcc", "stuff": "morestuff"}`

}
