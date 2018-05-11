package ranking

type Weigher interface {
	//权重计算
	WeightCalc(Weigher) bool
	//计算权重返回的key
	GroupKey() string
}

func Weight(data []Weigher, fc func(Weigher)) {
	if len(data) == 0 {
		return
	}
	ms := make(map[string]*sortItems, 1)
	for _, item := range data {
		k := item.GroupKey()
		if _, ok := ms[k]; !ok {
			ms[k] = &sortItems{}
		}
		ms[k].addItem(item)
	}

	for _, b := range ms {
		fc(b.calcWeigth()) //计算权重
	}
}

type sortItems struct {
	target Weigher //最终需要的 , 用于比较
	data   []Weigher
}

func (b *sortItems) addItem(item Weigher) {
	b.data = append(b.data, item)
}

//计算权重
func (b *sortItems) calcWeigth() Weigher {
	for _, data := range b.data {
		if b.target == nil {
			b.target = data
			continue
		}
		if data.WeightCalc(b.target) {
			b.target = data //重新赋值
			continue
		}
	}
	return b.target
}
