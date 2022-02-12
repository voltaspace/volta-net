package show


func RemoveEmptySlice(oldSlice []string) (newSlice []string){
	for _,v := range oldSlice {
		if v != "" {
			newSlice = append(newSlice,v)
		}
	}
	return newSlice
}