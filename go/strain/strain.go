package strain

type Ints []int
type Lists [][]int
type Strings []string

func (list Ints) Keep(fn func(int) bool) Ints {
	if list == nil {
		return list
	}

	newList := make(Ints, 0, len(list))

	for _, v := range list {
		if fn(v) == true {
			newList = append(newList, v)
		}
	}

	return newList
}

func (list Ints) Discard(fn func(int) bool) Ints {
	if list == nil {
		return list
	}

	newList := make(Ints, 0, len(list))

	for _, v := range list {
		if fn(v) == false {
			newList = append(newList, v)
		}
	}

	return newList
}

func (list Lists) Keep(fn func([]int) bool) Lists {
	if list == nil {
		return list
	}

	newList := make(Lists, 0, len(list))

	for _, v := range list {
		if fn(v) == true {
			newList = append(newList, v)
		}
	}

	return newList
}

func (list Lists) Discard(fn func([]int) bool) Lists {
	if list == nil {
		return list
	}

	newList := make(Lists, 0, len(list))

	for _, v := range list {
		if fn(v) == false {
			newList = append(newList, v)
		}
	}

	return newList
}

func (list Strings) Keep(fn func(string) bool) Strings {
	if list == nil {
		return list
	}

	newList := make(Strings, 0, len(list))

	for _, v := range list {
		if fn(v) == true {
			newList = append(newList, v)
		}
	}

	return newList
}

func (list Strings) Discard(fn func(string) bool) Strings {
	if list == nil {
		return list
	}

	newList := make(Strings, 0, len(list))

	for _, v := range list {
		if fn(v) == false {
			newList = append(newList, v)
		}
	}

	return newList
}
