package objects

import (
	"math/rand"
	"time"
	"fmt"
)

// GetObjectsList comment on exported function
func GetObjectsList() map[string]int {
	objectsMap := make(map[string]int)
	var objectsList = [3]string{
		"chest",
		"chair",
		"table",
	}

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	numOfObjects := r.Intn(5)
	listLen := len(objectsList)

	for i := 0; i < numOfObjects; i++ {
		objectName := objectsList[r.Intn(listLen)]

		if _, ok := objectsMap[objectName]; ok {
			objectsMap[objectName]++
		} else {
			fmt.Println(objectsMap)
			objectsMap[objectName] = 1
		}
	}

	return objectsMap
}
