package Insert_Interval



///**
// * Definition for an interval.
type Interval struct {
	   Start int
	   End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	returnItervals := []Interval{}
	nextInterval := Interval{Start:-1, End:-1}
	length := len(intervals)
	if length == 0{
		returnItervals = append(returnItervals, newInterval)
		return returnItervals
	}
	for index, ite := range intervals{
		if nextInterval.Start >= 0 {
			if nextInterval.End >= 0 {
				returnItervals = append(returnItervals, ite)
				continue
			}
			if ite.Start > newInterval.End {
				nextInterval.End = newInterval.End
				returnItervals = append(returnItervals, nextInterval)
				returnItervals = append(returnItervals, ite)
				continue
			}

			if ite.Start <= newInterval.End {
				if ite.End >= newInterval.End {
					nextInterval.End = ite.End
					returnItervals = append(returnItervals, nextInterval)
					continue
				}

				if index == length -1 {
					nextInterval.End = newInterval.End
					returnItervals = append(returnItervals, nextInterval)
				}

			}



		} else {
			if newInterval.Start <= ite.Start {
				nextInterval.Start = newInterval.Start
				if newInterval.End < ite.Start {
					returnItervals = append(returnItervals, newInterval)
					returnItervals = append(returnItervals, ite)
					nextInterval.End = newInterval.End
					continue
				}

				if newInterval.End <= ite.End {
					nextInterval.End = ite.End
					returnItervals = append(returnItervals, nextInterval)
					continue
				}

				if (newInterval.End >= ite.End) && (index == length -1) {
					nextInterval.End = newInterval.End
					returnItervals = append(returnItervals, nextInterval)
				}

			}

			if newInterval.Start > ite.Start {
				if newInterval.Start > ite.End {
					returnItervals = append(returnItervals, ite)
					if index == length -1 {
						returnItervals = append(returnItervals, newInterval)
					}
					continue
				}

				nextInterval.Start = ite.Start
				if newInterval.End <= ite.End {
					nextInterval.End =ite.End
					returnItervals = append(returnItervals, nextInterval)
					continue
				}
				if (newInterval.End >= ite.End) && (length-1 == index) {
					nextInterval.End = newInterval.End
					returnItervals = append(returnItervals, nextInterval)
					continue
				}

			}

		}

	}
	return returnItervals
}