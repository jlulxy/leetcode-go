package main


type TrivalRecord struct {
	stationName string
	time        int
}

type stationRecord struct {
	times   int
	NowRecord float64
}

type UndergroundSystem struct {
	trivalData map[int]TrivalRecord
	StationConData map[string](map[string]stationRecord)

}


func Constructor() UndergroundSystem {
	cons := UndergroundSystem{}
	cons.trivalData = make(map[int]TrivalRecord,20000)
	StationConData := make(map[string](map[string]stationRecord),10)
	cons.StationConData = StationConData
	return cons
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	// 记录行程
	this.trivalData[id] = TrivalRecord{stationName,t}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	// 提取进展行程
	inRecode := this.trivalData[id]
	inName := inRecode.stationName
	inTime := inRecode.time
	// 判断是否存在
	if data ,ok := this.StationConData[inName][stationName];ok{
		tempdata := stationRecord{}
		tempdata.NowRecord = (float64(data.times) * data.NowRecord + float64(t-inTime))/float64(data.times+1)
		tempdata.times = data.times + 1
		this.StationConData[inName][stationName] = tempdata
	}else {
		if inStation,ok := this.StationConData[inName];ok{
			inStation[stationName] = stationRecord{1,float64(t-inTime)}
		}else{
			temp := make(map[string]stationRecord)
			temp[stationName] = stationRecord{1,float64(t-inTime)}
			this.StationConData[inName]= temp
		}
	}

}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	redata :=float64(0)
	if data ,ok := this.StationConData[startStation][endStation];ok{
		redata =  data.NowRecord
	}
	return redata
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
