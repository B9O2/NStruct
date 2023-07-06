package ScrollArray

type ScrollArray struct {
	eid    uint
	maxId  uint
	array  []interface{}
	cursor uint
}

// Append 增加元素，返回唯一的元素id
func (ra *ScrollArray) Append(e interface{}) uint {
	ra.array[ra.cursor] = e
	if ra.cursor >= ra.maxId {
		ra.cursor = 0
	} else {
		ra.cursor += 1
	}
	ra.eid += 1
	return ra.eid - 1
}

// LoadWithEid 通过唯一的元素id查找元素，布尔返回值为假则代表元素id不存在或已被清除
func (ra *ScrollArray) LoadWithEid(eid uint) (interface{}, bool) {
	d := ra.eid - eid
	size := ra.maxId + 1
	if d > size || d <= 0 {
		return nil, false
	} else {
		id := eid % size
		return ra.array[id], true
	}
}

// Load 通过相对id查找元素
func (ra *ScrollArray) Load(id uint) (interface{}, bool) {
	size := ra.maxId + 1
	if id >= size {
		return nil, false
	}
	id += ra.cursor
	if id >= size {
		id = id % size
	}
	return ra.array[id], true
}

func (ra *ScrollArray) Range(f func(interface{}) bool) {
	id := ra.cursor
	lastId := int(ra.cursor) - 1
	if lastId < 0 {
		lastId = int(ra.maxId+1) + lastId
	}
	for {
		if !f(ra.array[id]) {
			break
		}
		if id == uint(lastId) {
			//遍历结束
			break
		} else {
			if id < ra.maxId {
				id += 1
			} else {
				id = 0
			}
		}
	}
}

func NewScrollArray(size uint) *ScrollArray {
	if size < 1 {
		size = 1
	}
	return &ScrollArray{
		array:  make([]interface{}, size),
		maxId:  size - 1,
		eid:    0,
		cursor: 0,
	}
}
