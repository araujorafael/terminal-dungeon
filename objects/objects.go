package objects

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

// GetMessage build the room message based on info passed by map
func GetMessage(objects map[string]int) string {
	if len(objects) == 0 {
		return "The room is empty; As you look for something inside, the room stares at you with his emptyness"
	}

	var roomMessage = "You look inside the room and find: "
	var itemsRoomItems []string

	for object, num := range objects {
		objectString := fmt.Sprintf("%d %s", num, object)
		itemsRoomItems = append(itemsRoomItems, objectString)
	}

	details := strings.Join(itemsRoomItems, ", ")
	return roomMessage + details + "; you dont know why, but appers the room is watching you"
}
