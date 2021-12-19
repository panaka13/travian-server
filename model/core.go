package model

import "fmt"

type ResourceSet struct {
	Wood  int
	Clay  int
	Iron  int
	Wheat int
}

func (r ResourceSet) serialize() string {
	return fmt.Sprintf("%d;%d;%d;%d", r.Wood, r.Clay, r.Iron, r.Wheat)
}

func (r *ResourceSet) deserialize(s string) {
	if s == "" {
		r.Wood, r.Clay, r.Iron, r.Wheat = 0, 0, 0, 0
		return
	}
	fmt.Sscanf(s, "%d;%d;%d;%d", &r.Wood, &r.Clay, &r.Iron, &r.Wheat)
}
